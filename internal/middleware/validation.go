package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/YehorChervonyi/SimpleAPI/internal/models"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func ValidateInput(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var request models.FactorialRequest
		if err := json.Unmarshal(body, &request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if request.A < 0 || request.B < 0 {
			errorRes := map[string]string{"error": "Incorrect input"}
			jsonRes, _ := json.Marshal(errorRes)
			http.Error(w, string(jsonRes), http.StatusBadRequest)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		h(w, r, ps)
	}
}
