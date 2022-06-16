package datab

import (
	"database/sql"
	"errors"
	"log"
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
		log.Fatal(err)
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
			log.Println(err)
			return nil, err
		}
		cars = append(cars, car1)
	}
	return cars, nil
}

func CheckForID(db *sql.DB, id string) error {
	query := "SELECT COUNT(carID) FROM car WHERE carID = " + id
	result, err := db.Query(query)

	if err != nil {
		log.Println(err)
		return err
	}
	defer result.Close()

	for result.Next() {
		var num int

		err := result.Scan(&num)
		if err != nil {
			log.Println(err)
			return err
		}

		if num < 1 {
			log.Println("Passed an ID which is not contained in the database")
			return errors.New("Passed an ID which is not contained in the database")
		}
	}

	return nil
}

func GetCarByID(db *sql.DB, id string) (car, error) {

	query := "select carID, carBrandName, carName,fuelTypeName ,carBodyName, driveWheelName, engineLocationName, " +
		"wheelBase, carLength, carWidth, carHeight, curbWeight, doorNumber, cityMPG, highwayMPG, " +
		"car.engineID, price " +
		"from car, carBrand, engine, fuelType, carBody, driveWheel, engineLocation " +
		"where car.carBrandID = carBrand.carBrandID and car.engineID = engine.engineID and " +
		"engine.fuelTypeID = fuelType.fuelTypeID and car.carBodyID = carBody.carBodyID and " +
		"car.driveWheelID = driveWheel.driveWheelID and car.engineLocationID = engineLocation.engineLocationID and " +
		"car.carID = " + id

	result, err := db.Query(query)
	defer result.Close()

	if err != nil {
		log.Println(err)
		return car{}, err
	}

	var car1 car
	for result.Next() {
		err = result.Scan(&car1.ID, &car1.CarBrand, &car1.CarName, &car1.FuelType, &car1.CarBody,
			&car1.DriveWheel, &car1.EngineLocation, &car1.WheelBase, &car1.CarLength, &car1.CarWidth,
			&car1.CarHeight, &car1.CurbWeight, &car1.DoorNum, &car1.CityMPG, &car1.HighwayMPG,
			&car1.EngineID, &car1.Price)

		if err != nil {
			log.Println(err)
			return car{}, err
		}
	}
	return car1, nil
}
