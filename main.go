package main

import (
	// "fmt"

	"plashoes-server/db"
	"plashoes-server/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "plashoes-server/models"
)

func main() {
	db.InitDB()

	// rows,_ := db.DB.Query("SELECT * FROM users")
	// defer rows.Close()

	// var usersSlice []models.User

	// for rows.Next() {
	// 	var user models.User
	// 	rows.Scan(&user.ID, &user.User_name, &user.Email, &user.Password, &user.Mobile_no, &user.Date_of_birth, &user.Gender, &user.Country, &user.Postal_code, &user.Address, &user.Country)

	// 	usersSlice = append(usersSlice, user)
	// }

	// fmt.Println(usersSlice)

	// db.DB.Exec("DELETE FROM users")
	
	server := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string {"POST", "GET", "UPDATE", "DELETE", "PUT"},
		AllowHeaders: []string {"Content-Type", "Authorization"},
		ExposeHeaders: []string {"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,

	}

	server.Use(cors.New(corsConfig))
	routes.RegisterRoutes(server)
	server.Run(":8080")
	

	// for _,item := range models.ProductsArray {
	// 	err := item.Save()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}

	// 	fmt.Println("Save success")
	// }

}