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

func GetAllProducts () ([]Product, error){
	query := "SELECT * FROM products"

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

var ProductsArray []Product = []Product{
	{
		Product_name: "Men's black running",
		Price: 79.90,
		Image: "https://res.cloudinary.com/djnzi39nh/image/upload/v1747847628/men_black_running_la3hl3.jpg",
		Free_shipping: true,
		Categories: "running",
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