package main

import "fmt"

type AbstractCustomer interface {
	isNil() bool
	getName() string
}

type RealCustommer struct {
	name string
}

func(r *RealCustommer)isNil() bool{
	return false;
}

func(r *RealCustommer)getName() string{
	return r.name;
}

type NullCustommer struct {
	name string
}

func(r *NullCustommer)isNil() bool{
	return true;
}

func(r *NullCustommer)getName() string{
	return "not available in customer database";
}

type CustomerFactory struct {
	names []string
}

func(c *CustomerFactory)getCustomer(name string) AbstractCustomer{
	c.names=[]string{"Rob", "Job", "Juile"}
	for _,n:=range c.names{
		if(n==name){
			return &RealCustommer{n}
		}
	}
	return new(NullCustommer)
}

func main(){
	custFac:=new(CustomerFactory);
	customer1 := custFac.getCustomer("Rob");
	customer2 := custFac.getCustomer("Bob");
	customer3 := custFac.getCustomer("Juile");
	customer4 := custFac.getCustomer("Laura");
	fmt.Println("Customers:");
	fmt.Println(customer1.getName())
	fmt.Println(customer2.getName())
	fmt.Println(customer3.getName())
	fmt.Println(customer4.getName())
}
