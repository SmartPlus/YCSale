angular.module('resources.categories', ['mongodb']);
angular.module('resources.categories').factory('Categories', ['mongoResource', function (mongoResource) {
  return mongoResource('categories');
}]);
