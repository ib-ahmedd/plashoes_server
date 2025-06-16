package models

import (
	"database/sql"
	"fmt"
	"plashoes-server/db"
	"slices"
)

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

func UpdateRating(productId int64, newRating int) error {
	query := "UPDATE products SET rating = ? WHERE id = ?"
	
	_,err := db.DB.Exec(query, newRating, productId)

	return err
}


func GetProductsFromDB(query string) ([]Product, error){
	rows,err := db.DB.Query(query)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var productsSlice []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories, &product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
		if err != nil {
			return nil,err
		}

		productsSlice = append(productsSlice, product)
	}

	return productsSlice,nil
}

func GetSingleProduct(id int64) ([]any, error) {
	query := "SELECT * FROM products WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var product Product
	err := row.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories,&product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	relatedProducts, err := getRelatedProducts(product.Categories, product.ID)

	if err != nil {
		return nil,err
	}

	reviews,err := getProductReviews(id)

	if err != nil {
		return nil,err
	}

	return []any{&product, relatedProducts, reviews}, nil
}

func getRelatedProducts(categ string, ID int64 ) ([]Product, error){
	query := "SELECT * FROM products WHERE categories = ? AND id != ?"
	rows,err := db.DB.Query(query, categ, ID) 
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var productsSlice []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories, &product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
		if err != nil {
			return nil,err
		}

		productsSlice = append(productsSlice, product)
	}

	return productsSlice,nil
}

func getProductReviews(productID int64) ([]Review,error){
	commentsQuery := "SELECT user_id, rating, review_title, review_detail, review_date, reviewer_name FROM reviews WHERE product_id = ?"

	
	rows,err := db.DB.Query(commentsQuery, productID)

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var reviews []Review

	for rows.Next(){
		var review Review

		err := rows.Scan(&review.UserID, &review.Stars, &review.ReviewTitle, &review.ReviewDetail, &review.ReviewDate, &review.ReviewerName)

		if err != nil {
			return nil,err
		}

		reviews = append(reviews, review)
	}

	return reviews,nil
}

func GetProductPage(page string) (map[string]any, error){
	query := "SELECT * FROM products WHERE gender = ?"

	if page == "Sale" {
		query = "SELECT * FROM products WHERE sale = true LIMIT 12"
	}
	rows,err := db.DB.Query(query, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productsSlice []Product
	var categories []string
	var prices []float64

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories, &product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
		if err != nil {
			return nil,err
		}

		if !slices.Contains(categories, product.Categories){
			categories = append(categories, product.Categories)
		}

		prices = append(prices, product.Price)
		productsSlice = append(productsSlice, product)
	}

	count := len(productsSlice)
	priceRange := map[string]float64 {"minPrice" : slices.Min(prices), "maxPrice": slices.Max(prices)}
	requestResponse := map[string]any{"data": productsSlice, "categoriesData": categories, "count": count, "range": priceRange}
	
	return requestResponse, err
}

func GetAllProducts() (map[string]any, error){
	productsQuery := "SELECT * FROM products LIMIT 12"
	countQuery := "SELECT COUNT(*) FROM products"
	categoriesQuery := "SELECT DISTINCT categories FROM products"
	priceRangeQuery := "SELECT MIN(price) AS min_col, MAX(price) AS max_col FROM products"
	

//----------------products query----------------
//----------------products query----------------
	var productsSlice []Product
	rows,err := db.DB.Query(productsQuery)

	if err != nil {
		return nil,err
	}
	defer rows.Close()


	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories, &product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
		if err != nil {
			return nil,err
		}

		productsSlice = append(productsSlice, product)
	}

//----------------count query----------------
//----------------count query----------------

	var count int

	err = db.DB.QueryRow(countQuery).Scan(&count)

	if err != nil {
		
		return nil,err
	}
//----------------categories query----------------
//----------------categories query----------------

	var categories []string

	catRows, err := db.DB.Query(categoriesQuery)

	if err != nil {
		return nil,err
	}

	defer catRows.Close()

	for catRows.Next() {
		var category string
		err := catRows.Scan(&category)
		if err != nil {
			return nil,err
		}

		categories = append(categories, category)
	}

//----------------price range query----------------
//----------------price range query----------------

	var minPrice, maxPrice float64

	err = db.DB.QueryRow(priceRangeQuery).Scan(&minPrice, &maxPrice)

	if err != nil {
		return nil,err
	}

	priceRange := map[string]float64{"minPrice": minPrice, "maxPrice": maxPrice}

	requestResponse := map[string]any{"data": productsSlice, "categoriesData": categories, "count": count, "range": priceRange}

	return requestResponse,nil
}

func FilterSortProducts(requestBody FilterSort)(map[string]any, error){
	sortOrder := getSortOrder(requestBody.Sort)
	var query string
	var rows *sql.Rows
	var countQuery string
	var err error
	var countErr error
	var count int
	
	if requestBody.Category != "Default" {
		if requestBody.Page == "Men" || requestBody.Page == "Women" {
			query = "SELECT * FROM products WHERE categories = ? AND price < ? AND gender = ? ORDER BY " + sortOrder + " LIMIT 12 OFFSET ?"
			rows,err = db.DB.Query(query, requestBody.Category, requestBody.PriceRange, requestBody.Page, requestBody.Offset)

			countQuery =  "SELECT COUNT(*) FROM products WHERE categories = ? AND gender = ? AND price < ?"
			countErr = db.DB.QueryRow(countQuery, requestBody.Category, requestBody.Page, requestBody.PriceRange).Scan(&count)
			
		} else if requestBody.Page == "Sale" {
			query = "SELECT * FROM products WHERE categories = ? AND sale = true AND price < ? ORDER BY " + sortOrder + " LIMIT 12 OFFSET ?"
			rows,err = db.DB.Query(query, requestBody.Category, requestBody.PriceRange, requestBody.Offset)

			countQuery = "SELECT COUNT(*) FROM products WHERE categories = ? AND sale = true AND price < ?"
			countErr = db.DB.QueryRow(countQuery, requestBody.Category, requestBody.PriceRange, requestBody.Offset).Scan(&count)

		} else {
			query = "SELECT * FROM products WHERE categories = ? AND price < ? ORDER BY " + sortOrder + " LIMIT 12 OFFSET ?"
			rows,err = db.DB.Query(query, requestBody.Category, requestBody.PriceRange, requestBody.Offset)

			countQuery = "SELECT COUNT(*) FROM products WHERE categories = ? AND price < ?"
			countErr = db.DB.QueryRow(countQuery, requestBody.Category, requestBody.PriceRange, requestBody.Offset).Scan(&count)
		}
	} else {
		if requestBody.Page == "Men" || requestBody.Page == "Women" {
			fmt.Println(requestBody.Page, requestBody.PriceRange, requestBody.Offset)
			query = "SELECT * FROM products WHERE gender = ? AND price < ? ORDER BY " + sortOrder + " LIMIT 12 OFFSET ?"
			rows,err = db.DB.Query(query, requestBody.Page, requestBody.PriceRange, requestBody.Offset)

			countQuery = "SELECT COUNT(*) FROM products WHERE gender = ? AND price < ?"
			countErr = db.DB.QueryRow(countQuery, requestBody.Page, requestBody.PriceRange, requestBody.Offset).Scan(&count)

		} else if requestBody.Page == "Sale" {
			query = "SELECT * FROM products WHERE sale = true AND price < ? ORDER BY " + sortOrder + " LIMIT 12 OFFSET ?"
			rows,err = db.DB.Query(query, requestBody.PriceRange, requestBody.Offset)

			countQuery = "SELECT COUNT(*) FROM products WHERE sale = true AND price < ?"
			countErr = db.DB.QueryRow(countQuery, requestBody.PriceRange, requestBody.Offset).Scan(&count)

		} else {
			query = "SELECT * FROM products WHERE price < ? ORDER BY " + sortOrder + " LIMIT 12 OFFSET ?"
			rows,err = db.DB.Query(query, requestBody.PriceRange, requestBody.Offset)

			countQuery = "SELECT COUNT(*) FROM products WHERE price < ?"
			countErr = db.DB.QueryRow(countQuery, requestBody.PriceRange, requestBody.Offset).Scan(&count)
		}
	}

	if countErr != nil { 
		return nil,err
	}

	if err != nil {
		return nil,err
	}

	var productsSlice []Product

	defer rows.Close()


	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Product_name, &product.Price, &product.Image, &product.Free_shipping, &product.Categories, &product.Sold, &product.Sale, &product.Rating, &product.Date_arrived, &product.Gender, &product.Color)
		if err != nil {
			return nil,err
		}

		productsSlice = append(productsSlice, product)
	}

	responseData := map[string]any{"data": productsSlice, "count": count}
	return responseData,nil
}

var ProductsArray []Product = []Product{
	{
		Product_name: "Men's black running",
		Price: 79.90,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/men_black_running_la3hl3.jpg",
		Free_shipping: true,
		Categories: "Running",
		Sold: 22,
		Sale: false,
		Rating: 4,
		Date_arrived: "2016-10-07",
		Gender: "Men",
		Color: "black",
	},
	{
		Product_name: "Men's Navy Running",
		Price: 79.90,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847629/men_navy_running_t9tqdy.jpg",
		Free_shipping: false,
		Categories: "Running",
		Sold: 102,
		Sale: false,
		Rating: 4,
		Date_arrived: "2020-05-12",
		Gender: "Men",
		Color: "navy",
	},
	{
		Product_name: "Men's Green Running",
		Price: 79.90,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/men_green_running_pk6aho.jpg",
		Free_shipping: true,
		Categories: "Running",
		Sold: 12,
		Sale: true,
		Rating: 2,
		Date_arrived: "2018-07-09",
		Gender: "Men",
		Color: "green",
	},
	{
		Product_name: "Men's Red Running",
		Price: 79.90,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847629/men_red_running_t5twuu.jpg",
		Free_shipping: false,
		Categories: "Running",
		Sold: 12,
		Sale: false,
		Rating: 3,
		Date_arrived: "2023-10-28",
		Gender: "Men",
		Color: "red",
	},
	{
		Product_name: "Men's Classic mint",
		Price: 104.20,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/men_classic_mint_tywrxa.jpg",
		Free_shipping: false,
		Categories: "Classic",
		Sold: 66,
		Sale: false,
		Rating: 0,
		Date_arrived: "2022-12-05",
		Gender: "Men",
		Color: "Blue",
	},
	{
		Product_name: "Men's Classic Blue",
		Price: 104.20,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/men_classic_blue_f9f0zn.jpg",
		Free_shipping: false,
		Categories: "Classic",
		Sold: 36,
		Sale: false,
		Rating: 4,
		Date_arrived: "2020-06-23",
		Gender: "Men",
		Color: "blue",
	},
	{
		Product_name: "Men's Moonstone Sneaker",
		Price: 93.50,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/men_moonstone_sneaker_dvh0bh.jpg",
		Free_shipping: true,
		Categories: "Sneaker",
		Sold: 45,
		Sale: true,
		Rating: 5,
		Date_arrived: "2019-02-20",
		Gender: "Men",
		Color: "grey",
	},
	{
		Product_name: "Men's Earth-tone Sneaker",
		Price: 93.50,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847627/men_earth-tone_sneaker_khxcnn.jpg",
		Free_shipping: false,
		Categories: "Sneaker",
		Sold: 23,
		Sale: false,
		Rating: 2,
		Date_arrived: "2020-01-15",
		Gender: "Men",
		Color: "brown",
	},
	{
		Product_name: "Women's Tosca City Run",
		Price: 89.40,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_tosca_city_run_i8aszi.jpg",
		Free_shipping: false,
		Categories: "Sneaker",
		Sold: 77,
		Sale: true,
		Rating: 3,
		Date_arrived: "2024-02-03",
		Gender: "Women",
		Color: "brown",
	},
	{
		Product_name: "Women's Candy City Run",
		Price: 89.40,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/women_candy_city_run_qykosn.jpg",
		Free_shipping: false,
		Categories: "Sneaker",
		Sold: 101,
		Sale: true,
		Rating: 5,
		Date_arrived: "2019-07-16",
		Gender: "Women",
		Color: "blue",
	},
	{
		Product_name: "Women's Choco City Run",
		Price: 89.40,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847629/women_choco_city_run_lmjfzw.jpg",
		Free_shipping: true,
		Categories: "Sneaker",
		Sold: 27,
		Sale: false,
		Rating: 4,
		Date_arrived: "2024-02-03",
		Gender: "Women",
		Color: "brown",
	},
	{
		Product_name: "Women's Pink Training",
		Price: 69.80,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_pink_training_yy6a06.jpg",
		Free_shipping: true,
		Categories: "Training",
		Sold: 89,
		Sale: false,
		Rating: 4,
		Date_arrived: "2022-09-10",
		Gender: "Women",
		Color: "pink",
	},
	{
		Product_name: "Women's Peach Training",
		Price: 69.80,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_peach_training_b8e6kv.jpg",
		Free_shipping: false,
		Categories: "Training",
		Sold: 90,
		Sale: true,
		Rating: 5,
		Date_arrived: "2019-03-22",
		Gender: "Women",
		Color: "peach",
	},
	{
		Product_name: "Women's Green Training",
		Price: 69.80,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_green_training_jzdsf2.jpg",
		Free_shipping: false,
		Categories: "Training",
		Sold: 47,
		Sale: false,
		Rating: 3,
		Date_arrived: "2016-06-07",
		Gender: "Women",
		Color: "green",
	},
	{
		Product_name: "Women's Blue Training",
		Price: 69.80,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_green_training_jzdsf2.jpg",
		Free_shipping: false,
		Categories: "Training",
		Sold: 37,
		Sale: false,
		Rating: 5,
		Date_arrived: "2017-10-30",
		Gender: "Women",
		Color: "blue",
	},
	{
		Product_name: "Women's Pink Suede",
		Price: 74.20,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847629/women_pink_suede_avkbig.jpg",
		Free_shipping: true,
		Categories: "Suede",
		Sold: 67,
		Sale: true,
		Rating: 3,
		Date_arrived: "2016-08-22",
		Gender: "Women",
		Color: "pink",
	},
	{
		Product_name: "Women's Cream Suede",
		Price: 74.20,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/women_cream_suede_tgpivh.jpg",
		Free_shipping: true,
		Categories: "Suede",
		Sold: 67,
		Sale: false,
		Rating: 5,
		Date_arrived: "2020-07-06",
		Gender: "Women",
		Color: "cream",
	},
	{
		Product_name: "Women's Tan Sneaker",
		Price: 109.70,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_tan_sneaker_eqkhge.jpg",
		Free_shipping: false,
		Categories: "Sneaker",
		Sold: 88,
		Sale: false,
		Rating: 3,
		Date_arrived: "2023-01-08",
		Gender: "Women",
		Color: "tan",
	},
	{
		Product_name: "Women's Orange Sneaker",
		Price: 109.70,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847630/women_orange_sneaker_du933p.jpg",
		Free_shipping: false,
		Categories: "Sneaker",
		Sold: 99,
		Sale: false,
		Rating: 4,
		Date_arrived: "2019-04-14",
		Gender: "Women",
		Color: "Orange",
	},
	{
		Product_name: "Women's Mint Sneaker",
		Price: 109.70,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847629/women_mint_sneaker_hrmjew.jpg",
		Free_shipping: true,
		Categories: "Sneaker",
		Sold: 78,
		Sale: true,
		Rating: 5,
		Date_arrived: "2017-10-07",
		Gender: "Women",
		Color: "blue",
	},
}