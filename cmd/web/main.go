package main

import (
	"fmt"
	"net/http"

	"github.com/robertbenavidez/bookings/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting Application on port", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
