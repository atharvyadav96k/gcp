package res

import "net/http"

// Success returns a 200 OK response.
func Success(w http.ResponseWriter, message string, data interface{}) {
	Send(w, http.StatusOK, message, data, nil)
}

// Created returns a 201 Created response.
func Created(w http.ResponseWriter, message string, data interface{}) {
	Send(w, http.StatusCreated, message, data, nil)
}

// NoContent returns a 204 No Content response.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
