package abstract_factory

import "fmt"

// ======= Product related code:

type AbstractProduct interface {
  printAttributes()
}

type ProductA struct {
  name string
}

func NewProductA(name string) *ProductA {
  p := new(ProductA)
  p.name = name
  return p
}

func (p *ProductA) printAttributes() {
  fmt.Printf("Name : %v\n", p.name)
}

type ProductB struct {
  name string
  size int
  volume float64
}

func NewProductB(name string, size int, volume float64) *ProductB {
  p := new(ProductB)
  p.name = name
  p.size = size
  p.volume = volume
  return p
}

func (p *ProductB) printAttributes() {
  fmt.Printf("Name %v, size %d, volume %f\n", p.name, p.size, p.volume)
}

// ======= Factory related code

type AbstractFactory interface {
  createProductA() AbstractProduct
  createProductB() AbstractProduct
}

type ConcreteFactoryA struct {
}

func (c *ConcreteFactoryA) createProductA() AbstractProduct {
  p := NewProductA("Factory A :: product-A")
  return p
}

func (c *ConcreteFactoryA) createProductB() AbstractProduct {
  p := NewProductB("Factory A :: product-B", 10, 21.3)
  return p
}

type ConcreteFactoryB struct {
}

func (c *ConcreteFactoryB) createProductA() AbstractProduct {
  p := NewProductA("Factory B :: product-A")
  return p
}

func (c *ConcreteFactoryB) createProductB() AbstractProduct {
  p := NewProductB("Factory B :: product-B", 12, 65.4)
  return p
}

type Client struct {
}

func (c *Client) Produce(factory AbstractFactory) {
  pA := factory.createProductA();
  pA.printAttributes()
  //fmt.Printf("%v\n", pA)
  pB := factory.createProductB();
  pB.printAttributes()
  //fmt.Printf("%v\n", pB)
}
