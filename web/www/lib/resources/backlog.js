angular.module('resources.productbacklog', ['mongodb']);
angular.module('resources.productbacklog').factory('ProductBacklog', ['mongoResource', function (mongoResource) {
  var ProductBacklog = mongoResource('productbacklog');

  ProductBacklog.forProject = function (projectId) {
    return ProductBacklog.query({projectId:projectId});
  };

  return ProductBacklog;
}]);
