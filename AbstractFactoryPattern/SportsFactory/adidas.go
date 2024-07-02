package main

type Adidas struct{}

func NewAdidas() *Adidas {
	return &Adidas{}
}

func (n *Adidas) makeShoes() IShoesFactory {
	return &AdidasShoes{
		Shoes: Shoes{
			logo: "Adidas",
			size: 18,
		},
	}
}

func (n *Adidas) makeShirts() IShirtsFactory {
	return &AdidasShirts{
		Shirts: Shirts{
			logo: "Adidas",
			size: 40,
		},
	}
}
