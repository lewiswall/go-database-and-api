package datab

import (
	"database/sql"
	"fmt"
)

type car struct {
	ID             int     `json:"id"`
	CarBrand       string  `json:"car_brand"`
	CarName        string  `json:"car_name"`
	FuelType       string  `json:"fuel_type"`
	CarBody        string  `json:"car_body"`
	DriveWheel     string  `json:"drive_wheel"`
	EngineLocation string  `json:"engine_location"`
	WheelBase      float64 `json:"wheel_base"`
	CarLength      float32 `json:"car_length"`
	CarWidth       float32 `json:"car_width"`
	CarHeight      float32 `json:"car_height"`
	CurbWeight     float32 `json:"curb_weight"`
	DoorNum        int     `json:"door_num"`
	CityMPG        int     `json:"city_mpg"`
	HighwayMPG     int     `json:"highway_mpg"`
	EngineID       int     `json:"engine_id"`
	Price          string  `json:"price"`
}

func GetAllCars(db *sql.DB) ([]car, error) {
	query := "select carID, carBrandName, carName,fuelTypeName ,carBodyName, driveWheelName, engineLocationName, " +
		"wheelBase, carLength, carWidth, carHeight, curbWeight, doorNumber, cityMPG, highwayMPG, " +
		"car.engineID, price " +
		"from car, carBrand, engine, fuelType, carBody, driveWheel, engineLocation " +
		"where car.carBrandID = carBrand.carBrandID and car.engineID = engine.engineID and " +
		"engine.fuelTypeID = fuelType.fuelTypeID and car.carBodyID = carBody.carBodyID and " +
		"car.driveWheelID = driveWheel.driveWheelID and car.engineLocationID = engineLocation.engineLocationID"

	results, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	var cars []car

	for results.Next() {
		var car1 car

		err = results.Scan(&car1.ID, &car1.CarBrand, &car1.CarName, &car1.FuelType, &car1.CarBody,
			&car1.DriveWheel, &car1.EngineLocation, &car1.WheelBase, &car1.CarLength, &car1.CarWidth,
			&car1.CarHeight, &car1.CurbWeight, &car1.DoorNum, &car1.CityMPG, &car1.HighwayMPG,
			&car1.EngineID, &car1.Price)

		if err != nil {
			return nil, err
		}
		cars = append(cars, car1)
	}
	return cars, nil
}

func GetCarByID(db *sql.DB, id string) (car, error) {
	query := "select carID, carBrandName, carName,fuelTypeName ,carBodyName, driveWheelName, engineLocationName, " +
		"wheelBase, carLength, carWidth, carHeight, curbWeight, doorNumber, cityMPG, highwayMPG, " +
		"car.engineID, price " +
		"from car, carBrand, engine, fuelType, carBody, driveWheel, engineLocation " +
		"where car.carBrandID = carBrand.carBrandID and car.engineID = engine.engineID and " +
		"engine.fuelTypeID = fuelType.fuelTypeID and car.carBodyID = carBody.carBodyID and " +
		"car.driveWheelID = driveWheel.driveWheelID and car.engineLocationID = engineLocation.engineLocationID and" +
		"carID = " + id

	result, err := db.Query(query)
	if err != nil {
		return car{}, err
	}
	defer result.Close()

	fmt.Println("here")

	var car1 car

	err = result.Scan(&car1.ID, &car1.CarBrand, &car1.CarName, &car1.FuelType, &car1.CarBody,
		&car1.DriveWheel, &car1.EngineLocation, &car1.WheelBase, &car1.CarLength, &car1.CarWidth,
		&car1.CarHeight, &car1.CurbWeight, &car1.DoorNum, &car1.CityMPG, &car1.HighwayMPG,
		&car1.EngineID, &car1.Price)

	if err != nil {
		return car{}, err
	}
	return car1, nil
}
