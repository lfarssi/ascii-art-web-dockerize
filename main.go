package main

import (
	"fmt"
	"net/http"

	"ascii-art/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiProcessHandler)

	
	fmt.Println("Server is running at http://localhost:8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
