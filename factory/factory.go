package factory

type Product interface {
	Name() string
}

type ProductA struct {
}

func (p *ProductA) Name() string {
	return "ProductA"
}

type ProductB struct {
}

func (p *ProductB) Name() string {
	return "ProductB"
}

type Creator interface {
	Create() Product
}

type CreatorA struct {
}

func (c *CreatorA) Create() Product {
	return new(ProductA)
}

type CreatorB struct {
}

func (c *CreatorB) Create() Product {
	return new(ProductB)
}
