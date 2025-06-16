package models

import (
	"plashoes-server/db"
	"time"
)

func (review Review) Save() error {
	query := "INSERT INTO reviews (user_id, product_id, order_id, reviewer_name, rating, review_title, review_detail, review_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	review.ReviewDate = time.Now().String()

	_, err := db.DB.Exec(query, review.UserID, review.ProductID, review.OrderID, review.ReviewerName, review.Stars, review.ReviewTitle, review.ReviewDetail, review.ReviewDate)

	return err
}

func (review Review)GetNewAverageProductRating() (int,error){
	sumQuery := "SELECT AVG(rating) FROM reviews WHERE product_id = ?"

	var average int

	err := db.DB.QueryRow(sumQuery, review.ProductID).Scan(&average)

	if err != nil {
		return 0,err
	}

	return average, err
}