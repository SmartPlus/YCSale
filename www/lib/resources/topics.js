angular.module('resources.topics', ['mongodb']);

angular.module('resources.topics')
.factory('Topics', ['mongoResource', function (mongoResource) {
  var Topics = mongoResource('topics');
  
  Topics.prototype.isTopicOwner = function(userId) {
	return this.owner === userId;
  };

  Topics.View.paging = function(pagesize, page) {
  	return Topics.View.query({pagesize: pagesize, page: page});
  }	
  return Topics;
}]);