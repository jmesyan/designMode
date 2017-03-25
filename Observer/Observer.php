<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 14:34
 */

class Subject{
    private $state;
    private $observers;
    public function getState(){
        return $this->state;
    }

    public function setState($state){
        $this->state=$state;
        $this->notifyUsers();

    }

    public function attach($observer){
        $this->observers[]=$observer;
    }


    public function notifyUsers(){
        foreach($this->observers as $observer){
            $observer->update();
        }
    }
}


abstract  class Observer{
    protected $subject;
    public abstract function update();
}


class BinaryObserver extends Observer{
    public function __construct($subject)
    {
        $this->subject=$subject;
        $this->subject->attach($this);
    }

    public function update()
    {
        // TODO: Implement update() method.
        var_dump("Binary String: ".$this->subject->getState());
    }
}

class OctalObserver extends Observer{
    public function __construct($subject)
    {
        $this->subject=$subject;
        $this->subject->attach($this);
    }

    public function update()
    {
        // TODO: Implement update() method.
        var_dump("Octal String: ".$this->subject->getState());
    }
}

class HexaObserver extends Observer{
    public function __construct($subject)
    {
        $this->subject=$subject;
        $this->subject->attach($this);
    }

    public function update()
    {
        // TODO: Implement update() method.
        var_dump("Hexa String: ".$this->subject->getState());
    }
}

//demo

$subject=new Subject();

new HexaObserver($subject);
new OctalObserver($subject);
new BinaryObserver($subject);

var_dump("First state change: 15");
$subject->setState(15);

var_dump("Secend state change: 10");
$subject->setState(10);