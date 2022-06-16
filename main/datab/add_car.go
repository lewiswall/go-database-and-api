package datab

import (
	"database/sql"
	"fmt"
	"log"
)

type CarComit struct {
	CarID            int     `json:"car_id"`
	BrandID          int     `json:"brand_id"`
	BodyID           int     `json:"body_id"`
	DriveWheelID     int     `json:"drive_wheel_id"`
	EngineID         int     `json:"engine_id"`
	EngineLocationID int     `json:"engine_location_id"`
	CarName          string  `json:"car_name"`
	Price            string  `json:"price"`
	WheelBase        float64 `json:"wheel_base"`
	CarLength        float64 `json:"car_length"`
	CarWidth         float64 `json:"car_width"`
	CarHeight        float64 `json:"car_height"`
	CurbWeight       float64 `json:"curb_weight"`
	DoorNumber       int     `json:"door_number"`
	CityMPG          int     `json:"city_mpg"`
	HighwayMPG       int     `json:"highway_mpg"`
}

func AddCar(db *sql.DB, c CarComit) {
	query := "insert into car (carBrandID, carBodyID, driveWheelID, engineID, engineLocationID, carName, " +
		"price, wheelBase, carLength, carWidth, carHeight, curbWeight, doorNumber, cityMPG, highwayMPG) " +
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.BrandID, c.BodyID, c.DriveWheelID, c.EngineID, c.EngineLocationID, c.CarName,
		c.Price, c.WheelBase, c.CarLength, c.CarWidth, c.CarHeight, c.CurbWeight, c.DoorNumber, c.CityMPG,
		c.HighwayMPG)

	if err != nil {
		log.Println(err)
	}

	car, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(car)

}

//{
//"car_id":0,
//"brand_id":1,
//"body_id":1,
//"drive_wheel_id":1,
//"engine_id":1,
//"engine_location_id":1,
//"car_name":"Lewiss",
//"price":"£100,000",
//"wheel_base":20,
//"car_length":20,
//"car_width":20,
//"car_height":20,
//"curb_weight":20,
//"door_number":5,
//"city_mpg":18,
//"highway_mpg":18
//}
