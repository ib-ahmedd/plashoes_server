package models

import (
	"plashoes-server/db"
)

func GetCartItems(userId int64) ([]CartItem,error){
	query := "SELECT cart.id, cart.product_id, product_name, image, quantity, price FROM products JOIN cart ON products.id = cart.product_id WHERE cart.user_id = ?"

	rows, err := db.DB.Query(query, userId)

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var cartItemsArray []CartItem

	for rows.Next() {
		var cartItem CartItem

		err := rows.Scan(&cartItem.ID, &cartItem.Image, &cartItem.Price, &cartItem.ProductID, &cartItem.Quantity, &cartItem.Product_name)

		if err != nil {
			return nil,err
		}

		cartItemsArray = append(cartItemsArray, cartItem)
	}

	return cartItemsArray,nil
}