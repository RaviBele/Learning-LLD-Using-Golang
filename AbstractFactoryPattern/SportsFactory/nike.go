package main

type Nike struct{}

func NewNike() *Nike {
	return &Nike{}
}

func (n *Nike) makeShoes() IShoesFactory {
	return &NikeShoes{
		Shoes: Shoes{
			logo: "Nike",
			size: 20,
		},
	}
}

func (n *Nike) makeShirts() IShirtsFactory {
	return &NikeShirts{
		Shirts: Shirts{
			logo: "Nike",
			size: 50,
		},
	}
}
