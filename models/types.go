package models

type Product struct {
	ID            int64   `json:"id"`
	Product_name  string  `json:"product_name"`
	Price         float64 `json:"price"`
	Image         string  `json:"image"`
	Free_shipping bool    `json:"free_shipping"`
	Categories    string  `json:"categories"`
	Sold          int     `json:"sold"`
	Sale          bool    `json:"sale"`
	Rating        float64 `json:"rating"`
	Date_arrived  string  `json:"date_arrived"`
	Gender        string  `json:"gender"`
	Color         string  `json:"color"`
}

type CartItem struct {
	ID           int64   `json:"id"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	UserID       int64   `json:"user_id"`
	ProductID    int64   `json:"product_id"`
	Quantity     int     `json:"quantity"`
	Product_name string  `json:"product_name"`
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

type OrderItem struct {
	ID            int64   `json:"id"`
	UserID        int64   `json:"user_id"`
	ProductID     int64   `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Image         string  `json:"image"`
	Price         string  `json:"price"`
	Quantity      int     `json:"quantity"`
	TotalPrice    float64 `json:"totalprice"`
	OrderStatus   string  `json:"order_status"`
	Reviewed      bool    `json:"reviewed"`
	DateOrdered   string  `json:"date_ordered"`
	DateDelivered string  `json:"date_delivered"`
}

type Review struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	ProductID    int64  `json:"product_id"`
	OrderID      int64  `json:"order_id"`
	ReviewerName string `json:"reviewer_name"`
	Stars        int    `json:"stars"`
	ReviewTitle  string `json:"review_title"`
	ReviewDetail string `json:"review_detail"`
	ReviewDate   string `json:"review_date"`
}

type FilterSort struct {
	Page       string
	PriceRange float64
	Sort       string
	Category   string
	Offset     int
}

type OTP struct {
	Email string
}

type ClientOTP struct {
	OTP int
}

type LoginDetails struct {
	Email    string
	Password string
}

type CartRequest struct {
	ID        int64 `json:"id"`
	UserID    int   `json:"user_id"`
	ProductID int   `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type OrderRequest struct {
	UserID     int64       `json:"user_id"`
	OrderItems []OrderItem `json:"ordered_items"`
}