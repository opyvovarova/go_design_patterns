package main

import (
  "go_design_patterns/abstract_factory"
  )

func main() {
  c := new(abstract_factory.Client)
  factoryA := new(abstract_factory.ConcreteFactoryA)
  c.Produce(factoryA)
  factoryB := new(abstract_factory.ConcreteFactoryB)
  c.Produce(factoryB)
}
