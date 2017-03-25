<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 16:15
 */
abstract class AbstractCustomer
{
    protected $name;

    public abstract function isNil();

    public abstract function getName();
}

class RealCustommer extends AbstractCustomer
{
    public function __construct($name)
    {
        $this->name = $name;
    }

    public function getName()
    {
        // TODO: Implement getName() method.
        return $this->name;
    }

    public function isNil()
    {
        // TODO: Implement isNil() method.
        return false;
    }
}

class NullCustommer extends AbstractCustomer
{
    public function getName()
    {
        // TODO: Implement getName() method.
        return "not available in customer database";
    }

    public function isNil()
    {
        // TODO: Implement isNil() method.
        return true;
    }
}


class CustomerFactory
{
    public static $names = array("Rob", "Job", "Juile");

    public static function getCustomer($name)
    {

        foreach(self::$names as $n){
            if($n==$name){
                return new RealCustommer($name);
            }
        }
        return new NullCustommer();
    }
}


//demo
$customer1=CustomerFactory::getCustomer("Rob");
$customer2=CustomerFactory::getCustomer("Bob");
$customer3=CustomerFactory::getCustomer("Juile");
$customer4=CustomerFactory::getCustomer("Laura");
var_dump("Customers:");
var_dump($customer1->getName());
var_dump($customer2->getName());
var_dump($customer3->getName());
var_dump($customer4->getName());