<?php

/**
 * Created by PhpStorm.
 * User: Jmesyan
 * Date: 2017/3/14
 * Time: 9:29
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

interface color
{
    public function fill();
}

class Red implements color
{
    public function fill()
    {
        echo "red fill() method<br/>";
    }
}

class Green implements color
{
    public function fill()
    {
        echo "Green fill() method<br/>";
    }
}

class Blue implements color
{
    public function fill()
    {
        echo "Blue fill() method<br/>";
    }
}

abstract class AbstractFactory
{
    abstract function getColor($color);

    abstract function getShape($shapeType);
}

class ShapeFactory extends AbstractFactory
{
    public function getShape($shapeType)
    {
        if ($shapeType == null) {
            return null;
        }
        $shapeType = strtolower($shapeType);
        if ($shapeType == "rectangle") {
            return new Rectangle();
        } elseif ($shapeType == "square") {
            return new Square();
        } elseif ($shapeType == "circle") {
            return new Circle();
        }
        return null;
    }

    public function getColor($color)
    {
        // TODO: Implement getColor() method.
        return null;
    }
}


class ColorFactory extends AbstractFactory
{
    public function getColor($color)
    {
        if ($color == null) {
            return null;
        }
        $color = strtolower($color);
        if ($color == "red") {
            return new Red();
        } elseif ($color == "green") {
            return new Green();
        } elseif ($color == "Blue") {
            return new Blue();
        }
        return null;
    }

    public function getShape($shapeType)
    {
        return null;
    }
}

class FactoryProducer
{
    public static function getFactory($choice)
    {
        $choice=strtolower($choice);
        if($choice=="shape"){
            return new shapeFactory();
        }else if($choice=="color"){
            return new ColorFactory();
        }
        return null;
    }
}

//demo
$shapeFactory=FactoryProducer::getFactory("shape");
$shape1=$shapeFactory->getShape("circle");
$shape1->draw();
$shape2=$shapeFactory->getShape("rectangle");
$shape2->draw();

$colorFactory=FactoryProducer::getFactory("color");
$color1=$colorFactory->getColor("green");
$color1->fill();
$color2=$colorFactory->getColor("red");
$color2->fill();