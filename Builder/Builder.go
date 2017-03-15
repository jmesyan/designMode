package main

import (
	"container/list"
	"fmt"
)


//定义抽象接口
type Packing interface {
	pack() string
}

type Item interface {
	name() string
	packing() Packing
	price() float64
}

//实现packing接口
type Wrapper struct {
}

func (Packing *Wrapper)pack() string {
	return "Wrapper"
}

type Bottle struct {
}

func (Packing *Bottle)pack() string {
	return "Bottle";
}


//实现Item接口

type Burger struct {

}

func (Item *Burger)packing() Packing {
	return new(Wrapper)
}

type vegBurger struct {
	Burger
}

func (Item *vegBurger)name() string {
	return "Verg Burger"
}

func (Item *vegBurger)price() float64 {
	return 21.3
}

type chickenBurger struct {
	Burger
}

func (Item *chickenBurger)name() string {
	return "chicken Burger"
}

func (Item *chickenBurger)price() float64 {
	return 22.6
}

type ColdDrink struct {

}

func (Item *ColdDrink)packing() Packing {
	return new(Bottle)
}

type Coke struct {
	ColdDrink
}

func (Item *Coke)name() string {
	return "Coke"
}

func (Item *Coke)price() float64 {
	return 15.6
}

type Pesi struct {
	ColdDrink
}

func (Item *Pesi)name() string {
	return "Pesi"
}

func (Item *Pesi)price() float64 {
	return 18.6
}

type Meal struct {
	items list.List
}

func (this *Meal)addItem(item Item) {
	this.items.PushBack(item)
}

func (this *Meal)getCost() float64 {
	cost := 0.0
	for e := this.items.Front(); e != nil; e = e.Next() {
		cost += e.Value.(Item).price()
	}
	return cost
}

func (this *Meal)showItems() {
	for e := this.items.Front(); e != nil; e = e.Next() {
		item := e.Value.(Item)
		fmt.Println("item: ", item.name(), ",Packing :", item.packing().pack(), ",Price :", item.price())
	}
}

type MealBuilder struct {

}

func (this *MealBuilder)prepareVegMeal() *Meal {
	meal := new(Meal)
	meal.addItem(new(vegBurger))
	meal.addItem(new(Coke))
	return meal
}

func (this *MealBuilder)prepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.addItem(new(chickenBurger))
	meal.addItem(new(Pesi))
	return meal
}

func main() {
	mealBuilder := new(MealBuilder)
	vegMeal := mealBuilder.prepareVegMeal()
	fmt.Println("Veg Meal")
	vegMeal.showItems()
	fmt.Println("Total cost:", vegMeal.getCost())

	nonVegMeal := mealBuilder.prepareNonVegMeal()
	fmt.Println("NonVeg Meal")
	nonVegMeal.showItems()
	fmt.Println("Total cost:", nonVegMeal.getCost())

}