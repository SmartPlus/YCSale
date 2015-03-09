angular.module('metric.table', ['ui.grid', 'ui.grid.resizeColumns']);

angular.module('metric.table').config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.when('/table/:game_code/:type', {
            templateUrl: '/tpl/metric/table.html',
            controller: 'MetricTableController'
        });
        $routeProvider.when('/table/:game_code', {
            templateUrl: '/tpl/metric/table_list.html',
            controller: 'MetricTableListController'
        })
}]);

angular.module('metric.table').controller('MetricTableListController', ['$scope', '$routeParams', '$http',
    function ($scope, $routeParams, $http) {
        $scope.types = ['metric', 'summary'];
        $scope.game_code = $routeParams.game_code;

        $scope.hrefs = {};
        _.forEach($scope.types, function (v) {
            $scope.hrefs[v] = '#/table/' + $scope.game_code + '/' + v;
        })
    }]);

angular.module('metric.table').controller('MetricTableController', ['$scope', '$routeParams', '$http',
    function ($scope, $routeParams, $http) {

        $scope.detailColumns = {};

        $scope.detailMenu = function (name, all) {
            return {
                title: all ? 'ALL' : 'Detail ' + name,
                icon: all ? 'ui-grid-icon-ok' : 'ui-grid-icon-info-circled',
                shown: function () {
                    return ($scope.detailColumns[name] === all);
                },
                action: function ($event) {
                    if (all) delete $scope.detailColumns[name];
                    else $scope.detailColumns[name] = true;
                    $scope.updateTable();
                }                
            }
        }

        $http.get('/grid_state', {
            params: _.extend({}, $routeParams)
        }).success(function (data) {
            _.forEach(data.columnDefs, function (c) {
                if (c.hasDetail) {
                    c.enableFiltering = true;
                    c.menuItems = (c.menuItems || []).concat([                   
                        $scope.detailMenu(c.field),
                        $scope.detailMenu(c.field, true)
                    ]);
                }
            })
            _.extend($scope.options, data);
            $scope.gridApi.grid.refresh();            
        }).error(function (err) {
            alert(err);
        });


        $scope.options = {           
            enableGridMenu: true,
            gridMenuShowHideColumns: true,
            enableColumnResizing: true,
            enableColumnMoving: true,
            enableFiltering: true,
            data: $scope.table, 
            onRegisterApi: function (gridApi) {
                $scope.gridApi = gridApi;
            }
        }



        $scope.time_units = ['daily', 'weekly'];
        $scope.time_unit = 'daily';
        $scope.firstRun = true;
        $scope.isCollapsed = true;

        $scope.updateTable = function () {
            var params = {};
            _.forEach($scope.detailColumns, function (v, k) {
                if (v) {
                    params['all_' + k.toLowerCase()] = true;
                }
            })

            params.time_unit = $scope.time_unit;
            $http.get('/data', {
                params: _.extend({}, $routeParams, params, {
                    from: $scope.from._date,
                    to: $scope.to._date
                })
            }).success(function (data) {
                $scope.options.data = data;
            }).error(function (data) {
                alert(data);
            });
        }

        $scope.formatDate = function (data) {
            data._date = moment(data.date).format('YYYY-MM-DD'); // for query
            data.date = moment(data.date).format('dddd, MMMM DD, YYYY');
        }

        $scope.today = function () {
            $scope.from = {
                date: moment('2014-01-01')
            };
            $scope.formatDate($scope.from);
            $scope.to = {
                date: moment()
            };
            $scope.formatDate($scope.to);
        };

        $scope.today();

        $scope.open = function ($event, type) {
            $event.preventDefault();
            $event.stopPropagation();

            type.opened = true;
        };

}]);