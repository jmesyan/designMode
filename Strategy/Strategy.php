<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 17:31
 */

interface Strategy {
    public function doOperation($num1,$num2);
}


class OperationAdd implements Strategy{

    public function doOperation($num1, $num2)
    {
        // TODO: Implement doOperation() method.
        return $num1+$num2;
    }
}



class OperationSubstract implements Strategy{

    public function doOperation($num1, $num2)
    {
        // TODO: Implement doOperation() method.
        return $num1-$num2;
    }
}


class OperationMultiply implements Strategy{

    public function doOperation($num1, $num2)
    {
        // TODO: Implement doOperation() method.
        return $num1*$num2;
    }
}

class Context{
    private $strategy;
    public function __construct($strategy)
    {
        $this->strategy=$strategy;
    }

    public function executeStrategy($num1,$num2){
        return $this->strategy->doOperation($num1,$num2);
    }
}

//demo
$context=new Context(new OperationAdd());
var_dump("10+5=".$context->executeStrategy(10,5));

$context=new Context(new OperationSubstract());
var_dump("10-5=".$context->executeStrategy(10,5));

$context=new Context(new OperationMultiply());
var_dump("10*5=".$context->executeStrategy(10,5));



