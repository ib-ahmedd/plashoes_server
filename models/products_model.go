package models

import "plashoes-server/db"

type Product struct {
	ID            int64
	Product_name  string
	Price         float64
	Image         string
	Free_shipping bool
	Categories    string
	Sold          int
	Sale          bool
	Rating        float64
	Date_arrived  string
	Gender        string
	Color         string
}

func (product Product) Save() error{
	query := "INSERT INTO products (product_name, price, image, free_shipping, categories, sold, sale, rating, date_arrived, gender, color) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt,err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec(product.Product_name, product.Price, product.Image, product.Free_shipping, product.Categories, product.Sold, product.Sale, product.Rating, product.Date_arrived, product.Gender, product.Color)
	return err
}

func GetAllProducts () ([]map[string]any, error){
	query := "SELECT * FROM products"

	rows,err := db.DB.Query(query)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var productsSlice []map[string]any

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories, &product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
		if err != nil {
			return nil,err
		}

		mappedProduct := createProductMap(product)

		productsSlice = append(productsSlice, mappedProduct)
	}

	return productsSlice,nil
}

func createProductMap (product Product) map[string]any {	
	return map[string]any {
		"product_name": product.Product_name,
		"price": product.Price,
		"image": product.Image,
		"free_shipping": product.Free_shipping,
		"categories": product.Categories,
		"sold": product.Sold,
		"sale": product.Sale,
		"rating": product.Rating,
		"date_arrived": product.Date_arrived,
		"gender": product.Gender,
		"color": product.Color,
	}
}