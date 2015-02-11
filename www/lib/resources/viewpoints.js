angular.module('resources.viewpoints', ['mongodb', 'resources.votes']);

angular.module('resources.viewpoints')
.factory('Viewpoints', ['mongoResource', 'Votes', function (mongoResource, Votes) {
	var Viewpoints = mongoResource('viewpoints');

	Viewpoints.addMethod('forTopic', function(topicId) {
		return this.query({topic: topicId});
	});

	Viewpoints.vote = function(vid, up) {		
		var vote = new Votes({_id: vid, up: up});
		return vote.$update();
	};

	return Viewpoints;
}]);