<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/26
 * Time: 6:36
 */
class HomeView
{
    public function show()
    {
        var_dump("displaying home page");
    }
}

class StudentView
{
    public function show()
    {
        var_dump("displaying student page");
    }
}

class Dispatcher
{
    private $studentView, $homeView;

    public function __construct()
    {
        $this->studentView = new StudentView();
        $this->homeView = new HomeView();
    }

    public function dispatch($request)
    {
        if ($request == "STUDENT") {
            $this->studentView->show();
        } else {
            $this->homeView->show();
        }
    }
}

class FrontController
{
    private $dispatcher;

    public function __construct()
    {
        $this->dispatcher = new Dispatcher();
    }

    private function isAuthenticUser()
    {
        var_dump("User is authenticated successfully.");
        return true;
    }

    private function trackRequest($request){
          var_dump("Page Request:".$request);
    }

    public function dispatchRequest($request){
        $this->trackRequest($request);
        if($this->isAuthenticUser()){
            $this->dispatcher->dispatch($request);
        }
    }
}

//demo
$frontController=new FrontController();
$frontController->dispatchRequest("HOME");
$frontController->dispatchRequest("STUDENT");