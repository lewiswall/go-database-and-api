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
	var ret string
	for _, brand := range carBrands {
		ret += brand.Brand + "     "
	}

	c.JSON(http.StatusOK, ret)
}

func getAllCars(c *gin.Context) {

}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:lewiswall@tcp(localhost:3306)/go-rest-api")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := gin.Default()
	router.GET("/brands", getAllBrands)
	router.GET("/cars", getAllCars)

	fmt.Println("hello")
	router.Run("localhost:8080")
}
