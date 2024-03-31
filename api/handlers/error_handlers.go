package handlers

import "net/http"

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		errorWriter(w, err.Error(), http.StatusBadRequest)
	}

	NotFoundErrorHandler = func(w http.ResponseWriter, err error) {
		errorWriter(w, err.Error(), http.StatusNotFound)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		errorWriter(w, "unexpected error occurred in server", http.StatusInternalServerError)
	}
)
