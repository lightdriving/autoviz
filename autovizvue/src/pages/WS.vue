<template>
    <canvas style="width:100%;height:100%;" id="c">

    </canvas>
</template>

<script>
    import * as light3d from 'light3d/src/index.js'

    export default {
        name: 'ws',
        components: {},
        data() {
            return {

                websocket: null,
                l: null,
                app: null,
                vertices: [],
                position_b:null,
            };
        },

        mounted() {
            this.l = new light3d.Render(document.getElementById("c"));
            this.l.config();

            this.app = new light3d.Singlecolor(this.l.gl, [0.9, 0.0, 0.0, 1.0], "mvp", "3.1").linkProgram().useProgram();

            let vertices = new light3d.Polygon(1000, 10, "Arrays", WebGLRenderingContext.POINTS).position;

            this.position_b = this.l.data(vertices);
            this.app.setattr(WebGLRenderingContext.ARRAY_BUFFER,  this.position_b, this.app.gl_Position, 3);

            this.l.setprogram("triangle", this.app);
            this.l.lookAt([0.0, 0,30], [0, 0, 0], [0, 1, 0]);
            this.l.perspective(45, this.l.canvas.width / this.l.canvas.height, 0.1, 100);
            this.l.toworld();
            this.initWebSocket();
        },
        methods: {
            initWebSocket(){
                console.log(new light3d.glstatus(this.l.gl).checkgl());

                if ('WebSocket' in window) {
                    this.websocket = new WebSocket("ws://127.0.0.1:10000/ws");
                } else {
                    alert('当前浏览器 Not support websocket')
                } //连接成功建立的回调方法 websocket.onopen = function () { console.log("WebSocket连接成功")

                this.websocket.onmessage = this.websocketonmessage;
            },
            websocketonmessage:function(event) {

                let v=JSON.parse(event.data);

                this.app.initmMatrix();
                this.app.rotate(0,[0,0,1], this.l.pMatrix,this.l.vMatrix);

                this.l.clear();
                this.l.updatedata(WebGLRenderingContext.ARRAY_BUFFER, this.position_b,v)
                this.l.gl.bindBuffer(WebGLRenderingContext.ARRAY_BUFFER, this.position_b);
                this.l.drawArrays(WebGLRenderingContext.POINTS, v.length/3);
                this.l.flush();
            }
        }


    }
</script>
