angular.module('app.course', ['services.crud']);

angular.module('app.course').config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.when('/course/:action/:Id?', {
            templateUrl: '/app/course/course_edit.html',
            controller: 'CourseEditController',
            resolve: {
                course: ['$route', '$http',
                    function ($route, $http) {
                        var params = $route.current.params;
                        if (params.action === 'new') {
                            return {};
                        } else {
                            return $http.get('/course/' + params.Id);
                        }
                }]
            }
        });
        $routeProvider.when('/courses', {
            templateUrl: '/app/course/course_list.html',
            controller: 'CourseListController'
        })
}]);

angular.module('app.course').controller('CourseListController', ['$scope', '$routeParams', '$http', '$q', '$location',
    function ($scope, $routeParams, $http, $q, $location) {
        $http.get('/courses')
            .success(function (courses) {
                $scope.courses = courses;
            });

        $scope.remove = function (course, $index, $event) {
            // Don't let the click bubble up to the ng-click on the enclosing div, which will try to trigger
            // an edit of this item.
            $event.stopPropagation();

            // Remove this course
            $http.delete('/course/' + course._id).success(function () {
                // It is gone from the DB so we can remove it from the local list too
                $scope.courses.splice($index, 1);
                console.log('crud.course.remove.success', course._id);
            }).error(function (e) {
                console.log('crud.course.remove.error', course._id, e);
            });
        };

        $scope.edit = function (course) {
            $location.path('course/edit/' + course._id);
        }

        $scope.new = function () {
            $location.path('course/new');
        }
    }]);


angular.module('app.course').controller('CourseEditController', ['$scope', '$routeParams', '$http', '$location', 'course',
    function ($scope, $routeParams, $http, $location, course) {
        $scope.course = course.data;

        $scope.new = function () {
            $http.post('/course', $scope.course)
                .success(function () {
                    console.log('course.add.success');
                    $location.path('courses');
                })
                .error(function (e) {
                    console.log('course.add.error', e);
                })
        }

        $scope.edit = function () {
            var courseId = $scope.course._id
            $http.put('/course/' + courseId, $scope.course)
                .success(function () {
                    console.log('course.edit.success', courseId);
                    $location.path('courses');
                }).error(function (e) {
                    console.log('course.edit.error', courseId, e);
                })
        }

        $scope.save = function () {
            if ($scope.course._id === undefined) {
                $scope.new();
            } else {
                $scope.edit();
            }
        }

}]);