<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 14:02
 */

class ChatRoom{
    public static function showMessage($user,$message){
        var_dump(date("Y-m-d H:i:s").":[{$user->getName()}]{$message}");
    }
}

class User{
    private $name;
    public function __construct($name)
    {
        $this->name=$name;
    }

    public function getName(){
        return $this->name;
    }

    public function setName($name){
        $this->name=$name;
    }

    public function sendMessage($message){
        ChatRoom::showMessage($this,$message);
    }
}


//demo
$robet=new User("Robot");
$john=new User("John");

$robet->sendMessage("Hi Join!");
$john->sendMessage("hello Robert!");