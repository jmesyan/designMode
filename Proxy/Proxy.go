package main

import "fmt"

type Image interface {
	display()
}

type RealImage struct {
	fileName string
}

func(r *RealImage)loadFromDisk(fileName string){
	fmt.Println("load disk file:",fileName)
}

func(r *RealImage)init(fileName string){
	r.fileName=fileName
	r.loadFromDisk(fileName)
}

func(r *RealImage)display(){
	fmt.Println("displaying:",r.fileName)
}

type ProxyImage struct {
	realImage *RealImage
	fileName string
}

func(p *ProxyImage)display(){
	if(p.realImage==nil){
		p.realImage=new(RealImage)
		p.realImage.init(p.fileName)
	}
	p.realImage.display()
}

func main(){
	proxyImage:=&ProxyImage{fileName:"abcx.jpg"}
	proxyImage.display()
	proxyImage.display()
}