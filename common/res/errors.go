package res

import (
	"net/http"

	common_error "github.com/atharvyadav96k/gcp/common/error"
)

// InternalServerError returns a 500 response.
func InternalServerError(w http.ResponseWriter, err []error) {
	Send(w, http.StatusInternalServerError, "Internal Server Error", nil, common_error.ErrorsToString(err))
}

// BadRequest returns a 400 response.
func BadRequest(w http.ResponseWriter, err []error) {
	Send(w, http.StatusBadRequest, "Bad Request", nil, common_error.ErrorsToString(err))
}

// NotFound returns a 404 response.
func NotFound(w http.ResponseWriter, err []error) {
	Send(w, http.StatusNotFound, "Not Found", nil, common_error.ErrorsToString(err))
}

// Forbidden returns a 403 response.
func Forbidden(w http.ResponseWriter, err []error) {
	Send(w, http.StatusForbidden, "Forbidden", nil, common_error.ErrorsToString(err))
}

// Unauthorized returns a 401 response.
func Unauthorized(w http.ResponseWriter, err []error) {
	Send(w, http.StatusUnauthorized, "Unauthorized", nil, common_error.ErrorsToString(err))
}
