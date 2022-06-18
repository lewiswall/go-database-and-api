package datab

import (
	"database/sql"
	"log"
)

func UpdateCar(db *sql.DB, c CarComit, id string) error {

	result, err := db.Exec("update car set carBrandID = ?, carBodyID = ?, driveWheelID = ?, engineID = ?, "+
		"engineLocationID = ?, carName = ?, price = ?, wheelBase = ?, carLength = ?, carWidth = ?, carHeight = ?, curbWeight = ?, "+
		"doorNumber = ?, cityMPG = ?, highwayMPG = ? where carID = ?", c.BrandID, c.BodyID, c.DriveWheelID, c.EngineID,
		c.EngineLocationID, c.CarName, c.Price, c.WheelBase, c.CarLength, c.CarWidth, c.CarHeight,
		c.CurbWeight, c.DoorNumber, c.CityMPG, c.HighwayMPG, id)

	if err != nil {
		log.Println(err)
		return err
	}

	car, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Rows Affected insert : %d", car)

	return nil
}
