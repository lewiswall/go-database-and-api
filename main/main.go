package main

import (
	"database-api/datab"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func getAllCars(c *gin.Context) {
	car := datab.GetAllCars(db)
	fmt.Println(car)
}

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:lewiswall@tcp(localhost:3306)/go-rest-api")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := gin.Default()
	router.GET("/cars", getAllCars)

	fmt.Println("hello")
	router.Run("localhost:8080")
}
