package main

import (
	app ""
	"fmt"
)

type Product struct {
	name    string
	price   int
	company string
}
type SellableProduct interface {
	buy() string
}

//constructor *act as

func newProduct(name string, price int, company string) *Product {

	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return &p //this is returning a copy of the product
}

func main() {



	// var p1 Product
	// p1.name = "Laptop"
	// p1.price = 1000
	// p1.company = "Dell"

	// fmt.Println("Product Name:", p1.name)
	// fmt.Println("Product Price:", p1.price)
	// fmt.Println("Product Company:", p1.company)

		c := app.NewCOmplex(1.0, 2.0)
	fmt.Println(c)

}
