<?php

/**
 * Created by PhpStorm.
 * User: 24560
 * Date: 2017/3/10
 * Time: 15:39
 */
interface shape
{
    public function draw();
}

class Rectangle implements shape
{
    public function draw()
    {
        echo "Rectangle draw() method<br/>";
    }
}

class Square implements shape
{
    public function draw()
    {
        echo "Square draw() method<br/>";
    }
}

class Circle implements shape
{
    public function draw()
    {
        echo "Circle draw() method<br/>";
    }
}

class shapeFactory
{
    public function getShape($shapeType)
    {
        if($shapeType==null){
            return null;
        }
        $shapeType=strtolower($shapeType);
        if($shapeType=="rectangle"){
            return new Rectangle();
        }elseif($shapeType=="square"){
            return new Square();
        }elseif($shapeType=="circle"){
            return new Circle();
        }
        return null;
    }
}

class FactoryPatternDemo{
   public function demo(){
       $shapeFactory=new shapeFactory();
       $shape1=$shapeFactory->getShape("Rectangle");
       $shape1->draw();
       $shape2=$shapeFactory->getShape("Circle");
       $shape2->draw();
       $shape3=$shapeFactory->getShape("Square");
       $shape3->draw();
   }
}

(new FactoryPatternDemo())->demo();