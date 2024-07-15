package models

type User struct {
	name    string
	card    *Card
	account *Account
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) SetAccount(account *Account) {
	u.account = account
}

func (u *User) SetCard(card *Card) {
	u.card = card
}

func (u *User) GetCard() *Card {
	return u.card
}

func (u *User) GetAccount() *Account {
	return u.account
}

func (u *User) GetName() string {
	return u.name
}
