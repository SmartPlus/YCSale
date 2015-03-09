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
                            return $http.get('/user/get/' + params.Id);
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
        $http.get('/user/getall')
            .success(function (users) {
                $scope.users = users;
            });

        $scope.edit = function (user, $index, $event) {
            $event.stopPropagation();
            $location.path('user/edit/' + user.id);
        }

        $scope.new = function () {
            $location.path('user/new');
        }
    }]);


angular.module('admin.user').controller('UserEditController', ['$scope', '$routeParams', '$http', '$location', 'user',
    function ($scope, $routeParams, $http, $location, user) {
        $scope.user = user.data;
        $scope.redirect = function () {
            $location.path('users');
        }

        $scope.add = function () {
            $http.post('/user/insert', $scope.user)
                .success($scope.redirect);
        }

        $scope.save = function () {
            var userId = $scope.user.id
            $http.put('/user/update/' + userId, $scope.user)
                .success($scope.redirect);
        }

        $scope.remove = function () {            
            // Remove this user
            $http.delete('/user/delete/' + $scope.user.id)
                .success($scope.redirect);
        };
}]);