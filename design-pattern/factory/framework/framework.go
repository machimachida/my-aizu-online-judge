package framework

type Product interface {
	Use()
}

type Factory interface {
	Create(owner string) Product
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type FactoryClass struct {
	Factory
}

func (f FactoryClass) Create(owner string) Product {
	p := f.CreateProduct(owner)
	f.RegisterProduct(p)
	return p
}
