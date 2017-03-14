/**
 * Created by Jmesyan on 2017/3/14.
 */
//shape
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

//color
//Red
function Red(){}
Red.prototype.fill=function(){
    console.log("Red fill() method");
}

//Green
function Green(){}
Green.prototype.fill=function(){
    console.log("Green fill() method");
}

//Blue
function Blue(){}
Blue.prototype.fill=function(){
    console.log("Blue fill() method");
}

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

shapeFactory.prototype.getColor=function(){
    return null;
}

function colorFactory(){}

colorFactory.prototype.getColor=function(colorType){
    if(colorType==null){
        return null;
    }
    colorType=colorType.toLowerCase();
    switch(colorType){
        case "red":
            return new Red();
            break;
        case "green":
            return new Green();
            break;
        case "blue":
            return new Blue();
            break;
        default:
            throw new Error("no color find!")
    }
    return null;
}

colorFactory.prototype.getShape=function(shapeType){
    return null;
}


function FactoryProducer(){}

FactoryProducer.prototype.getFactory=function(choice){
    if(choice==null){
        return null;
    }
    choice=choice.toLowerCase();
    if(choice=="shape"){
        return new shapeFactory();
    }else if(choice=="color"){
        return new colorFactory();
    }
    return null;
}

