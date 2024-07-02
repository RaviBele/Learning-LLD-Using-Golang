package main

import "fmt"

type Item struct {
	observerMap map[string]Observer
	name        string
	inStock     bool
}

func NewItem(name string) *Item {
	return &Item{name: name}
}

func (item *Item) updateAvailability() {
	fmt.Printf("Item %s is now available.\n", item.name)
	item.inStock = true
	item.notifyAll()
}

func (item *Item) registerObserver(observer Observer) {
	if item.observerMap == nil {
		item.observerMap = make(map[string]Observer)
	}
	item.observerMap[observer.getID()] = observer
}

func (item *Item) unregisterObserver(observer Observer) {
	delete(item.observerMap, observer.getID())
}

func (item *Item) notifyAll() {
	for _, observer := range item.observerMap {
		observer.update(item.name)
	}
}
