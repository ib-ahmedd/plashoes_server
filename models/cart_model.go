package models

import (
	"fmt"
	"plashoes-server/db"
)

func (item CartRequest) Save() error {
	query := "INSERT INTO cart (user_id, product_id, quantity) VALUES (?, ?, ?)"

	_,err := db.DB.Exec(query, item.UserID, item.ProductID, item.Quantity)

	return err
}

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

		err := rows.Scan(&cartItem.ID, &cartItem.ProductID, &cartItem.Product_name, &cartItem.Image, &cartItem.Quantity, &cartItem.Price) 

		if err != nil {
			return nil,err
		}

		cartItemsArray = append(cartItemsArray, cartItem)
	}

	fmt.Println(cartItemsArray)

	return cartItemsArray,nil
	
}

