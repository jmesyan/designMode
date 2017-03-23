package main

import (
	"container/list"
	"fmt"
)

type Employee struct {
	name string
	dept string
	salary int
	subordinates *list.List
}

func(e *Employee)add(n *Employee){
	e.subordinates.PushBack(n)
}

func(e *Employee)remove(n *Employee){
	for l:=e.subordinates.Front();l!=nil;l=l. Next(){
		if(l.Value.(*Employee)==n){
			e.subordinates.Remove(l)
			break;
		}
	}
}

func(e *Employee)getSubordinates() *list.List{
	return e.subordinates
}

func(e *Employee)printEmployee(){
	fmt.Println( "Employee:[ Name : " ,e.name, ", dept : ",e.dept, ", salary :" ,e.salary , " ]");
}

func main(){
	ceo:=&Employee{"John", "CEO", 30000,list.New()}
	headSales:=&Employee{"Robert", "Head Sales", 20000,list.New()}
	headMarking:=&Employee{"Michel", "Head Marketing", 20000,list.New()}
	clerk1:=&Employee{"Laura", "Marketing", 10000,list.New()}
	clerk2:=&Employee{"Bob", "Marketing", 10000,list.New()}
	salesExecutive1:=&Employee{"Richard", "Sales", 10000,list.New()}
	salesExecutive2:=&Employee{"Rob", "Sales", 10000,list.New()}

	ceo.add(headSales)
	ceo.add(headMarking)

	headSales.add(salesExecutive1)
	headSales.add(salesExecutive2)

	headMarking.add(clerk1)
	headMarking.add(clerk2)

	ceo.printEmployee()

	for e:=ceo.subordinates.Front();e!=nil;e=e.Next(){
		headEmployee:=e.Value.(*Employee)
		headEmployee.printEmployee()
		for e:=headEmployee.subordinates.Front();e!=nil;e=e.Next(){
			employee:=e.Value.(*Employee)
		        employee.printEmployee()
		}
	}
}