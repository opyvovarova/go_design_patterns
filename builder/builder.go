package builder

type Builder interface {
	buildPart1()
	buildPart2()
	buildPart3()
}

type Product struct {
	s string
	i int
	f float64
}

type ConcreteBuilderA struct {
	product Product
}

type ConcreteBuilderB struct {
	product Product
}

func (b *ConcreteBuilderA) buildPart1() {
	b.product.s = "A"
}

func (b *ConcreteBuilderA) buildPart2() {
	b.product.i = 1
}

func (b *ConcreteBuilderA) buildPart3() {
	b.product.f = 6.4
}

func (b *ConcreteBuilderA) Build() Product {
	return Product{s : b.product.s, i : b.product.i,	f : b.product.f}
}

func (b *ConcreteBuilderB) buildPart1() {
	b.product.s = "B"
}

func (b *ConcreteBuilderB) buildPart2() {
	b.product.i = 2
}

func (b *ConcreteBuilderB) buildPart3() {
	b.product.f = 6.5
}

func (b *ConcreteBuilderB) Build() Product {
	p := Product{
		s : b.product.s,
		i : b.product.i,
		f : b.product.f}
	return p
}

type Director struct {
}

func (d Director) Construct(builder Builder) {
	builder.buildPart1()
	builder.buildPart2()
	builder.buildPart3()
}
