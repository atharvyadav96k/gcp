package entity

// Email represents a user email address as a domain value.
//
// It is a custom type to encapsulate validation and behavior
// related to email handling (e.g., Validate, String).
type Email string

// Password represents a user password as a domain value.
//
// It supports validation, hashing, and comparison logic
// through associated methods.
type Password string

// PhoneNumber represents a structured phone number.
//
// Fields:
//   - CountryCode: international dialing code (e.g., "+91")
//   - Number: local phone number
//
// This type supports validation and formatting via methods.
type PhoneNumber struct {
	CountryCode string
	Number      string
}
