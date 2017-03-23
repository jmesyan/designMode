<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/23
 * Time: 22:08
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


class ShapeMaker
{
    private $circle, $rectangle, $square;

    public function __construct()
    {
        $this->circle = new Circle();
        $this->rectangle = new Rectangle();
        $this->square = new Square();
    }

    public function drawCircle(){
        $this->circle->draw();
    }

    public function drawRectangle(){
        $this->rectangle->draw();
    }

    public function drawSquare(){
        $this->square->draw();
    }
}

//demo
$shapeMaker=new ShapeMaker();
$shapeMaker->drawCircle();
$shapeMaker->drawRectangle();
$shapeMaker->drawSquare();
