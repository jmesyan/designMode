<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/26
 * Time: 7:23
 */
interface Filter{
    public function execute($request);
}

class AuthenticationFilter implements Filter{
    public function execute($request)
    {
       // TODO: Implement execute() method
        var_dump("Authenticating request:".$request);
    }
}

class DebugFilter implements Filter{
    public function execute($request)
    {
        // TODO: Implement execute() method
        var_dump("request log:".$request);
    }
}


class Target{
    public function execute($request)
    {
        // TODO: Implement execute() method
        var_dump("Executing request:".$request);
    }
}


class FilterChain {
    private $filters=array(),$target;
    public function addFilter($filter){
        $this->filters[]=$filter;
    }

    public function execute($request){
        foreach($this->filters as $filter){
            $filter->execute($request);
        }
        $this->target->execute($request);
    }

    public function setTarget($target){
        $this->target=$target;
    }
}

class FilterManager{
    private $filterChain;
    public function __construct($target)
    {
        $this->filterChain=new FilterChain();
        $this->filterChain->setTarget($target);
    }

    public function setFilter($filter){
        $this->filterChain->addFilter($filter);
    }

    public function filterRequest($request){
        $this->filterChain->execute($request);
    }
}

class Client{
    private $filterManager;
    public function setFilterManager($filterManager){
        $this->filterManager=$filterManager;
    }

    public function sendRequest($request){
        $this->filterManager->filterRequest($request);
    }
}

//demo
$filterManager=new FilterManager(new Target());
$filterManager->setFilter(new AuthenticationFilter());
$filterManager->setFilter(new DebugFilter());

$client=new Client();
$client->setFilterManager($filterManager);
$client->sendRequest("Home");
