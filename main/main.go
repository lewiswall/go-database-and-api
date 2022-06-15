package main

import (
	"database-api/datab"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func getAllBrands(c *gin.Context) {
	carBrands, err := datab.GetAllBrands(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "There was an error while trying to query the database")
	}
	c.JSON(http.StatusOK, carBrands)
}

func getAllCars(c *gin.Context) {
	cars, err := datab.GetAllCars(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "There was an error while trying to query the database")
	}
	c.JSON(http.StatusOK, cars)
}

func getCarByID(c *gin.Context) {
	params := c.Request.URL.Query()
	id := params.Get("id")

	car, err := datab.GetCarByID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "There was an error while trying to query the database")
	}
	c.JSON(http.StatusOK, car)
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:lewiswall@tcp(localhost:3306)/go-rest-api")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/brands", getAllBrands)
	router.GET("/cars", getAllCars)
	router.GET("/car", getCarByID)

	fmt.Println("Server Starting")
	router.Run("localhost:8080")
}
