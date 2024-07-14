package application

type Coin int

const (
	ONE  Coin = 1
	TWO  Coin = 2
	FIVE Coin = 5
	TEN  Coin = 10
)

type Product struct {
	name  string
	count int
	code  int
	price int
}

func NewProduct(name string, code int, price int, count int) *Product {
	return &Product{
		count: count,
		code:  code,
		price: price,
		name:  name,
	}
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetCode() int {
	return p.code
}

func (p *Product) GetPrice() int {
	return p.price
}

func (p *Product) GetCount() int {
	return p.count
}

func (p *Product) UpdateCount() {
	p.count -= 1
}
