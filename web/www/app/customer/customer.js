angular.module('app.customer', ['services.crud']);

angular.module('app.customer').config(['$routeProvider',
    function ($routeProvider) {
        
        $routeProvider.when('/customers', {
            templateUrl: '/app/customer/customer_list.html',
            controller: 'CustomerListController'
        })
        $routeProvider.when('/customers/new', {
            templateUrl: '/app/customer/customer_edit.html',
            controller: 'CustomerNewController'
            
        });
        $routeProvider.when('/customer/:id', {
            templateUrl: '/app/customer/customer.html',
            controller: 'CustomerDetailController',
            resolve: {
                customer: ['$route', '$http',
                    function ($route, $http) {
                        var params = $route.current.params;
                        return $http.get('/customer/' + params.id);
                }]
            }
        });
}]);

angular.module('app.customer').controller('CustomerListController', ['$scope', '$routeParams', '$http', '$q', '$location',
    function ($scope, $routeParams, $http, $q, $location) {
        $http.get('/customers')
            .success(function (customers) {
                $scope.customers = customers;
            });

        $scope.remove = function (customer, $index, $event) {
            // Don't let the click bubble up to the ng-click on the enclosing div, which will try to trigger
            // an edit of this item.
            $event.stopPropagation();

            // Remove this customer
            $http.delete('/customer/' + customer._id).success(function () {
                // It is gone from the DB so we can remove it from the local list too
                $scope.customers.splice($index, 1);
                console.log('crud.customer.remove.success', customer._id);
            }).error(function (e) {
                console.log('crud.customer.remove.error', customer._id, e);
            });
        };

        $scope.edit = function (customer) {
            $location.path('customer/' + customer._id);
        }

        $scope.new = function () {
            $location.path('customers/new');
        }
    }]);


angular.module('app.customer').controller('CustomerNewController', ['$scope', '$routeParams', '$http', '$location',
    function ($scope, $routeParams, $http, $location, customer) {
        $scope.customer = {}

        $scope.new = function () {
            $http.post('/customer', $scope.customer)
                .success(function () {
                    console.log('customer.add.success');
                    $location.path('customers');
                })
                .error(function (e) {
                    console.log('customer.add.error', e);
                })
        }

        $scope.save = function () {
            $scope.new();
        }

}]);

angular.module('app.customer').controller('CustomerDetailController', ['$scope', '$routeParams', '$http', '$location', 'customer',
    function ($scope, $routeParams, $http, $location, customer) {
        $scope.customer = customer.data;
        $scope.action = function (a) {            
            if (a == "register") {
                $scope.getAvailableCourses()
            }

            $scope.customer_action = "/app/customer/customer_" + a + ".html"
        }

        
        $scope.getAvailableCourses = function () {
            if ($scope.courses === undefined) {
                $scope.getMyCourses()

                $http.get('/courses')
                .success(function (data) {
                    $scope.courses = data                    
                }).error(function (err) {
                    alert(err);
                });    
            }            
        }

        $scope.getMyCourses = function () {
            if ($scope.myCourses === undefined) {
                var customerId = $scope.customer._id
                $http.get('/customer/' + customerId + '/courses')
                .success(function (data) {                    
                    $scope.myCourses = data
                }).error(function (e) {
                    console.log('customer.course.error', customerId, e);
                })
            }                
        }

        $scope.action('view')
        $scope.getMyCourses()

        $scope.save = function () {
            var customerId = $scope.customer._id
            $http.put('/customer/' + customerId, $scope.customer)
                .success(function () {
                    console.log('customer.edit.success', customerId);
                    $scope.action('view')
                }).error(function (e) {
                    console.log('customer.edit.error', customerId, e);
                })
        }


        $scope.register = function (course, $index, $event) {
            // Don't let the click bubble up to the ng-click on the enclosing div, which will try to trigger
            // an edit of this item.
            $event.stopPropagation();

            var customerId = $scope.customer._id
            console.log(course)
            $http.post('/customer/' + customerId + '/register/' + course._id).success(function () {
                $scope.action('course')
                console.log('crud.customer.register.success', customer._id, course._id);
            }).error(function (e) {
                console.log('crud.customer.register.error', customer._id, course._id, e);
            });
        };

        $scope.unregister = function (course, $index, $event) {
            // Don't let the click bubble up to the ng-click on the enclosing div, which will try to trigger
            // an edit of this item.
            $event.stopPropagation();

            var customerId = $scope.customer._id
            console.log(course)
            $http.post('/customer/' + customerId + '/unregister/' + course._id).success(function () {
                $scope.action('course')
                console.log('crud.customer.unregister.success', customer._id, course._id);
            }).error(function (e) {
                console.log('crud.customer.unregister.error', customer._id, course._id, e);
            });
        };
}]);