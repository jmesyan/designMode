package main

import (
	"fmt"
	"container/list"
)

type Stock struct {
	name     string
	quantity int
}

func (s *Stock)buy() {
	fmt.Println("stock [name", s.name, ",quantity:", s.quantity, "] bought")
}

func (s *Stock)sell() {
	fmt.Println("stock [name", s.name, ",quantity:", s.quantity, "] sell")
}

type Order interface {
	execute()
}

type BuyOrder struct {
	stock *Stock
}

func (b *BuyOrder)execute() {
	b.stock.buy()
}

type SellOrder struct {
	stock *Stock
}

func (b *SellOrder)execute() {
	b.stock.sell()
}

type Broker struct {
	orderList *list.List
}

func (b *Broker)takeOrder(order Order) {
	b.orderList.PushBack(order)
}

func (b *Broker)placeOrder() {
	for o := b.orderList.Front(); o != nil; o = o.Next() {
		order := o.Value.(Order)
		order.execute()

	}
	//清空
	//for o:=b.orderList.Front();o!=nil;o=o.Next(){
	//	b.orderList.Remove(o)
	//}-无法完全清空
	b.orderList = list.New()
}

func main() {
	stock := &Stock{"abc", 10}
	buyOrder := &BuyOrder{stock}
	sellOrder := &SellOrder{stock}
	broker := &Broker{list.New()}
	broker.takeOrder(buyOrder)
	broker.takeOrder(sellOrder)
	broker.takeOrder(buyOrder)
	broker.placeOrder()
}



