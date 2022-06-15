package datab

import (
	"database/sql"
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
