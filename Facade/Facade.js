/**
 * Created by jamesyan on 2017/3/23.
 */
//Rectangle
function Rectangle(){}
Rectangle.prototype.draw=function(){
    console.log("Rectangle draw() method");
}

//Square
function Square(){}
Square.prototype.draw=function(){
    console.log("Square draw() method");
}

//Square
function Circle(){}
Circle.prototype.draw=function(){
    console.log("Circle draw() method");
}

var ShapeMaker=function(circle,rectangle,square){
    this.circle=circle;
    this.rectangle=rectangle;
    this.square=square;
}

ShapeMaker.prototype.drawCircle=function () {
    this.circle.draw()
}

ShapeMaker.prototype.drawRectangle=function () {
    this.rectangle.draw()
}

ShapeMaker.prototype.drawSquare=function () {
    this.square.draw()
}