package main

import (
	"fmt"
)

func main() {
	car1 := NewCar("Honda", "Civic", "Grey", 4, 50000, 2020)
	car2 := NewCar("Toyota", "Corolla", "ash", 3, 60000, 2008)
	
	// instantiate a new store
	store := NewStore()

	//  adding item
	store.AddItem(car1)
	store.AddItem(car2)
	
	// listing all products in the store
	store.ListAllProduct()
	fmt.Println()

	// sell item
	store.SellItem(car1, 1)
	store.SellItem(car2, 5)

	// list all sold item
	store.ListSoldItem()
	fmt.Println()
	store.ListAllProduct()

}
