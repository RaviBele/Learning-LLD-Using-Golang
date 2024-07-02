package main

type ISportFactory interface {
	makeShoes() IShoesFactory
	makeShirts() IShirtsFactory
}

func GetSportsFactory(name string) ISportFactory {
	if name == "nike" {
		return NewNike()
	}

	if name == "adidas" {
		return NewAdidas()
	}

	return nil
}
