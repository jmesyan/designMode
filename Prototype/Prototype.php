<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/18
 * Time: 18:03
 */
abstract class Shape
{
    private $id;
    protected $type;
    public  $time;

    abstract function draw();

    public function getType()
    {
        return $this->type;
    }

    public function getId()
{
    return $this->id;
}

    public function setId($id)
    {
        $this->id = $id;
    }

    public function getTime()
    {
        return $this->time->format('Y-m-d H:i:s');
    }

    public function setTime($time)
    {
        $this->time = $time;
    }

    public function setType($type){
        $this->type=$type;
    }

    public function __clone()
    {
        // TODO: Implement __clone() method.
        $this->time=clone $this->time;//如果没有这句，传递的time类型浅拷贝中后面对象改变前面对象也会改变

    }


}


class Rectangle extends shape
{
    public function __construct()
    {
        $this->type = "Rectangle";
    }

    public function draw()
    {
        echo "Inside Rectangle::draw() method";
    }
}

class Square extends shape
{
    public function __construct()
    {
        $this->type = "Square";
    }

    public function draw()
    {
        echo "Inside Square::draw() method";
    }
}

class Circle extends shape
{
    public function __construct()
    {
        $this->type = "Circle ";
    }

    public function draw()
    {
        echo "Inside Circle ::draw() method";
    }
}

class ShapeCache
{
    private static $shapeMap = array();

    public static function getShape($shapeId)
    {
        $cachedShape = self::$shapeMap[$shapeId];
        return clone $cachedShape;
    }

    public static function loadCache()
    {
        $circle=new Circle();
        $circle->setId(1);
        $circle->setTime(new DateTime("2014-07-05", new DateTimeZone("UTC")));
        self::$shapeMap[$circle->getId()]=$circle;

        $square=new Square();
        $square->setId(2);
        $square->setTime(new DateTime("2014-07-06", new DateTimeZone("UTC")));
        self::$shapeMap[$square->getId()]=$square;

        $rectangle=new Rectangle();
        $rectangle->setId(3);
        $rectangle->setTime(new DateTime("2014-07-07", new DateTimeZone("UTC")));
        self::$shapeMap[$rectangle->getId()]=$rectangle;
    }
}


//demo
$shapeCache=new ShapeCache();
$shapeCache::loadCache();

$cloneShape1=$shapeCache::getShape(1);
var_dump("Shape1 :".$cloneShape1->getType());
var_dump("Shape1 Time :".$cloneShape1->getTime());

$cloneShape2=$shapeCache::getShape(2);
var_dump("Shape2 :".$cloneShape2->getType());

$cloneShape3=$shapeCache::getShape(3);
var_dump("Shape3 :".$cloneShape3->getType());

$cloneShape4=$shapeCache::getShape(1);
$cloneShape4->time->add(new DateInterval('P10D'));
var_dump("Shape4 Time :".$cloneShape4->getTime());
var_dump("Shape1 Time :".$cloneShape1->getTime());


