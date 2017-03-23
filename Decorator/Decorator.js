/**
 * Created by Jmesyan on 2017/3/23.
 */

/**
 interface Shape{
    public function draw();
}
 **/

var Circle=function(){
    this.draw=function(){
        console.log("Shape:Circle");
    }
};

var Rectangle=function(){
    this.draw=function(){
        console.log("Shape:Rectangle");
    }
};

var ShapeDector=function(dectorShape) {
    this.dectorShape = dectorShape;
};

ShapeDector.prototype.draw=function() {
   this.dectorShape.draw();
};

var RedShapeDector=function(dectorShape){
    ShapeDector.call(this,dectorShape);
}

RedShapeDector.prototype=Object.create(ShapeDector.prototype)

RedShapeDector.prototype.draw=function() {
    this.dectorShape.draw();
    this.setRedBorder();

};

RedShapeDector.prototype.setRedBorder=function(){
    console.log("Border Color:Red");
}


