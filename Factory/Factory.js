/**
 * Created by jamesyan on 2017/3/13.
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
//定义shape工厂
function shapeFactory(){}

shapeFactory.prototype.getShape=function(shapeType){
    if(shapeType==null){
        return null;
    }
    shapeType=shapeType.toLowerCase();
    switch(shapeType){
        case "rectangle":
            return new Rectangle();
            break;
        case "square":
            return new Square();
            break;
        case "circle":
            return new Circle();
            break;
        default:
            throw new Error("no graph find!")
    }
    return null;
}