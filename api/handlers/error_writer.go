package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AdityaMayukhSom/alex_mux_go/api/types"
)

func errorWriter(w http.ResponseWriter, msg string, code int) {
	var resp types.Error = types.Error{
		Code:    code,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}
