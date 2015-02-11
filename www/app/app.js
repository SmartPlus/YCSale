angular.module('app', [
'ngRoute',
'ngResource',
'ui.bootstrap',
'services.breadcrumbs',
'services.httpRequestTracker',
'security',
'notyModule',
'metric.table'
]);

angular.module('app').run(['security',
    function (security) {

        // Get the current user when the application starts
        // (in case they are still logged in from a previous session)
        security.requestCurrentUser().then(function (user) {
            if (!user) security.showLogin();
        })
}]);

angular.module('app').controller('AppCtrl', ['$scope', 'noty',
    function ($scope, noty) {
        $scope.noty = noty;
        $scope.tpls = {
            header: '/tpl/header.html'
        };

}]);

angular.module('app').controller('HeaderCtrl', ['$scope', '$location', '$route', '$http', 'breadcrumbs', 'httpRequestTracker', 'security',
  function ($scope, $location, $route, $http, breadcrumbs, httpRequestTracker, security) {


        $scope.$watch(function () {
            return security.currentUser;
        }, function (currentUser) {
            if (currentUser) {
                $http.get('/game_code')
                    .success(function (data) {
                        $scope.game_codes = data;
                    })
                    .error(function (data) {
                        console.warn(arguments);
                    })
            }
        });



        $scope.location = $location;
        $scope.breadcrumbs = breadcrumbs;

        $scope.home = function () {
            $location.path('/');
        };

        $scope.isNavbarActive = function (navBarPath) {
            console.log(navBarPath);
            return navBarPath === breadcrumbs.getFirst().name;
        };

        $scope.hasPendingRequests = function () {
            return httpRequestTracker.hasPendingRequests();
        };
        $scope.security = security;
}]);