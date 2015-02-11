angular.module('directives.crud.buttons', [])

.directive('crudButtons', function () {
  return {
    restrict:'E',
    replace:true,
    template:
      '<div>' +
      '  <button type="button" class="btn btn-default btn-primary save" ng-disabled="!canSave()" ng-click="save()">Save</button>' +
      '  <button type="button" class="btn btn-default btn-warning revert" ng-click="revertChanges()" ng-disabled="!canRevert()">Revert changes</button>'+
      '  <button type="button" class="btn btn-default btn-danger remove" ng-click="remove()" ng-show="canRemove()">Remove</button>'+
      '</div>'
  };
});