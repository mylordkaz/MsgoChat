package handlers

import (
	"fmt"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, you've reached the user service")
}