<template>
    <div id="c" style="width:100%;height:100%;">

    </div>
</template>

<script>

    import echarts from 'echarts';
    import 'echarts-gl';

    export default {
        name: 'ws',
        components: {},
        data() {
            return {

                websocket: null,
                myChart:null,
                option:null,
            };
        },

        mounted() {

            this.initgl();
            this.initWebSocket();
        },

        methods: {
            initgl() {
                this.myChart = echarts.init(document.getElementById('c'));
                var symbolSize = 2.5;
                this.option = {
                    grid3D: {},
                    xAxis3D: {
                        type: 'category'
                    },
                    yAxis3D: {},
                    zAxis3D: {},
                    series: [{
                        type: 'scatter3D',
                        data: [],
                    }],
                };
            },
            initWebSocket() {


                if ('WebSocket' in window) {
                    this.websocket = new WebSocket("ws://127.0.0.1:10000/e");
                } else {
                    alert('当前浏览器 Not support websocket')
                } //连接成功建立的回调方法 websocket.onopen = function () { console.log("WebSocket连接成功")

                this.websocket.onmessage = this.websocketonmessage;
            },
            websocketonmessage: function (event) {

                let v = JSON.parse(event.data);
                console.log("+")
                this.option.series[0].data=[].concat(v);
                this.myChart.setOption(this.option);
            }
        }


    }
</script>
