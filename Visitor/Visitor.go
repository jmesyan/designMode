package main

import (
	"fmt"
)

type ComputerPartVister struct {

}

func (c *ComputerPartVister)visit(part ComputerPart) {
	switch part.(type){
	case *Mouse:
		fmt.Println("Displaying Mouse.")
	case *Monitor:
		fmt.Println("Displaying Monitor.")
	case *Keyboard:
		fmt.Println("Displaying Keyboard.")
	case *Computer:
		fmt.Println("Displaying Computer.")
	}
}

type ComputerPart interface {
	accept(visitor *ComputerPartVister)
}

type Mouse struct {

}

func (m *Mouse)accept(visitor *ComputerPartVister) {
	visitor.visit(m)
}

type Keyboard struct {

}

func (m *Keyboard)accept(visitor *ComputerPartVister) {
	visitor.visit(m)
}

type Monitor struct {

}

func (m *Monitor)accept(visitor *ComputerPartVister) {
	visitor.visit(m)
}

type Computer struct {
	parts []ComputerPart
}

func (c *Computer)accept(visitor *ComputerPartVister) {
	for _, part := range c.parts {
		part.accept(visitor)
	}

	visitor.visit(c)
}

func main() {
	parts := []ComputerPart{new(Monitor), new(Mouse), new(Keyboard)}
	computer := &Computer{parts}
	computer.accept(new(ComputerPartVister))
}


