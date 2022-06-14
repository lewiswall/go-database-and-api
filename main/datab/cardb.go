package datab

import (
	"database/sql"
	"fmt"
)

type brand struct {
	Brand string `json:"name"`
}

func GetAllBrands(db *sql.DB) ([]brand, error) {
	results, err := db.Query("Select carBrandName FROM carBrand")
	if err != nil {
		return nil, err
	}

	defer results.Close()

	var brands []brand

	for results.Next() {
		var brand1 brand

		err = results.Scan(&brand1.Brand)

		if err != nil {
			panic(err.Error())
		}
		brands = append(brands, brand1)

	}
	return brands, nil
}

type car struct {
	ID             int     `json:"id"`
	CarBrand       string  `json:"car_brand"`
	CarBody        string  `json:"car_body"`
	DriveWheel     string  `json:"drive_wheel"`
	EngineID       int     `json:"engine_id"`
	EngineLocation string  `json:"engine_location"`
	CarName        string  `json:"car_name"`
	Price          float64 `json:"price"`
	WheelBase      int     `json:"wheel_base"`
	FuelType       string  `json:"fuel_type"`
	CarLength      float32 `json:"car_length"`
	CarWidth       float32 `json:"car_width"`
	CarHeight      float32 `json:"car_height"`
	CurbWeight     float32 `json:"curb_weight"`
	DoorNum        int     `json:"door_num"`
	CityMPG        int     `json:"city_mpg"`
	HighwayMPG     int     `json:"highway_mpg"`
}

func GetAllCars(db *sql.DB) ([]car, error) {
	query := "SELECT carID, carBrandName, carName, fuelTypeName, carBodyName, " +
		"DriveWheelName, engineLocationName, wheelBase, carLength, " +
		"carWidth, carHeight, curbWeight, doorNumber, cityMPG, highwayMPG, " +
		" engine.engineID , price, " +
		"FROM car, carBody, carBrand, engine, fuelType, aspiration, " +
		"engineType, fuelSystem, driveWheel, engineLocation " +
		"WHERE car.carBrandID = carBrand.carBrandID " +
		"AND car.engineID = engine.engineID " +
		"AND carBody.carBodyID = car.CarBodyID " +
		"AND fuelSystem.fuelSystemID = engine.fuelSystemID " +
		"AND driveWheel.driveWheelID = car.driveWheelID " +
		"AND aspiration.aspirationID = engine.aspirationID " +
		"AND engineType.engineTypeID = engine.engineTypeID " +
		"AND engineLocation.engineLocationID = car.engineLocationID " +
		"AND engine.fuelTypeID = fuelType.fuelTypeID " +
		"AND Car.carID = " + "1" + ";"

	results, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var car1 car

		err = results.Scan(&car1.ID, &car1.CarBrand, &car1.CarName, &car1.FuelType, &car1.CarBody, &car1.DriveWheel,
			&car1.EngineLocation, &car1.WheelBase, &car1.CarLength, &car1.CarWidth, &car1.CarHeight, &car1.CurbWeight,
			&car1.DoorNum, &car1.CityMPG, &car1.HighwayMPG, &car1.EngineID, &car1.Price)

		if err != nil {
			print(err)
			return nil, err
		}

		fmt.Println(car1)

	}
	return nil, nil
}

//func database() {
//	fmt.Println("Go Database API")
//
//	// Opening connection to database running on local machine
//	db, err := sql.Open("mysql", "root:lewiswall@tcp(localhost:3306)/go-rest-api")
//
//	//Handles error thrown as database is opening
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// Defer closing the database until the end of the current function
//	defer db.Close()
//
//	results, err := db.Query("Select carBrandName FROM carBrand")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	defer results.Close()
//
//	var cars []Car
//
//	for results.Next() {
//		var car Car
//
//		err = results.Scan(&car.Name)
//
//		if err != nil {
//			panic(err.Error())
//		}
//		cars = append(cars, car)
//
//	}
//
//	for _, c := range cars {
//		fmt.Println(c)
//	}
//
//}
