<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 10:46
 */

interface Iterater{
    public function hasNext();
    public function next();
}

interface Container{
    public function getIterator();
}

class NameIterator implements Iterater{
    private $index=0;
    private $names;
    public function __construct(&$names)
    {
     $this->names=$names;
    }

    public function hasNext()
    {
        // TODO: Implement hasNext() method.
        if($this->index<count($this->names)){
            return true;
        }
        return false;
    }

    public function next()
    {
        // TODO: Implement next() method.
        if($this->hasNext()){
            return $this->names[$this->index++];
        }
        return null;
    }
}

class NameRepository implements Container{
    private $names=array("Robert","John","Julie","Lora");
    public function getIterator()
    {
        // TODO: Implement getIterator() method.
           return new NameIterator($this->names);
    }

}


//demo
$namesRepository=new NameRepository();
for($iter=$namesRepository->getIterator();$iter->hasNext();){
    $name=$iter->next();
    var_dump("Name :".$name);
}

