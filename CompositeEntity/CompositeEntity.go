package main

import "fmt"

type DependentObject1 struct {
	data string
}

func(d *DependentObject1)setData(data string){
	d.data=data
}

func(d *DependentObject1)getData() string{
	return d.data
}

type DependentObject2 struct {
	data string
}

func(d *DependentObject2)setData(data string){
	d.data=data
}

func(d *DependentObject2)getData() string{
	return d.data
}


type CoarseGrainedObject struct {
	do1 *DependentObject1
	do2 *DependentObject2
}

func(c *CoarseGrainedObject)setData(data1,data2 string){
	c.do1.setData(data1)
	c.do2.setData(data2)
}

func(c *CoarseGrainedObject)getData()[]string{
	return []string{c.do1.getData(),c.do2.getData()}
}

type CompositeEntity struct {
	cgo *CoarseGrainedObject
}

func(c *CompositeEntity)setData(data1,data2 string){
	c.cgo.setData(data1,data2)
}

func(c *CompositeEntity)getData()[]string{
	return c.cgo.getData()
}

type Client struct {
    compositeEntity *CompositeEntity
}

func(c *Client)printData(){
	for _,composite:=range c.compositeEntity.getData(){
		fmt.Println("Data:",composite)
	}
}


func(c *Client)setData(data1,data2 string){
	c.compositeEntity.setData(data1,data2)
}


func main(){
	client:=&Client{&CompositeEntity{&CoarseGrainedObject{new(DependentObject1),new(DependentObject2)}}}
	client.setData("Test", "Data")
	client.printData()

	client.setData("secend Test", "Data1")
	client.printData()

}



