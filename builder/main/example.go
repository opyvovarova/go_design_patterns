package main

import (
	"fmt"
	"go_design_patterns/builder/email"
	"go_design_patterns/builder"
)

func main() {
	e := email.EmailBuilder().To("john@doe.com").From("jane@doe.com").Subject("hola").Body("Hola carino").Build()
	fmt.Printf("%v", e)

	director := new(builder.Director)

	concreteBuilderA := new(builder.ConcreteBuilderA)
	director.Construct(concreteBuilderA)
	fmt.Printf("%v", concreteBuilderA.Build())

	concreteBuilderB := new(builder.ConcreteBuilderB)
	director.Construct(concreteBuilderB)
	fmt.Printf("%v", concreteBuilderB.Build())
}
