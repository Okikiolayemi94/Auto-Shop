package shops

import (
	"fmt"
)

type Car struct {
	ID                string  
	Brand             string  
	Model             string 
	Quantity          float64 
	Price             float64
	Color             string  
	YearOfManufacture int64  
}

type CarOut struct {
	Car
}

// NewCar ... creates a new car type during use
func NewCar(brand, model, color string, quantity, price float64, year int64) *Car {
	return &Car{
		ID:                GenerateID(),
		Brand:             brand,
		Model:             model,
		Quantity:          quantity,
		Price:             price,
		Color:             color,
		YearOfManufacture: year,
	}
}
func (c *Car) GetName() string {
	return c.Brand + " " + c.Model
}

func (c *Car) GetID() string {
	return c.ID
}

func (c *Car) GetQuantity() float64 {
	return c.Quantity
}

func (c *Car) GetPrice() float64 {
	return c.Price
}

func (c *Car) GetBrand() string {
	return c.Brand
}

func (c *Car) GetModel() string {
	return c.Model
}

func (c *Car) GetColor() string {
	return c.Color
}

func (c *Car) SetQuantityIn(q float64) {
	c.Quantity -= q
}

func (c Car) SetQuantityOut(q float64) {
	c.Quantity += q
}

func (c *Car) GetYearOfManufacture() int64 {
	return c.YearOfManufacture
}

func (c *Car) GetStruct() interface{} {
	v := &Car{
		ID:                c.GetID(),
		Quantity:          c.GetQuantity(),
		Brand:             c.GetBrand(),
		Model:             c.GetModel(),
		Price:             c.GetPrice(),
		Color:             c.GetColor(),
		YearOfManufacture: c.GetYearOfManufacture(),
	}
	return v
}

func (c *Car) DisplayProductStatus() (string, bool) {
	msg := ""
	if c.Quantity > 0 {
		msg = fmt.Sprintf("Product is In stock")
		return msg, true
	}
	msg = fmt.Sprintf("Product currently out of stock")
	return msg, false
}
