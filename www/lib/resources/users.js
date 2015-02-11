angular.module('resources.users', ['mongodb']);
angular.module('resources.users').factory('Users', ['mongoResource', function (mongoResource) {

  var userResource = mongoResource('users');
  userResource.prototype.getFullName = function () {
    return this.lastName + " " + this.firstName + " (" + this.email + ")";
  };

  return userResource;
}]);

angular.module('resources.users').factory('UserCache', ['Users', 'security', function (Users, security) {
	var users = {};
	var UserCache = {};
	var waiting = {};

	function prepare (arr, key) {
		var newWaiting = [];
		arr.forEach(function(d) {
		  var id = d[key];
		  if((users[id] || waiting[id]) === undefined) { 		  	
		  	waiting[id] = [];
		    newWaiting.push(id);
		  }
		});    
		return newWaiting;    
	}

	function doCallback(u) {
		var id = u._id;
		var tasks = waiting[id];
		for(var i = tasks.length - 1; i >= 0; i--) {
			tasks[i](u);			
		}
		delete waiting[id];
		users[id] = u;
	}

	function saveCache (arr) {
		arr.forEach(doCallback);
		return users;
	}

	// maintain a compatible promise-based API

	UserCache.fetch = function (arr, key) {
		var queue = prepare(arr, key);
		if (queue.length > 0) {
			return Users.View.getByIds(queue).then(function(docs) {
		  		return saveCache(docs);		  		
			})  	
		}
		else return users;		
	}

	UserCache.get = function (id, callback) {
		if (!callback) { return users[id]; }
		if (users[id]) { callback(users[id]); }
		else if (waiting[id]) { waiting[id].push(callback); }
		else {
			waiting[id] = [callback];
			Users.View.getByIds([id]).then(function(docs){
				saveCache(docs);
			})
		}
	};

	security.onLoginSuccess(function(currentUser) {
		UserCache.me = currentUser;
		users[currentUser._id] = currentUser;
	});
	
	return UserCache;
}]);
