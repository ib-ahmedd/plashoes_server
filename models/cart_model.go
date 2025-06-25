package models

import (
	"fmt"
	"plashoes-server/db"
)

func (item CartItem) Save() error {
	query := "INSERT INTO cart (user_id, product_id, quantity) VALUES (?, ?, ?)"

	_,err := db.DB.Exec(query, item.UserID, item.ProductID, item.Quantity)

	return err
}

func (item CartItem) Update() error {

	fmt.Println(item.UserID, item.ProductID, item.Quantity)
	query := "UPDATE cart SET quantity = ? WHERE user_id = ? AND product_id = ?"

	_,err := db.DB.Exec(query, item.Quantity, item.UserID, item.ProductID)

	return err
}

func (item CartItem) UpdateQuantity(itemID int64) error {
	query := "UPDATE cart SET quantity = ? WHERE id = ?"

	_,err := db.DB.Exec(query, item.Quantity, itemID)

	return err
}

func ItemInCart(userID int64, productID int64) (bool, error) {
	query := "SELECT COUNT(1) FROM cart WHERE user_id = ? AND product_id = ?"

	var count int
	err := db.DB.QueryRow(query, userID, productID).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
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

	return cartItemsArray,nil
	
}

func DeleteCartItem(itemID int64) error {
	query := "DELETE FROM cart WHERE id = ?"

	_,err := db.DB.Exec(query, itemID)

	return err
}

func EmptyCart(userID int64) error{
	query := "DELETE FROM cart WHERE user_id = ?"

	_,err := db.DB.Exec(query, userID)

	return err
}