package entity

// Validator defines a contract for types that can validate their own data.
//
// Any type implementing this interface should provide domain-specific
// validation logic inside the Validate method.
//
// Returns:
//   - error: nil if valid, otherwise a descriptive validation error
//
// Example:
//   type Email string
//
//   func (e Email) Validate() error {
//       // custom validation logic
//   }
//
//   var v Validator = Email("test@example.com")
//   err := v.Validate()
type validatable interface {
	Validate() error
}
