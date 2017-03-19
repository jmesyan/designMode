/**
 * Created by jamesyan on 2017/3/19.
 */

/**
 interface DrawApi
 {
     public function drawCircle($radius, $x, $y);
 }

 type InterShape interface {
	draw()
}

 **/

var RedCircle=function(){};
RedCircle.prototype.drawCircle=function(radius,x,y){
    console.log("Drawing Circle[ color: red, radius: ",radius, ", x: ",x,", y:",y,"]\r\n");
}

var GreenCircle=function(){};

GreenCircle.prototype.drawCircle=function(radius,x,y){
    console.log("Drawing Circle[ color: green, radius: ",radius, ", x: ",x,", y:",y,"]\r\n");
}


var Shape=function(drawAPI){
    this.drawAPI=drawAPI;
}

var Circle=function(x,y,radius,drawAPI){
    Shape.call(this,drawAPI);
    this.x=x;
    this.y=y;
    this.radius=radius;
}

Circle.prototype.draw=function(){
    this.drawAPI.drawCircle(this.radius,this.x,this.y)
}

