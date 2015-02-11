(function () {

    function crudRouteProvider($routeProvider) {

        // This $get noop is because at the moment in AngularJS "providers" must provide something
        // via a $get method.
        // When AngularJS has "provider helpers" then this will go away!
        this.$get = angular.noop;

        // Again, if AngularJS had "provider helpers" we might be able to return `routesFor()` as the
        // crudRouteProvider itself.  Then we would have a much cleaner syntax and not have to do stuff
        // like:
        //
        // ```
        // myMod.config(function(crudRouteProvider) {
        //   var routeProvider = crudRouteProvider.routesFor('MyBook', '/myApp');
        // });
        // ```
        //
        // but instead have something like:
        //
        //
        // ```
        // myMod.config(function(crudRouteProvider) {
        //   var routeProvider = crudRouteProvider('MyBook', '/myApp');
        // });
        // ```
        //
        // In any case, the point is that this function is the key part of this "provider helper".
        // We use it to create routes for CRUD operations.  We give it some basic information about
        // the resource and the urls then it it returns our own special routeProvider.
        this.routesFor = function (resourceName, urlPrefix, routePrefix) {
            var baseUrl = resourceName.toLowerCase();
            var baseRoute = '/' + resourceName.toLowerCase();
            routePrefix = routePrefix || urlPrefix;

            // Prepend the urlPrefix if available.
            if (angular.isString(urlPrefix) && urlPrefix !== '') {
                baseUrl = urlPrefix + '/' + baseUrl;
            }

            // Prepend the routePrefix if it was provided;
            if (routePrefix !== null && routePrefix !== undefined && routePrefix !== '') {
                baseRoute = '/' + routePrefix + baseRoute;
            }

            // Create the templateUrl for a route to our resource that does the specified operation.
            var templateUrl = function (operation) {
                return baseUrl + '/' + resourceName.toLowerCase() + '-' + operation.toLowerCase() + '.tpl.html';
            };
            // Create the controller name for a route to our resource that does the specified operation.
            var controllerName = function (operation) {
                return resourceName + operation + 'Ctrl';
            };

            // This is the object that our `routesFor()` function returns.  It decorates `$routeProvider`,
            // delegating the `when()` and `otherwise()` functions but also exposing some new functions for
            // creating CRUD routes.  Specifically we have `whenList(), `whenNew()` and `whenEdit()`.
            var routeBuilder = {
                // Create a route that will handle showing a list of items
                whenList: function (resolveFns) {
                    routeBuilder.when(baseRoute, {
                        templateUrl: templateUrl('List'),
                        controller: controllerName('List'),
                        resolve: resolveFns
                    });
                    return routeBuilder;
                },
                // Create a route that will handle creating a new item
                whenNew: function (resolveFns) {
                    routeBuilder.when(baseRoute + '/new', {
                        templateUrl: templateUrl('Edit'),
                        controller: controllerName('Edit'),
                        resolve: resolveFns
                    });
                    return routeBuilder;
                },
                // Create a route that will handle editing an existing item
                whenEdit: function (resolveFns) {
                    routeBuilder.when(baseRoute + '/:itemId', {
                        templateUrl: templateUrl('Edit'),
                        controller: controllerName('Edit'),
                        resolve: resolveFns
                    });
                    return routeBuilder;
                },
                // Pass-through to `$routeProvider.when()`
                when: function (path, route) {
                    $routeProvider.when(path, route);
                    return routeBuilder;
                },
                // Pass-through to `$routeProvider.otherwise()`
                otherwise: function (params) {
                    $routeProvider.otherwise(params);
                    return routeBuilder;
                },
                // Access to the core $routeProvider.
                $routeProvider: $routeProvider
            };
            return routeBuilder;
        };
    }
    // Currently, v1.0.3, AngularJS does not provide annotation style dependencies in providers so,
    // we add our injection dependencies using the $inject form
    crudRouteProvider.$inject = ['$routeProvider'];

    // Create our provider - it would be nice to be able to do something like this instead:
    //
    // ```
    // angular.module('services.crudRouteProvider', [])
    //   .configHelper('crudRouteProvider', ['$routeProvider, crudRouteProvider]);
    // ```
    // Then we could dispense with the $get, the $inject and the closure wrapper around all this.
    angular.module('services.crudRouteProvider', ['ngRoute']).provider('crudRoute', crudRouteProvider);
})();


angular.module('services.crud', ['services.crudRouteProvider']);
angular.module('services.crud').factory('crudEditMethods', function () {

    return function (itemName, item, formName, successcb, errorcb) {

        var mixin = {};

        mixin[itemName] = item;
        mixin[itemName + 'Copy'] = angular.copy(item);

        mixin.save = function () {
            this[itemName].$saveOrUpdate(successcb, successcb, errorcb, errorcb);
        };

        mixin.canSave = function () {
            return this[formName].$valid && !angular.equals(this[itemName], this[itemName + 'Copy']);
        };

        mixin.revertChanges = function () {
            this[itemName] = angular.copy(this[itemName + 'Copy']);
        };

        mixin.canRevert = function () {
            return !angular.equals(this[itemName], this[itemName + 'Copy']);
        };

        mixin.remove = function () {
            if (this[itemName].$id()) {
                this[itemName].$remove(successcb, errorcb);
            } else {
                successcb();
            }
        };

        mixin.canRemove = function () {
            return item.$id();
        };

        /**
         * Get the CSS classes for this item, to be used by the ng-class directive
         * @param {string} fieldName The name of the field on the form, for which we want to get the CSS classes
         * @return {object} A hash where each key is a CSS class and the corresponding value is true if the class is to be applied.
         */
        mixin.getCssClasses = function (fieldName) {
            var ngModelController = this[formName][fieldName];
            return {
                error: ngModelController.$invalid && ngModelController.$dirty,
                success: ngModelController.$valid && ngModelController.$dirty
            };
        };

        /**
         * Whether to show an error message for the specified error
         * @param {string} fieldName The name of the field on the form, of which we want to know whether to show the error
         * @param  {string} error - The name of the error as given by a validation directive
         * @return {Boolean} true if the error should be shown
         */
        mixin.showError = function (fieldName, error) {
            return this[formName][fieldName].$error[error];
        };

        return mixin;
    };
});

angular.module('services.crud').factory('crudListMethods', ['$location',
    function ($location) {

        return function (pathPrefix) {

            var mixin = {};

            mixin['new'] = function () {
                $location.path(pathPrefix + '/new');
            };

            mixin['edit'] = function (itemId) {
                $location.path(pathPrefix + '/' + itemId);
            };

            return mixin;
        };
}]);