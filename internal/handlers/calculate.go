package handlers

import (
	"encoding/json"
	"github.com/YehorChervonyi/SimpleAPI/internal/factorial"
	"github.com/YehorChervonyi/SimpleAPI/internal/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
)

func CalculateFactorial(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request models.FactorialRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var aFactorial, bFactorial int
	go func() {
		defer wg.Done()
		aFactorial = factorial.Calculate(request.A)
	}()
	go func() {
		defer wg.Done()
		bFactorial = factorial.Calculate(request.B)
	}()

	wg.Wait()

	response := models.FactorialResponse{
		AFactorial: aFactorial,
		BFactorial: bFactorial,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
