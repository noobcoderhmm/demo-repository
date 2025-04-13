package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Light-Yagamii/demo-repository/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/validate-email", ValidateEmailHandler).Methods("GET")
	r.HandleFunc("/sum", SumHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func ValidateEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	valid := utils.IsEmailValid(email)

	json.NewEncoder(w).Encode(map[string]bool{"valid": valid})
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid numbers", http.StatusBadRequest)
		return
	}

	sum := utils.CalculateSum(a, b)
	json.NewEncoder(w).Encode(map[string]int{"sum": sum})
}
