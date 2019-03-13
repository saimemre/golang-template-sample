package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	model "./models"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := model.Page{Title: "Selamlar Go Bey"}
	products := loadProducts()
	categories := loadCategories()
	relationMappings := loadIRelationMappings()

	var newProducts []model.Product

	for _, product := range products {

		for _, relationMapping := range relationMappings {
			if product.ID == relationMapping.ProductID {
				for _, category := range categories {
					if relationMapping.CategoryID == category.ID {
						product.Categories = append(product.Categories, category)
					}
				}
			}

		}

		newProducts = append(newProducts, product)
	}

	fmt.Printf("%+v\n", newProducts)
	view := model.View{Page: page, Products: newProducts}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, view)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func loadProducts() []model.Product {
	bytes, _ := ioutil.ReadFile("json/products.json")
	var products []model.Product
	json.Unmarshal(bytes, &products)
	return products
}

func loadCategories() []model.Category {
	bytes, _ := ioutil.ReadFile("json/categories.json")
	var categories []model.Category
	json.Unmarshal(bytes, &categories)
	return categories
}

func loadIRelationMappings() []model.RelationMapping {
	bytes, _ := ioutil.ReadFile("json/relation.json")
	var relationMappings []model.RelationMapping
	json.Unmarshal(bytes, &relationMappings)
	return relationMappings
}
