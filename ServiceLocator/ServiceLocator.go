package main

import (
	"fmt"
	"strings"
)

type Service interface {
	getName() string
	execute()
}

type Service1 struct {

}

func(s *Service1)getName() string{
	return "Service1"
}

func(s *Service1)execute(){
	fmt.Println("execute service1")
}

type Service2 struct {

}

func(s *Service2)getName() string{
	return "Service2"
}

func(s *Service2)execute(){
	fmt.Println("execute service2")
}

type InitialContext struct {

}
func(l *InitialContext)lookup(serviceName string)Service{
	serviceName=strings.ToLower(serviceName)
	if(serviceName=="service1"){
		fmt.Println("Looking up and creating a new Service1 object")
		return new(Service1)
	}else if(serviceName=="service2"){
		fmt.Println("Looking up and creating a new Service2 object")
		return new(Service2)
	}
	return nil
}

type Cache struct {
	services []Service
}

func(c *Cache)getService(serviceName string)Service{
        for _,service:=range c.services{
		if(service.getName()==serviceName){
			return service
		}
	}
	return nil
}

func(c *Cache)addService(newService Service){
	exists:=false;
	for _,service:=range c.services{
		if(service.getName()==newService.getName()){
			exists=true
			break
		}
	}
	if !exists{
		c.services=append(c.services,newService)
	}
}

type ServiceLocator struct {
	context *InitialContext
	cache *Cache
}

func(s *ServiceLocator)init(){
	s.context=new(InitialContext)
	s.cache=new(Cache)
}

func(s *ServiceLocator)getService(serviceName string)Service{
      sevice:=s.cache.getService(serviceName)
	if(sevice!=nil){
		return sevice;
	}
	sevice1:=s.context.lookup(serviceName)
	s.cache.addService(sevice1)
	return sevice1
}



func main(){
	serviceLocator:=new(ServiceLocator)
	serviceLocator.init()
	service := serviceLocator.getService("Service1");
	service.execute();
	service = serviceLocator.getService("Service2");
	service.execute();
	service = serviceLocator.getService("Service1");
	service.execute();
	service = serviceLocator.getService("Service2");
	service.execute();
	
}