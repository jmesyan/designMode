<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/23
 * Time: 23:22
 */
interface Shape
{
    public function draw();
}

class Circle implements Shape
{
    private $color, $x, $y, $radius;

    public function __construct($color)
    {
        $this->color = $color;
    }

    public function setX($x)
    {
        $this->x = $x;
    }


    public function setY($y)
    {
        $this->y = $y;
    }

    public function setRadius($radius)
    {
        $this->radius = $radius;
    }

    public function draw()
    {
        // TODO: Implement draw() method.
        var_dump("Circle: Draw() [Color : ".$this->color.", x : ".$this->x.", y :".$this->y.", radius :".$this->radius);
    }
}

class ShapeFactory
{
    private static $circleMap = array();

    public static function getCircle($color)
    {

        if(isset(self::$circleMap[$color])){
            $circle = self::$circleMap[$color];
        }else{
            $circle = new Circle($color);
            self::$circleMap[$color] = $circle;
            var_dump("Creating circle of color : ".$color);
        }
        return $circle;
    }

}


//demo
$colors=array("Red", "Green", "Blue", "White", "Black");
for($i=0; $i < 20; ++$i) {
    $randColor=$colors[array_rand($colors)];
    $circle=ShapeFactory::getCircle($randColor);
    $circle->setX(rand(50,100));
    $circle->setY(rand(50,100));
    $circle->setRadius(100);
    $circle->draw();
}