package models

import (
	"cars/pranay/github.com/config"
	"database/sql"
	"fmt"
)

type Car struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Year  int64   `json:"year"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

var Cars = make(map[int64]Car)

func (car *Car) Insert() {
	query := "INSERT INTO cars (name, model, year, brand, price) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	if err := config.DB.QueryRow(query, car.Name, car.Model, car.Year, car.Brand, car.Price).Scan(&car.Id); err != nil {
		fmt.Println("Error inserting car", err)
	}
}
func (car *Car) Get() error {
	query := "SELECT id, name, model, year, brand, price FROM cars WHERE id = $1"
	if err := config.DB.QueryRow(query, car.Id).Scan(&car.Id, &car.Name, &car.Model, &car.Year, &car.Brand, &car.Price); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Error getting car", err)
		}
		return err
	}
	return nil
}
func (car *Car) Update() {
	query := "UPDATE cars SET name = $1, model = $2, year = $3, brand = $4, price = $5 WHERE id = $6"
	if _, err := config.DB.Exec(query, car.Name, car.Model, car.Year, car.Brand, car.Price, car.Id); err != nil {
		fmt.Println("Error updating car", err)
	}
}

func (car *Car) Delete() error {
	query := "DELETE FROM cars WHERE id = $1"
	if _, err := config.DB.Exec(query, car.Id); err != nil {
		fmt.Println("Error deleting car", err)
		return err
	}
	return nil
}
