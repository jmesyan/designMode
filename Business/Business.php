<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 21:21
 */

interface BusinessService{
    public function doProcessing();
}

class EJBService  implements BusinessService{
    public function doProcessing()
    {
        // TODO: Implement doProcessing() method.
        var_dump("Processing task by invoking EJB Service");
    }
}


class JMSService  implements BusinessService{
    public function doProcessing()
    {
        // TODO: Implement doProcessing() method.
        var_dump("Processing task by invoking JMS Service");
    }
}


class BusinessLookUp{
    public function getBusinessService($serviceType){
        if($serviceType=="EJB"){
            return new EJBService();
        }else{
            return new JMSService();
        }

    }
}

class BusinessDelegate {
    private $lookupService;
    private $businessService;
    private $serviceType;
    public function __construct()
    {
        $this->lookupService=new BusinessLookUp();
    }

    public function setSevericeType($serviceType){
        $this->serviceType=$serviceType;
    }

    public function doTask(){
        $this->businessService=$this->lookupService->getBusinessService($this->serviceType);
        $this->businessService->doProcessing();
    }


}

class Client {
    private $businessService;
    public function __construct($BusinessDelegate)
    {
        $this->businessService=$BusinessDelegate;
    }
    public function doTask(){
        $this->businessService->doTask();
    }
}

//dmeo

$businessDelegate=new BusinessDelegate();
$client=new Client($businessDelegate);
$businessDelegate->setSevericeType("EJB");
$client->doTask();

$businessDelegate->setSevericeType("JMS");
$client->doTask();