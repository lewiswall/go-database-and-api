package datab

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteCar(db *sql.DB, id string) error {
	query, err := db.Prepare("delete from car where carID = ?")
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := query.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	car, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(car)
	return nil
}
