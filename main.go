package main

import (
	"fmt"
	// array "plashoes-server/Array"
	"plashoes-server/models"

	"plashoes-server/db"
)

func main() {
	db.InitDB()

	// for _,item := range array.ProductsArray {
	// 	err := item.Save()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}

	// 	fmt.Println("Save success")
	// }

	products,err := models.GetAllProducts()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(products)
}