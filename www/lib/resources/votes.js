angular.module('resources.votes', ['mongodb']);

angular.module('resources.votes')
.factory('Votes', ['mongoResource', function (mongoResource) {
  var Votes = mongoResource('votes');
  return Votes;
}]);