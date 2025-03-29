package main

import (
	"calculator/calculator"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestData struct {
	Number1  float32 `json:"number1"`
	Number2  float32 `json:"number2"`
	Operator string  `json:"operator"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/calculate?", calculateHandler).Methods("POST")

	fmt.Print("Server is up and running")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Error in starting server:", err)
	}
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("%v", err)})
		}
	}()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData RequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	fmt.Println("Request Data", requestData)

	var number1 float32 = requestData.Number1
	var number2 float32 = requestData.Number2
	var operator string = requestData.Operator

	var answer float32 = calculator.Calculator(number1, operator, number2)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)

}
