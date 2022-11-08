package autos

import (
	"autos/shops"
	"fmt"
)

func main() {
	car1 := shops.NewCar("Honda", "Civic", "Grey", 4, 50000, 2020)
	car2 := shops.NewCar("Toyota", "Corolla", "ash", 3, 60000, 2008)

	// instantiate a new store
	store := shops.NewStore()

	//  adding item to the store
	store.AddItem(car1)
	store.AddItem(car2)
	

	// listing all products available in the store
	store.ListAllProduct()
	fmt.Println()

	// sell an item
	store.SellItem(car1, 1)
	store.SellItem(car2, 2)

	// list all sold item
	store.ListSoldItem()
	fmt.Println()
	store.ListAllProduct()

}
