package product

type ProductInterface interface {
	GetStruct() interface{}
	DisplayProductStatus() (string, bool)
	GetName() string
	GetID() string
	GetQuantity() float64
	GetPrice() float64
	GetBrand() string
	GetModel() string
	GetYearOfManufacture() int64
	SetQuantityIn(quantity float64)
	SetQuantityOut(quantity float64)
}

type Product struct {
	Product ProductInterface
}
