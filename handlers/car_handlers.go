package handlers

import (
	"cars/pranay/github.com/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

func CarHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	url := r.URL.String()
	fmt.Println("URL: ", url)
	fmt.Println("Path: ", path)
	entity := strings.TrimPrefix(path, "/cars")
	entity = strings.TrimPrefix(entity, "/")
	fmt.Println("Entity: ", entity)
	switch r.Method {
	case "GET":
		if entity != "" {
			id, _ := strconv.Atoi(entity)
			getCar(id, w)
		} else {
			http.Error(w, "We don't support GET method", http.StatusBadRequest)
		}
	case "POST":
		if entity == "" {
			createCar(r, w)
		} else {
			http.Error(w, "Incorrect post request", http.StatusBadRequest)
		}
	case "DELETE":
		if entity != "" {
			id, _ := strconv.Atoi(entity)
			deleteCar(id, w)
		} else {
			http.Error(w, "We don't support DELETE method", http.StatusBadRequest)
		}
	case "PUT":
		if entity == "" {
			updateCar(r, w)
		} else {
			http.Error(w, "We don't support PUT method", http.StatusBadRequest)
		}
	}
}

func createCar(r *http.Request, w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Incorrect post request", http.StatusBadRequest)
		return
	}
	car.Insert()
	fmt.Printf("Car saved with ID %v", car.Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func deleteCar(id int, w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	car.Id = int64(id)
	if err := car.Delete(); err != nil {
		fmt.Println("Error deleting car", err)
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}
	fmt.Printf("Car deleted successfully with ID %d", car.Id)
	w.WriteHeader(http.StatusNoContent)
}

func getCar(id int, w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	car.Id = int64(id)
	if err := car.Get(); err != nil {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func updateCar(r *http.Request, w http.ResponseWriter) {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Incorrect post request", http.StatusBadRequest)
		return
	}
	car.Update()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}
