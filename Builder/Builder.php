<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/15
 * Time: 20:28
 */
//设定接口
interface Item
{
    public function name();

    public function packing();

    public function price();
}

interface Packing
{
    public function Pack();
}


//创建打包实体类

class Wrapper implements Packing
{
    public function Pack()
    {
        return "Wrapper";
    }
}

class Bottle implements Packing
{
    public function Pack()
    {
        return "Bottle";
    }
}

//创建食物抽象类
abstract class Burger implements Item
{
    public function packing()
    {
        // TODO: Implement packing() method.
        return new Wrapper();
    }

    public abstract function price();
}


abstract class ColdDrink implements Item
{
    public function packing()
    {
        // TODO: Implement packing() method.
        return new Bottle();
    }

    public abstract function price();
}


//实现食物具体类
class vegBurger extends Burger
{
    public function price()
    {
        // TODO: Implement price() method.
        return 25.0;
    }

    public function name()
    {
        // TODO: Implement name() method.
        return "veg burger";
    }
}

class ChickenBurger extends Burger
{
    public function price()
    {
        // TODO: Implement price() method.
        return 50.5;
    }

    public function name()
    {
        // TODO: Implement name() method.
        return "chicken burger";
    }
}

class Coke extends ColdDrink
{
    public function price()
    {
        // TODO: Implement price() method.
        return 21.6;
    }

    public function name()
    {
        // TODO: Implement name() method.
        return "Coke";
    }
}

class Pesi extends ColdDrink
{
    public function price()
    {
        // TODO: Implement price() method.
        return 25.8;
    }

    public function name()
    {
        // TODO: Implement name() method.
        return "Pesi";
    }
}


//创建点餐类
class Meal
{
    public $items = array();
    public function addItem($item){
        array_push($this->items,$item);
    }

    public function getCost(){
        $cost=0;
        foreach($this->items as $item){
            $cost+=$item->price();
        }
        return $cost;
    }

    public function showItems(){
        foreach($this->items as $item){
            var_dump("item: ".$item->name().",Packing :".$item->packing()->pack().",Price :".$item->price());
        }
    }
}

//实验对象

class MealBuilder{
    public function prepareVegMeal(){
        $meal=new Meal();
        $meal->addItem(new vegBurger());
        $meal->addItem(new Coke());
        return $meal;
    }

    public function prepareNonVegMeal(){
        $meal=new Meal();
        $meal->addItem(new ChickenBurger());
        $meal->addItem(new Pesi());
        return $meal;
    }
}

//demo
$mealBuilder=new MealBuilder();

$vegMeal=$mealBuilder->prepareVegMeal();
var_dump("Veg Meal:<br/>");
$vegMeal->showItems();
var_dump("Total Cost:{$vegMeal->getCost()}<br/>");

$vegMeal2=$mealBuilder->prepareNonVegMeal();
var_dump("NonVeg Meal:<br/>");
$vegMeal2->showItems();
var_dump("Total Cost:{$vegMeal2->getCost()}<br/>");
