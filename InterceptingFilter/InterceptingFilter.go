package main

import "fmt"

type Filter interface {
	execute(request string)
}

type AuthFilter struct {

}

func(a *AuthFilter)execute(request string){
	fmt.Println("Authing Reqeust:",request)
}

type DebugFilter struct {

}

func(a *DebugFilter)execute(request string){
	fmt.Println("log Reqeust:",request)
}

type Target struct {

}

func(a *Target)execute(request string){
	fmt.Println("executing Reqeust:",request)
}

type ChainFilter struct {
	filters []Filter
	target *Target
}

func(c *ChainFilter)setTarget(target *Target){
	c.target=target
}

func(c *ChainFilter)addFilter(filter Filter){
	c.filters=append(c.filters,filter)
}

func(c *ChainFilter)execute(request string){
	for _,filter:=range c.filters{
		filter.execute(request)
	}
	c.target.execute(request)
}


type FilterManager struct {
	filterChain *ChainFilter
}

func(f *FilterManager)init(target *Target){
	f.filterChain=new(ChainFilter)
	f.filterChain.setTarget(target)
}

func(f *FilterManager)setFilter(filter Filter){
	f.filterChain.addFilter(filter)
}

func(f *FilterManager)filterRequest(request string){
	f.filterChain.execute(request)
}


type Client struct {
	filterManager *FilterManager
}

func(c *Client)setFilterManager(filterManager *FilterManager){
	c.filterManager=filterManager
}

func(c *Client)sendRequest(request string){
	c.filterManager.filterRequest(request)
}


func main(){
	filterManager:=new(FilterManager)
	filterManager.init(new(Target))
	filterManager.setFilter(new(AuthFilter))
	filterManager.setFilter(new(DebugFilter))

	client:=new(Client)
	client.setFilterManager(filterManager)
	client.sendRequest("HOME")
}