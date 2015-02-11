angular.module('admin.user', ['services.crud']);

angular.module('admin.user').config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.when('/user/:action/:Id?', {
            templateUrl: '/admin/tpl/user_edit.html',
            controller: 'UserEditController',
            resolve: {
                user: ['$route', '$http',
                    function ($route, $http) {
                        var params = $route.current.params;
                        if (params.action === 'new') {
                            return {};
                        } else {
                            return $http.get('/user/' + params.Id);
                        }
                }]
            }
        });
        $routeProvider.when('/users', {
            templateUrl: '/admin/tpl/user_list.html',
            controller: 'UserListController'
        })
}]);

angular.module('admin.user').controller('UserListController', ['$scope', '$routeParams', '$http', '$q', '$location',
    function ($scope, $routeParams, $http, $q, $location) {
        $http.get('/users')
            .success(function (users) {
                $scope.users = users;
            });

        $scope.remove = function (user, $index, $event) {
            // Don't let the click bubble up to the ng-click on the enclosing div, which will try to trigger
            // an edit of this item.
            $event.stopPropagation();

            // Remove this user
            $http.delete('/user/' + user.Id).success(function () {
                // It is gone from the DB so we can remove it from the local list too
                $scope.users.splice($index, 1);
                console.log('crud.user.remove.success', user.Id);
            }).error(function (e) {
                console.log('crud.user.remove.error', user.Id, e);
            });
        };

        $scope.edit = function (user) {
            $location.path('user/edit/' + user.Id);
        }

        $scope.new = function () {
            $location.path('user/new');
        }
    }]);


angular.module('admin.user').controller('UserEditController', ['$scope', '$routeParams', '$http', '$location', 'user',
    function ($scope, $routeParams, $http, $location, user) {
        $scope.user = user.data;

        $scope.new = function () {
            $http.post('/user', $scope.user)
                .success(function () {
                    console.log('user.add.success');
                    $location.path('users');
                })
                .error(function (e) {
                    console.log('user.add.error', e);
                })
        }

        $scope.edit = function () {
            $http.put('/user/' + $scope.user.Id, $scope.user)
                .success(function () {
                    console.log('user.edit.success', $scope.user.Id);
                    $location.path('users');
                }).error(function (e) {
                    console.log('user.edit.error', $scope.user.Id, e);
                })
        }

        $scope.save = function () {
            if ($scope.user.Id === undefined) {
                $scope.new();
            } else {
                $scope.edit();
            }
        }

}]);