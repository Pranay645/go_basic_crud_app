package main

import (
	"cars/pranay/github.com/config"
	"cars/pranay/github.com/handlers"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()
	http.HandleFunc("/cars", handlers.CarHandler)
	http.HandleFunc("/cars/", handlers.CarHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":5173", nil)

}
