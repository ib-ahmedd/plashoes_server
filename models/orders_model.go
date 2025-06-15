package models

import (
	"fmt"
	"plashoes-server/db"
)

func (orderItem OrderItem) Save() error{
	query := "INSERT INTO orders (user_id, product_id, date_ordered, quantity, total_price, order_status, reviewed, date_delivered) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.DB.Exec(query, orderItem.UserID, orderItem.ProductID, orderItem.DateOrdered, orderItem.Quantity, orderItem.TotalPrice, orderItem.OrderStatus, orderItem.Reviewed, orderItem.DateDelivered)

	return err
} 

func GetOrders(userID int64) ([]OrderItem, error){
	query := "SELECT orders.id, orders.product_id, image, product_name, total_price, quantity, date_ordered, order_status FROM products JOIN orders ON products.id = orders.product_id WHERE orders.user_id = ?"

	rows,err := db.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ordersSlice []OrderItem

	for rows.Next() {
		var order OrderItem

		err := rows.Scan(&order.ID, &order.ProductID, &order.Image, &order.ProductName, &order.TotalPrice, &order.Quantity, &order.DateOrdered, &order.OrderStatus)

		if err != nil {
			return nil, err
		}

		ordersSlice = append(ordersSlice, order)
	}

	return ordersSlice,err
}

func GetPendingReviews(userID int64)([]OrderItem, error){
	query := "SELECT orders.id, orders.product_id, image, product_name, date_delivered FROM products JOIN orders ON products.id = orders.product_id WHERE orders.user_id = ? AND orders.order_status = 'Delivered'"

	rows,err := db.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ordersSlice []OrderItem

	for rows.Next() {
		var order OrderItem

		err := rows.Scan(&order.ID, &order.ProductID, &order.Image, &order.ProductName, &order.DateDelivered)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		ordersSlice = append(ordersSlice, order)
	}

	return ordersSlice,err
}

func GetOrderDetails(orderID int64)(OrderItem, error){
	query := "SELECT orders.id, orders.product_id, image, price, product_name, total_price, quantity, date_ordered, order_status FROM products JOIN orders ON products.id = orders.product_id WHERE orders.id = ?"

	var orderdetails OrderItem

	err := db.DB.QueryRow(query, orderID).Scan(&orderdetails.ID, &orderdetails.ProductID, &orderdetails.Image, &orderdetails.Price, &orderdetails.ProductName, &orderdetails.TotalPrice, &orderdetails.Quantity, &orderdetails.DateOrdered, &orderdetails.OrderStatus)

	if err != nil {
		fmt.Println(err)
		return OrderItem{}, nil
	}

	return orderdetails,err
}