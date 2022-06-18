package datab

import (
	"database/sql"
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

func AddCar(db *sql.DB, c CarComit) error {
	query := "insert into car (carBrandID, carBodyID, driveWheelID, engineID, engineLocationID, carName, " +
		"price, wheelBase, carLength, carWidth, carHeight, curbWeight, doorNumber, cityMPG, highwayMPG) " +
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(c.BrandID, c.BodyID, c.DriveWheelID, c.EngineID, c.EngineLocationID, c.CarName,
		c.Price, c.WheelBase, c.CarLength, c.CarWidth, c.CarHeight, c.CurbWeight, c.DoorNumber, c.CityMPG,
		c.HighwayMPG)

	if err != nil {
		log.Println(err)
		return err
	}

	car, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Rows Affected insert : %d", car)

	return nil
}
