package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Its Working")

}

func main() {

	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting Server at: /8080....")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {

		fmt.Println("Error starting server", nil)
	}

}
