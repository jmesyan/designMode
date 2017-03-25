<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 22:35
 */
class DependentObject1
{
    private $data;
    public function setData($data){
        $this->data=$data;
    }

    public function getData(){
        return $this->data;
    }

}

class DependentObject2
{
    private $data;
    public function setData($data){
        $this->data=$data;
    }

    public function getData(){
        return $this->data;
    }

}

class CoarseGrainedObject{
    private $do1,$do2;
    public function __construct()
    {
        $this->do1=new DependentObject1();
        $this->do2=new DependentObject2();
    }

    public function setData($data1,$data2){
        $this->do1->setData($data1);
        $this->do2->setData($data2);
    }

    public function getData(){
        return array($this->do1->getData(),$this->do2->getData());
    }
}


class CompositeEntity{
    private $cgo ;
    public function __construct()
    {
        $this->cgo=new CoarseGrainedObject();
    }

    public function setData($data1,$data2){
        $this->cgo->setData($data1,$data2);
    }

    public function getData(){
        return $this->cgo->getData();
    }
}


class Client{
    private $compositeEntity;
    public function __construct()
    {
        $this->compositeEntity=new CompositeEntity();
    }

    public function printData(){
        foreach($this->compositeEntity->getData() as $data){
            var_dump("Data:".$data);
        }
    }

    public function setData($data1,$data2){
        $this->compositeEntity->setData($data1,$data2);
    }
}

//demo

$client=new Client();
$client->setData("Test", "Data");
$client->printData();
$client->setData("Second Test", "Data1");
$client->printData();
