<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/14
 * Time: 22:18
 */

class Singleton{
    private static $instance;

    private function __construct()
    {

    }

    public static function getInstance(){

        if(!self::$instance instanceof self){
            self::$instance=new self();
            echo "11111<br/>";
        }
        echo  "22222<br/>";
        return self::$instance;
    }

    public function showMessage(){
        echo "hello world<br/>";
    }

}

//test

$object=Singleton::getInstance();
$object->showMessage();

$object2=Singleton::getInstance();
$object2->showMessage();