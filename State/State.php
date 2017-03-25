<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 15:18
 */

interface  State{
    public function doAction($context);
}

class startState implements State{
    public function doAction($context)
    {
        // TODO: Implement doAction() method.
        var_dump("player is in start state");
        $context->setState($this);
    }

    public function __toString()
    {
        // TODO: Implement __toString() method.
        return "start state";
    }
}

class stopState implements State{
    public function doAction($context)
    {
        // TODO: Implement doAction() method.
        var_dump("player is in stop state");
        $context->setState($this);
    }

    public function __toString()
    {
        // TODO: Implement __toString() method.
        return "stop state";
    }
}


class Context{
    private $state;
    public function __construct()
    {
        $this->state=null;
    }

    public function setState($state){
        $this->state=$state;
    }

    public function getState(){
        return $this->state;
    }
}

//demo
$context=new Context();
$startState=new startState();
$startState->doAction($context);
echo ($context->getState());

$stopState=new stopState();
$stopState->doAction($context);
echo ($context->getState());

