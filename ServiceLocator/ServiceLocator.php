<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/26
 * Time: 8:32
 */
interface Severice{
    public function getName();
    public function execute();
}

class Service1 implements Severice{
    public function getName()
    {
        // TODO: Implement getName() method.
        return "Service1";
    }

    public function execute()
    {
        // TODO: Implement execute() method.
        var_dump("execute severice1");
    }
}

class Service2 implements Severice{
    public function getName()
    {
        // TODO: Implement getName() method.
        return "Service2";
    }

    public function execute()
    {
        // TODO: Implement execute() method.
        var_dump("execute severice2");
    }
}

class InitialContext{
    public function lookup($jndiName){
        $jndiName=strtolower($jndiName);
        if($jndiName=="service1"){
            var_dump("Looking up and creating a new Service1 object");
            return new Service1();
        }elseif($jndiName=="service2"){
            var_dump("Looking up and creating a new Service2 object");
            return new Service2();
        }
    }
}

class Cache{
    private $services=array();
    public function getService($serviceName){
        foreach($this->services as $service){
            if($service->getName()==$serviceName){
                return $service;
            }
        }
        return null;
    }

    public function addService($newService){
        $exists=false;
        foreach($this->services as $service){
            if($service->getName()==$newService->getName()){
                $exists=true;
                break;
            }
        }
        if(!$exists){
            $this->services[]=$newService;
        }

    }
}

class ServiceLocator{
    public  static $cache;

    public static function getService($serviceName){
        $service=self::$cache->getService($serviceName);
        if($service!=null){
            return $service;
        }

        $context=new InitialContext();
        $service1=$context->lookup($serviceName);
        self::$cache->addService($service1);
        return $service1;
    }
}

//demo
ServiceLocator::$cache=new Cache();
$service = ServiceLocator::getService("Service1");
$service->execute();
$service = ServiceLocator::getService("Service2");
$service->execute();
$service = ServiceLocator::getService("Service1");
$service->execute();
$service = ServiceLocator::getService("Service2");
$service->execute();