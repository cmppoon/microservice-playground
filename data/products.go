package data

import "fmt"

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"sku"`
}

type Products []*Product

var ErrProductNotFound = fmt.Errorf("product not found")

func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (*Product,error) {
	idx := findProductByID(id)

	if idx == -1 {
		return nil, ErrProductNotFound
	}

	return productList[idx], nil
}

func findProductByID(id int) int {
	for idx, prod := range productList {
		if prod.Id == id {
			return idx
		}
	}

	return -1
}

var productList = Products{
	&Product{
		Id:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		Sku:         "abc323",
	},
	&Product{
		Id:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		Sku:         "fjd34",
	},
}
