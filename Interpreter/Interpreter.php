<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 9:56
 */

interface Expression{
    public function interpret($context);
}

class TerminalExpression implements Expression{
    private $data;
    public function __construct($data)
    {
        $this->data=$data;
    }

    public function interpret($context)
    {
        // TODO: Implement interpret() method.
        if(strpos($context,$this->data)===false){
            return false;
        }else{
            return true;
        }

    }
}


class AndExpression implements Expression{
    private $exp1,$exp2;
    public function __construct($exp1,$exp2)
    {
        $this->exp1=$exp1;
        $this->exp2=$exp2;
    }

    public function interpret($context)
    {
        // TODO: Implement interpret() method.
        return $this->exp1->interpret($context)&&$this->exp2->interpret($context);
    }
}


class OrExpression implements Expression{
    private $exp1,$exp2;
    public function __construct($exp1,$exp2)
    {
        $this->exp1=$exp1;
        $this->exp2=$exp2;
    }

    public function interpret($context)
    {
        // TODO: Implement interpret() method.
        return $this->exp1->interpret($context)||$this->exp2->interpret($context);
    }
}

//demo
//规则：Robert 和 John 是男性
$robert=new TerminalExpression("Robert");
$join=new TerminalExpression("Join");
$isMale=new OrExpression($robert,$join);

$julie = new TerminalExpression("Julie");
$married = new TerminalExpression("Married");
$isMarriedWoman=new AndExpression($julie,$married);

var_dump("John is male? ",$isMale->interpret("abc"));
var_dump("Juile is a married women?",$isMarriedWoman->interpret("Married Julie"));