package models

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

type FilterSort struct {
	Page       string
	PriceRange float64
	Sort       string
	Category   string
	Offset     int
}