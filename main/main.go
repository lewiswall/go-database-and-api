package main

import (
	"database-api/datab"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	checkCarID(c, id)

	car, err := datab.GetCarByID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "There was an error while trying to query the database")
	}
	c.JSON(http.StatusOK, car)
}

func checkCarID(c *gin.Context, id string) {
	err := datab.CheckForID(db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Passed ID is not available")
	}
}

func delCarByID(c *gin.Context) {
	params := c.Request.URL.Query()
	id := params.Get("id")
	checkCarID(c, id)

	err := datab.DeleteCar(db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "There was a problem deleting the car")
	}

}

func addCar(c *gin.Context) {
	var newCar datab.CarComit
	if err := c.BindJSON(&newCar); err != nil {
		log.Println(err)
		return
	}

	err := datab.AddCar(db, newCar)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to insert car into database. Check your inputs.")
		return
	}
	newCar.CarID = 000
	c.JSON(http.StatusOK, newCar)
}

func updateCar(c *gin.Context) {
	params := c.Request.URL.Query()
	id := params.Get("id")
	checkCarID(c, id)

	var updatedCar datab.CarComit
	if err := c.BindJSON(&updatedCar); err != nil {
		log.Println(err)
		return
	}

	err := datab.UpdateCar(db, updatedCar, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Could not update car")
	}
	c.JSON(http.StatusOK, "Car Updated")
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
	router.DELETE("/del-car", delCarByID)
	router.POST("/add-car", addCar)
	router.PUT("/update-car", updateCar)

	fmt.Println("Server Starting")
	router.Run("localhost:8080")
}

//{
//"car_id": 0,
//"brand_id": 1,
//"body_id": 1,
//"drive_wheel_id": 1,
//"engine_id": 1,
//"engine_location_id": 1,
//"car_name": "Lewiss",
//"price": "£100,000",
//"wheel_base": 20,
//"car_length": 20,
//"car_width": 20,
//"car_height": 20,
//"curb_weight": 20,
//"door_number": 5,
//"city_mpg": 18,
//"highway_mpg": 18
//}
