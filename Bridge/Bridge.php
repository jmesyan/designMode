<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/19
 * Time: 18:58
 */
interface DrawApi
{
    public function drawCircle($radius, $x, $y);
}

class RedCircle implements DrawApi
{
    public function drawCircle($radius, $x, $y)
    {
        // TODO: Implement drawCircle() method.
        echo "Drawing Circle[ color: red, radius: " . $radius . ", x: " . $x . ", y:" . $y . "]<br/>";
    }
}

class GreenCircle implements DrawApi
{
    public function drawCircle($radius, $x, $y)
    {
        // TODO: Implement drawCircle() method.
        echo "Drawing Circle[ color: green, radius: " . $radius . ", x: " . $x . ", y:" . $y . "]<br/>";
    }
}

abstract class Shape
{
    protected $drawAPI;

    protected function __construct($drawAPI)
    {
        $this->drawAPI = $drawAPI;

    }

    public abstract function draw();
}

class Circle extends Shape
{
    private $x, $y, $radius;

    public function __construct($x, $y, $radius, $drawAPI)
    {
        parent::__construct($drawAPI);
        $this->x = $x;
        $this->y = $y;
        $this->radius = $radius;
    }

    public function draw()
    {
        $this->drawAPI->drawCircle($this->radius, $this->x, $this->y);
    }
}

//demo
$redCircle=new Circle(100,100,10,new RedCircle());
$greenCircle=new Circle(150,150,15,new GreenCircle());

$redCircle->draw();
$greenCircle->draw();