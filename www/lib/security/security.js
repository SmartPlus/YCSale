// Based loosely around work by Witold Szczerba - https://github.com/witoldsz/angular-http-auth
angular.module('security', ['security.service', 'security.login', 'security.interceptor', 'security.authorization']);

angular.module('security.service', ['security.retryQueue']);

angular.module('security.service').factory('security', ['$http', '$q', '$location', 'securityRetryQueue', '$modal',
    function ($http, $q, $location, queue, $modal) {
        // Redirect to the given url (defaults to '/')
        function redirect(url) {
            url = url || '/';
            $location.path(url);
        }
        var loginDialog = null;
        var allowRedirects = ['/', '/admin'];
        var loginCallback = angular.noop;

        function openLoginDialog() {
            if (loginDialog) {
                throw new Error('Trying to open a dialog that is already open!');
            }
            loginDialog = $modal.open({
                templateUrl: '/lib/security/tpl/login.html',
                controller: 'LoginFormController',
            })
            loginDialog.result.then(onLoginDialogClose);
        }

        function closeLoginDialog(success) {
            if (loginDialog) {
                loginDialog.close(success);
            }
        }

        function onLoginDialogClose(success) {
            loginDialog = null;
            if (success) {
                queue.retryAll();
            } else {
                queue.cancelAll();
                if (allowRedirects.indexOf($location.path().split('/')[1]) !== -1) {
                    redirect();
                }
            }
        }

        // Register a handler for when an item is added to the retry queue
        queue.onItemAddedCallbacks.push(function (retryItem) {
            if (queue.hasMore()) {
                service.showLogin();
            }
        });

        // The public API of the service
        var service = {

            // Get the first reason for needing a login
            getLoginReason: function () {
                return queue.retryReason();
            },

            // Show the modal login dialog
            showLogin: function () {
                if (loginDialog === null) {
                    openLoginDialog();
                }
            },

            // Attempt to authenticate a user by the given email and password
            login: function (email, password) {
                var request = $http.post('/login', {
                    email: email,
                    password: password
                });
                return request.then(function (response) {
                    service.currentUser = response.data.user;
                    if (service.isAuthenticated()) {
                        closeLoginDialog(true);
                        loginCallback(service.currentUser);
                    }
                    return service.isAuthenticated();
                });
            },

            onLoginSuccess: function (cb) {
                loginCallback = cb;
            },
            // Give up trying to login and clear the retry queue
            cancelLogin: function () {
                closeLoginDialog(false);
            },

            // Logout the current user and redirect
            logout: function (redirectTo) {
                $http.post('/logout').then(function () {
                    service.currentUser = null;
                    redirect(redirectTo);
                });
            },

            // Ask the backend to see if a user is already authenticated - this may be from a previous session.
            requestCurrentUser: function () {
                if (service.isAuthenticated()) {
                    return $q.when(service.currentUser);
                } else {
                    return $http.get('/current-user').then(function (response) {
                        service.currentUser = response.data.user;
                        return service.currentUser;
                    });
                }
            },

            // Require that there is an authenticated user
            // (use this in a route resolve to prevent non-authenticated users from entering that route)
            requireAuthenticatedUser: function () {
                var promise = service.requestCurrentUser().then(function (userInfo) {
                    if (!service.isAuthenticated()) {
                        return queue.pushRetryFn('unauthenticated-client', service.requireAuthenticatedUser);
                    }
                });
                return promise;
            },

            // Require that there is an administrator logged in
            // (use this in a route resolve to prevent non-administrators from entering that route)
            requireAdminUser: function () {
                var promise = service.requestCurrentUser().then(function (userInfo) {
                    if (!service.isAdmin()) {
                        return queue.pushRetryFn('unauthorized-client', service.requireAdminUser);
                    }
                });
                return promise;
            },
            // Information about the current user
            currentUser: null,
            // Is the current user authenticated?
            isAuthenticated: function () {
                return !!service.currentUser;
            },

            // Is the current user an adminstrator?
            isAdmin: function () {
                return !!(service.currentUser && service.currentUser.admin);
            }
        };

        return service;
}])


// The loginToolbar directive is a reusable widget that can show login or logout buttons
// and information the current authenticated user
angular.module('security.login', ['security.service']);
angular.module('security.login').directive('loginToolbar', ['security',
    function (security) {
        var directive = {
            templateUrl: '/lib/security/tpl/bar.html',
            restrict: 'E',
            replace: true,
            scope: true,
            link: function ($scope, $element, $attrs, $controller) {
                // function re-eveluation triggers update
                $scope.isAuthenticated = security.isAuthenticated;
                $scope.login = security.showLogin;
                $scope.logout = security.logout;

                $scope.$watch(function () {
                    return security.currentUser;
                }, function (currentUser) {
                    $scope.currentUser = currentUser;
                });
            }
        };
        return directive;
}]);

angular.module('security.login').controller('LoginFormController', ['$scope', 'security',
    function ($scope, security) {
        // The model for this form 
        $scope.user = {};

        // Any error message from failing to login
        $scope.authError = null;

        // The reason that we are being asked to login - for instance because we tried to access something to which we are not authorized
        // We could do something diffent for each reason here but to keep it simple...
        $scope.authReason = null;
        if (security.getLoginReason()) {
            $scope.authReason = (security.isAuthenticated()) ?
                'login.reason.notAuthorized' : 'login.reason.notAuthenticated';
        }

        // Attempt to authenticate the user specified in the form's model
        $scope.login = function () {
            // Clear any previous security errors
            $scope.authError = null;

            // Try to login
            security.login($scope.user.email, $scope.user.password).then(function (loggedIn) {
                if (!loggedIn) {
                    // If we get here then the login failed due to bad credentials
                    $scope.authError = 'login.error.invalidCredentials';
                }
            }, function (x) {
                // If we get here then there was a problem with the login request to the server
                $scope.authError = 'login.error.serverError ' + x;
            });
        };



        $scope.clearForm = function () {
            $scope.user = {};
        };

        $scope.cancelLogin = function () {
            security.cancelLogin();
        };
}]);

/************************************************************************/
angular.module('security.retryQueue', [])

// This is a generic retry queue for security failures.  Each item is expected to expose two functions: retry and cancel.
.factory('securityRetryQueue', ['$q', '$log',
    function ($q, $log) {
        var retryQueue = [];
        var service = {
            // The security service puts its own handler in here!
            onItemAddedCallbacks: [],

            hasMore: function () {
                return retryQueue.length > 0;
            },
            push: function (retryItem) {
                retryQueue.push(retryItem);
                // Call all the onItemAdded callbacks
                angular.forEach(service.onItemAddedCallbacks, function (cb) {
                    try {
                        cb(retryItem);
                    } catch (e) {
                        $log.error('securityRetryQueue.push(retryItem): callback threw an error' + e);
                    }
                });
            },
            pushRetryFn: function (reason, retryFn) {
                // The reason parameter is optional
                if (arguments.length === 1) {
                    retryFn = reason;
                    reason = undefined;
                }

                // The deferred object that will be resolved or rejected by calling retry or cancel
                var deferred = $q.defer();
                var retryItem = {
                    reason: reason,
                    retry: function () {
                        // Wrap the result of the retryFn into a promise if it is not already
                        $q.when(retryFn()).then(function (value) {
                            // If it was successful then resolve our deferred
                            deferred.resolve(value);
                        }, function (value) {
                            // Othewise reject it
                            deferred.reject(value);
                        });
                    },
                    cancel: function () {
                        // Give up on retrying and reject our deferred
                        deferred.reject();
                    }
                };
                service.push(retryItem);
                return deferred.promise;
            },
            retryReason: function () {
                return service.hasMore() && retryQueue[0].reason;
            },
            cancelAll: function () {
                while (service.hasMore()) {
                    retryQueue.shift().cancel();
                }
            },
            retryAll: function () {
                while (service.hasMore()) {
                    retryQueue.shift().retry();
                }
            }
        };
        return service;
}]);

/************************************************************************/
angular.module('security.interceptor', ['security.retryQueue'])
// This http interceptor listens for authentication failures
.factory('securityInterceptor', ['$injector', 'securityRetryQueue',
    function ($injector, queue) {
        return {
            responseError: function (originalResponse) {
                // Intercept failed requests                
                if (originalResponse.status === 401) {
                    // The request bounced because it was not authorized - add a new request to the retry queue
                    return queue.pushRetryFn('unauthorized-server', function retryRequest() {
                        // We must use $injector to get the $http service to prevent circular dependency
                        return $injector.get('$http')(originalResponse.config);
                    });
                } else {
                    noty({
                        text: originalResponse.data,
                        type: "warning"
                    });
                    return originalResponse;
                }
            }
        };
}]).config(['$httpProvider',
    function ($httpProvider) {
        $httpProvider.interceptors.push('securityInterceptor');
}]);


/***************************************************************/
angular.module('security.authorization', ['security'])

// This service provides guard methods to support AngularJS routes.
// You can add them as resolves to routes to require authorization levels
// before allowing a route change to complete
.provider('securityAuthorization', {

    requireAdminUser: ['securityAuthorization',
        function (securityAuthorization) {
            return securityAuthorization.requireAdminUser();
  }],

    requireAuthenticatedUser: ['securityAuthorization',
        function (securityAuthorization) {
            return securityAuthorization.requireAuthenticatedUser();
  }],

    $get: ['security', 'securityRetryQueue',
        function (security, queue) {
            var service = {

                // Require that there is an authenticated user
                // (use this in a route resolve to prevent non-authenticated users from entering that route)
                requireAuthenticatedUser: function () {
                    var promise = security.requestCurrentUser().then(function (userInfo) {
                        if (!security.isAuthenticated()) {
                            return queue.pushRetryFn('unauthenticated-client', service.requireAuthenticatedUser);
                        }
                    });
                    return promise;
                },

                // Require that there is an administrator logged in
                // (use this in a route resolve to prevent non-administrators from entering that route)
                requireAdminUser: function () {
                    var promise = security.requestCurrentUser().then(function (userInfo) {
                        if (!security.isAdmin()) {
                            return queue.pushRetryFn('unauthorized-client', service.requireAdminUser);
                        }
                    });
                    return promise;
                }

            };

            return service;
  }]
});