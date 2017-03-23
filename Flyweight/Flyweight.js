/**
 * Created by jamesyan on 2017/3/24.
 */


var Circle=function(color){
    this.color=color;
}

Circle.prototype.setX=function(x){
    this.x=x;
}

Circle.prototype.setY=function(y){
    this.y=y;
}

Circle.prototype.setRadius=function(radius){
    this.radius=radius;
}

Circle.prototype.draw=function(){
    console.log("Circle: Draw() [Color : ", this.color, ", x : ", this.x, ", y :", this.y, ", radius :", this.radius);
}


var ShapeFactory=function(){
    this.shapeMap={};
}
ShapeFactory.prototype.getCircle=function(color){
    if(this.shapeMap.hasOwnProperty(color)){
        return this.shapeMap[color];
    }else{
        var circle=new Circle(color);
        this.shapeMap[color]=circle;
        console.log("creating Color:"+color);
        return circle;
    }
}