package main

import (
	"cars/pranay/github.com/config"
	"cars/pranay/github.com/handlers"
	"cars/pranay/github.com/middleware"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()
	mux := http.NewServeMux()

	mux.HandleFunc("/cars", handlers.CarHandler)
	mux.HandleFunc("/cars/", handlers.CarHandler)
	wrappedMux := middleware.Logger(mux)
	wrappedMux = middleware.SecurityHeaders(wrappedMux)
	fmt.Println("Server is running on port 5173")
	http.ListenAndServe(":5173", wrappedMux)

}
