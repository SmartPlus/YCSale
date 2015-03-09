angular.module('chart.details', []);

angular.module('chart.details').config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.when('/details', {
            templateUrl: '/tpl/chart/details.html',
            controller: 'DetailsChart'
        });
}]);

angular.module('chart.details').controller('DetailsChart', ['$scope', '$resource', '$q',
    function ($scope, $resource, $q) {
        var Graph = $resource('/data');

        $scope.fetchData = function () {
            return Graph.query({
                from: $scope.from._date,
                to: $scope.to._date,
                period: $scope.period
            }).$promise;
        }

        $scope.showChart = function (data) {
            chart('#chart', $scope.config, data);
        };

        $scope.updateChart = function () {            
            $scope.showChart([]);
            $scope.fetchData().then(function (data) {                
                $scope.showChart(data);
            });
        }

        $scope.config = {
            'one': {
                name: 'One User - One Account',
                type: 'bar',
                on: true
            },
            'unique': {
                name: 'Unique',
                type: 'bar'
            },
            'total': {
                name: 'Total',
                type: 'bar',
                on: false
            },
            'average': {
                name: 'Ratio: unique/total',
                type: 'line',
                axis: 'y2'
            }
        };

        $scope.formatDate = function (data) {
            data._date = moment(data.date).format('YYYY-MM-DD'); // for query
            data.date = moment(data.date).format('dddd, MMMM DD, YYYY');
        }

        $scope.today = function () {
            $scope.from = {
                date: $scope.formatDate({})
            };
            $scope.to = {
                date: $scope.formatDate({})
            };
        };

        $scope.today();

        $scope.open = function ($event, type) {
            $event.preventDefault();
            $event.stopPropagation();

            type.opened = true;
        };

        $scope.periods = [1, 2, 3, 7, 12, 16];
        $scope.period = 1;

        $scope.showChart([]);

}]);

function chart(id, conf, data) {
    var chart = c3.generate({
        bindto: id,
        data: {
            json: data,
            keys: {
                x: 'date',
                value: _.keys(conf)
            },
            axes: _.omit(_.mapValues(conf, 'axis'), _.isUndefined),
            names: _.omit(_.mapValues(conf, 'name'), _.isUndefined),
            types: _.omit(_.mapValues(conf, 'type'), _.isUndefined)
        },
        axis: {
            x: {
                label: {
                    text: 'Date',
                    position: 'middle'
                },
                type: 'timeseries',
                tick: {
                    format: '%Y-%m-%d'
                }
            },

            y: {
                label: {
                    text: 'Count',
                    position: 'middle'
                }
            },

            y2: {
                show: true,
                min: 0,
                max: 1
            }
        },
        grid: {
            x: {
                show: true
            },
            y: {
                show: true
            }
        }
    });
}