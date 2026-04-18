package res

import (
	"net/http"

	"github.com/atharvyadav96k/gcp/common"
)

// Send writes a JSON response to the client.
//
// Parameters:
//   - w: http response writer
//   - status: HTTP status code
//   - message: optional message
//   - data: response payload
//   - errs: optional error messages
func Send(w http.ResponseWriter, status int, message string, data interface{}, errs []string) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errs,
	}

	jsonData, err := common.ToJSON(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // ✅ FIXED
	w.Write(jsonData)
}
