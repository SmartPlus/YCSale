angular.module('admin', [
'ngRoute',
'ngResource',
'ui.bootstrap',
'services.breadcrumbs',
'services.httpRequestTracker',
'security',
'admin.user'
]);

angular.module('admin').run(['security',
    function (security) {
        // Get the current user when the application starts
        // (in case they are still logged in from a previous session)
        security.requestCurrentUser();
}]);

angular.module('admin').controller('AdminController', ['$scope',
    function ($scope) {

        $scope.tpls = {
            header: '/admin/tpl/header.html'
        };

}]);


angular.module('admin').controller('HeaderCtrl', ['$scope', '$location', '$route', 'breadcrumbs', 'httpRequestTracker',
  function ($scope, $location, $route, breadcrumbs, httpRequestTracker) {

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
}]);

angular.module('admin').directive('ngConfirmClick', [
  function(){
    return {
      priority: -1,
      restrict: 'A',
      link: function(scope, element, attrs){
        element.bind('click', function(e){
          var message = attrs.ngConfirmClick;
          if(message && !confirm(message)){
            e.stopImmediatePropagation();
            e.preventDefault();
          }
        });
      }
    }
  }
]);