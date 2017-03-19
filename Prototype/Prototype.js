/**
 * Created by jamesyan on 2017/3/19.
 */

var shape=function(type,desc){
    this.type=type;
    this.draw=function(){
       console.log(desc)
    };
}

shape.prototype.getId=function(){
    return this.id;
}

shape.prototype.setId=function(id){
    this.id=id;
}
shape.prototype.getType=function(){
    return this.type;
}

shape.prototype.setType=function(type){
    this.type=type;
}

var Rectangle=function(){
    shape.call(this,"Rectangle","Inside Rectangle::draw() method");
}

Rectangle.prototype=Object.create(shape.prototype);

var Circle=function(){
    shape.call(this,"Circle","Inside Circle::draw() method");
}

Circle.prototype=Object.create(shape.prototype);

var Square=function(){
    shape.call(this,"Square","Inside Square::draw() method");
}

Square.prototype=Object.create(shape.prototype);


var ShapeCache=function(){
    this.shapeMap={};
}

ShapeCache.prototype.loadCache=function(){
    var circle=new Circle();
    circle.setId("1");
    this.shapeMap[circle.getId()]=circle;

    var square=new Square();
    square.setId("2");
    this.shapeMap[square.getId()]=square;

    var rectangle=new Rectangle();
    rectangle.setId("3");
    this.shapeMap[rectangle.getId()]=rectangle;

}

ShapeCache.prototype.getShape=function(shapeId){
    var shape=this.shapeMap[shapeId];
    return Object.create(shape);
}