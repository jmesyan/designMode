package main

import "fmt"

type BusinessService interface {
	doProcessing()
}

type EJBService struct {

}

func(e *EJBService)doProcessing(){
	fmt.Println("Processing task by invoking EJB Service");
}

type JMSService struct {

}

func(e *JMSService)doProcessing(){
	fmt.Println("Processing task by invoking JMS Service");
}


type BusinessLookup struct {

}
func(b *BusinessLookup)getBusinessService(name string) BusinessService{
	if(name=="EJB"){
		return new(EJBService)
	}else{
		return new(JMSService)
	}
}


type BusinessDelegate struct {
	serviceType string
	businessLookup *BusinessLookup
}

func(b *BusinessDelegate)setSeviceType(serviceType string){
	b.serviceType=serviceType
}

func(b *BusinessDelegate)doTask(){
	service:=b.businessLookup.getBusinessService(b.serviceType)
	service.doProcessing()
}

type Client struct {
	businessDelegate *BusinessDelegate
}

func(c *Client)doTask(){
	c.businessDelegate.doTask()
}

func main(){
	businessDelegate:=&BusinessDelegate{"EJB",new(BusinessLookup)}
	client:=&Client{businessDelegate}
	client.doTask()

	businessDelegate.setSeviceType("JMS");
	client.doTask();

}

