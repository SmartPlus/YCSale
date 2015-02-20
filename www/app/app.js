//  To add a model
// 1. Add link to script in html 
// 2. Add module to app.js
// 3. Link to model in header
// 4. Model Controller
// 5. Fields


angular.module('app', [
'ngRoute',
'ngResource',
'ui.bootstrap',
'services.breadcrumbs',
'services.httpRequestTracker',
'security',
'notyModule',
'app.customer',
'app.course'
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
            header: '/app/tpl/header.html'
        };

}]);

angular.module('app').controller('HeaderCtrl', ['$scope', '$location', '$route', '$http', 'breadcrumbs', 'httpRequestTracker', 'security',
  function ($scope, $location, $route, $http, breadcrumbs, httpRequestTracker, security) {

        $scope.location = $location;
        $scope.breadcrumbs = breadcrumbs;

        $scope.home = function () {
            $location.path('/');
        };

        $scope.isNavbarActive = function (navBarPath) {

            return navBarPath === breadcrumbs.getFirst().name;
        };

        $scope.hasPendingRequests = function () {
            return httpRequestTracker.hasPendingRequests();
        };
        $scope.security = security;
}]);