<template>
  <canvas style="width:100%;height:100%;" id="c">

  </canvas>
</template>

<script>
  import * as light3d from 'light3d/src/index.js'
  export default {
    name: 'l3d',
    components: {

    },
    data(){
      return {
          timer: '',
        l:null,
        count : 0,
          app:null,
          vertices:null,
      };
    },
    mounted () {
        this.l=new  light3d.Render(document.getElementById("c"));
        this.l.config();

        this.app=new light3d.Singlecolor( this.l.gl,[0.9, 0.0, 0.0, 1.0],"mvp","1.0").linkProgram().useProgram();

        this.vertices = new light3d.Polygon(5,1,"Arrays",WebGLRenderingContext.TRIANGLES).position;

        let position_b = this.l.data( this.vertices);
        this.app.setattr(WebGLRenderingContext.ARRAY_BUFFER,position_b, this.app.gl_Position,3);

        this.l.setprogram("triangle", this.app);
        this.l.lookAt([0.0, 0,2], [0, 0, 0], [0, 1, 0]);
        this.l.perspective(45,  this.l.canvas.width/ this.l.canvas.height,0.1, 100);

        console.log(new light3d.glstatus( this.l.gl).checkgl());

        this.timer = setInterval(this.render, 1000/30);
    },
      methods: {
          render(){
              this.l.clear();
              this.count++;

              var rad = (this.count % 360) * Math.PI / 180;

              this.app.initmMatrix();
              this.app.rotate(rad,[0,1,1], this.l.pMatrix, this.l.vMatrix);

              this.l.gl.clear(WebGLRenderingContext.COLOR_BUFFER_BIT);

              this.l.drawArrays(WebGLRenderingContext.TRIANGLES, this.vertices.length);
              this.l.flush();
          },

      },
      beforeDestroy() {
          clearInterval(this.timer);
      }
  }
</script>
