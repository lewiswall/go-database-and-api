package datab

import (
	"database/sql"
	"fmt"
)

type car struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllCars(db *sql.DB) car {
	fmt.Print("working")
	return car{}
}

// func database() {
// 	fmt.Println("Go Database API")

// 	// Opening connection to database running on local machine
// 	db, err := sql.Open("mysql", "root:lewiswall@tcp(localhost:3306)/go-rest-api")

// 	//Handles error thrown as database is opening
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// Defer closing the database until the end of the current function
// 	defer db.Close()

// 	results, err := db.Query("Select carBrandName FROM carBrand")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer results.Close()

// 	var cars []Car

// 	for results.Next() {
// 		var car Car

// 		err = results.Scan(&car.Name)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		cars = append(cars, car)

// 	}

// 	for _, c := range cars {
// 		fmt.Println(c)
// 	}

// }
