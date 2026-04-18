package entity

import (
	"errors"
	"net/mail"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// ValidateStruct validates all fields of a struct that implement the Validator interface.
//
// It iterates over struct fields and calls Validate() on each field that satisfies
// the Validator interface (pointer or value receiver).
//
// Parameters:
//   - s: struct or pointer to struct to validate
//
// Returns:
//   - []error: list of validation errors (empty if all fields are valid)
//
// Example:
//
//	errs := ValidateStruct(user)
//	if len(errs) > 0 {
//	    // handle validation errors
//	}
func ValidateStruct(s interface{}) []error {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return []error{errors.New("input must be a struct")}
	}

	var errList []error

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		var valid validatable

		// Try pointer receiver
		if field.CanAddr() {
			if v, ok := field.Addr().Interface().(validatable); ok {
				valid = v
			}
		}

		// Try value receiver
		if valid == nil {
			if v, ok := field.Interface().(validatable); ok {
				valid = v
			}
		}

		if valid != nil {
			if err := valid.Validate(); err != nil {
				errList = append(errList, err)
			}
		}
	}

	return errList
}

// String returns the string representation of Email.
func (e *Email) String() string {
	return string(*e)
}

// Validate checks if the Email is properly formatted.
//
// It ensures:
//   - not empty
//   - valid RFC format
//   - proper domain structure
//
// Returns:
//   - error if email is invalid
func (e *Email) Validate() error {
	emailStr := strings.TrimSpace(string(*e))
	if emailStr == "" {
		return errors.New("email address cannot be empty")
	}

	addr, err := mail.ParseAddress(emailStr)
	if err != nil {
		return errors.New("invalid email format")
	}

	if addr.Address != emailStr {
		return errors.New("email contains invalid characters or extra spaces")
	}

	parts := strings.Split(emailStr, "@")
	if len(parts) != 2 {
		return errors.New("invalid email structure")
	}

	domain := parts[1]
	if !strings.Contains(domain, ".") ||
		strings.HasPrefix(domain, ".") ||
		strings.HasSuffix(domain, ".") {
		return errors.New("email domain is invalid")
	}

	return nil
}

// HashPassword hashes the Password using bcrypt.
//
// It replaces the plain password with its hashed version.
//
// Returns:
//   - error if hashing fails
//
// Example:
//
//	err := p.HashPassword()
func (p *Password) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*p), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*p = Password(hashedPassword)
	return nil
}

// Validate checks if the Password meets strength requirements.
//
// Rules:
//   - minimum length: 8
//   - maximum length: 72 (bcrypt limit)
//   - must contain uppercase, lowercase, and number
//
// Returns:
//   - error if password is weak or invalid
func (p *Password) Validate() error {
	pStr := string(*p)
	if len(pStr) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if len(pStr) > 72 {
		return errors.New("password is too long (maximum 72 characters)")
	}

	var hasUpper, hasLower, hasNumber bool
	for _, char := range pStr {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	return nil
}

// ComparePassword compares a hashed password with a plain password.
//
// Parameters:
//   - plainPassword: raw password input
//
// Returns:
//   - error if password does not match or comparison fails
//
// Example:
//
//	err := storedPassword.ComparePassword("user-input")
func (p Password) ComparePassword(plainPassword string) error {
	if p == "" {
		return errors.New("hashed password is empty")
	}

	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(plainPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("invalid password")
		}
		return err
	}

	return nil
}

// String returns the full phone number (country code + number).
func (pn *PhoneNumber) String() string {
	return pn.CountryCode + pn.Number
}

// Validate checks if the PhoneNumber is valid.
//
// Rules:
//   - country code must start with '+' and be at least 2 chars
//   - number must contain only digits (non-digits are ignored)
//   - length must be between 6 and 15 digits
//
// Returns:
//   - error if phone number is invalid
func (pn *PhoneNumber) Validate() error {
	if !strings.HasPrefix(pn.CountryCode, "+") || len(pn.CountryCode) < 2 {
		return errors.New("invalid country code (must start with +)")
	}

	cleanNumber := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, pn.Number)

	if len(cleanNumber) < 6 {
		return errors.New("phone number is too short")
	}
	if len(cleanNumber) > 15 {
		return errors.New("phone number is too long")
	}

	return nil
}
