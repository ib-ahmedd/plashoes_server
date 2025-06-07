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

type User struct {
	ID            int64  `json:"id"`
	User_name     string `json:"user_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Mobile_no     int    `json:"phone"`
	Date_of_birth string `json:"DOB"`
	Gender        string `json:"gender"`
	Country       string `json:"country"`
	Postal_code   int    `json:"postalcode"`
	Address       string `json:"address"`
	Country_code  string `json:"code"`
}

type OTP struct {
	Email string
}

type ClientOTP struct {
	OTP int
}