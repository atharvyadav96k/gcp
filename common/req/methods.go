package req

import (
	"encoding/json"
	"net/http"
)

// ParseBody decodes the JSON request body into the provided destination struct.
//
// Parameters:
//   - r: incoming HTTP request
//   - v: pointer to the struct where decoded data will be stored
//
// Returns:
//   - error if decoding fails or body is invalid
//
// NOTE:
//   - v must be a pointer, otherwise decoding will fail
//   - the request body is closed after decoding
//
// Example:
//
//	var body CreateUserRequest
//	err := ParseBody(r, &body)
func ParseBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(v)
}
