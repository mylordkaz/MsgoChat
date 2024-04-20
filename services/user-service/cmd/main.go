package main

import (
	"fmt"
	"net/http"
)



func main () {
	http.HandleFunc("/", handleUser)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've reached the user service")
}