angular.module('resources.sprints', ['mongodb']);
angular.module('resources.sprints').factory('Sprints', ['mongoResource', function (mongoResource) {

  var Sprints = mongoResource('sprints');
  Sprints.forProject = function (projectId) {
    return Sprints.query({projectId:projectId});
  };
  return Sprints;
}]);