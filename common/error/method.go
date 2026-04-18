package common_error

// ErrorsToString converts a slice of error objects into a slice of strings.
//
// It extracts the error message from each error using err.Error().
//
// Parameters:
//   - errs: slice of error values
//
// Returns:
//   - []string: slice containing error messages
//
// Example:
//   errs := []error{
//       errors.New("invalid email"),
//       errors.New("password too short"),
//   }
//   messages := ErrorsToString(errs)
//   // ["invalid email", "password too short"]
func ErrorsToString(errs []error) []string {
	strErrs := make([]string, 0, len(errs)) // preallocate for performance

	for _, err := range errs {
		if err != nil {
			strErrs = append(strErrs, err.Error())
		}
	}

	return strErrs
}
