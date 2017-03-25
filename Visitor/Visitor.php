<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 19:34
 */
interface ComputerPart
{
    public function accept($computerPartVisitor);
}

class Keyboard implements ComputerPart
{
    public function accept($computerPartVisitor)
    {
        // TODO: Implement accept() method.
        $computerPartVisitor->visit($this);
    }
}


class Monitor implements ComputerPart
{
    public function accept($computerPartVisitor)
    {
        // TODO: Implement accept() method.
        $computerPartVisitor->visit($this);
    }
}

class Mouse implements ComputerPart
{
    public function accept($computerPartVisitor)
    {
        // TODO: Implement accept() method.
        $computerPartVisitor->visit($this);
    }
}


class Computer implements ComputerPart
{
    private $parts;

    public function __construct()
    {
        $this->parts=array(new Mouse(),new Keyboard(),new Monitor());
    }

    public function accept($computerPartVisitor)
    {
        // TODO: Implement accept() method.
        foreach($this->parts as $part){
            $part->accept($computerPartVisitor);
        }

        $computerPartVisitor->visit($this);
    }
}

interface computerPartVisitor{
    public function visit($part);
}

class ComputerPartDisplayVisitor  implements computerPartVisitor{
    public function visit($part)
    {
        // TODO: Implement visit() method.
        if(!is_object($part)) var_dump("not found part");
        $partName=get_class($part);
        switch ($partName){
            case "Computer":
                var_dump("Displaying Computer.");
                break;
            case "Mouse":
                var_dump("Displaying Mouse.");
                break;
            case "Keyboard":
                var_dump("Displaying Keyboard.");
                break;
            case "Monitor":
                var_dump("Displaying Monitor.");
                break;
        }
    }
}

//demo
$computer=new Computer();
$computer->accept(new ComputerPartDisplayVisitor());

