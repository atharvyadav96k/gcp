package test

import (
	"os"
	"path/filepath"
	"testing"

	"encoding/json"

	"github.com/atharvyadav96k/gcp/common/entity"
)

// EntityValidationCase represents a single entity validation test case
type EntityValidationCase struct {
	Name          string `json:"name"`
	EntityType    string `json:"entity_type"`
	Input         string `json:"input,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	Number        string `json:"number,omitempty"`
	ExpectedValid bool   `json:"expected_valid"`
	ErrorMessage  string `json:"error_message,omitempty"`
}

// EntityValidationCases wraps the list of test cases
type EntityValidationCases struct {
	Cases []EntityValidationCase `json:"entity_validation_cases"`
}

// loadEntityValidationCases loads test cases from the JSON file
func loadEntityValidationCases(t *testing.T) []EntityValidationCase {
	testTablePath := filepath.Join("..", "test_table", "entity_validation_cases.json")
	file, err := os.Open(testTablePath)
	if err != nil {
		t.Fatalf("Failed to open entity validation test cases file: %v", err)
	}
	defer file.Close()

	var tc EntityValidationCases
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tc); err != nil {
		t.Fatalf("Failed to decode entity validation test cases: %v", err)
	}

	return tc.Cases
}

// TestEmailValidationFromTable tests email validation using table-driven tests
func TestEmailValidationFromTable(t *testing.T) {
	testCases := loadEntityValidationCases(t)

	for _, tc := range testCases {
		if tc.EntityType != "email" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			email := entity.Email(tc.Input)
			err := email.Validate()

			if tc.ExpectedValid && err != nil {
				t.Errorf("Expected valid email but got error: %v", err)
			}

			if !tc.ExpectedValid {
				if err == nil {
					t.Errorf("Expected validation error but got none")
				} else if tc.ErrorMessage != "" && err.Error() != tc.ErrorMessage {
					t.Errorf("Error message mismatch. Got: %q, Want: %q", err.Error(), tc.ErrorMessage)
				}
			}
		})
	}
}

// TestPasswordValidationFromTable tests password validation using table-driven tests
func TestPasswordValidationFromTable(t *testing.T) {
	testCases := loadEntityValidationCases(t)

	for _, tc := range testCases {
		if tc.EntityType != "password" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			password := entity.Password(tc.Input)
			err := password.Validate()

			if tc.ExpectedValid && err != nil {
				t.Errorf("Expected valid password but got error: %v", err)
			}

			if !tc.ExpectedValid {
				if err == nil {
					t.Errorf("Expected validation error but got none")
				} else if tc.ErrorMessage != "" && err.Error() != tc.ErrorMessage {
					t.Errorf("Error message mismatch. Got: %q, Want: %q", err.Error(), tc.ErrorMessage)
				}
			}
		})
	}
}

// TestPhoneNumberValidationFromTable tests phone number validation using table-driven tests
func TestPhoneNumberValidationFromTable(t *testing.T) {
	testCases := loadEntityValidationCases(t)

	for _, tc := range testCases {
		if tc.EntityType != "phone_number" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			phoneNumber := entity.PhoneNumber{
				CountryCode: tc.CountryCode,
				Number:      tc.Number,
			}
			err := phoneNumber.Validate()

			if tc.ExpectedValid && err != nil {
				t.Errorf("Expected valid phone number but got error: %v", err)
			}

			if !tc.ExpectedValid {
				if err == nil {
					t.Errorf("Expected validation error but got none")
				} else if tc.ErrorMessage != "" && err.Error() != tc.ErrorMessage {
					t.Errorf("Error message mismatch. Got: %q, Want: %q", err.Error(), tc.ErrorMessage)
				}
			}
		})
	}
}

// TestNewEmailFromTable tests NewEmail constructor using table-driven tests
func TestNewEmailFromTable(t *testing.T) {
	testCases := loadEntityValidationCases(t)

	for _, tc := range testCases {
		if tc.EntityType != "new_email" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			_, err := entity.NewEmail(tc.Input)

			if tc.ExpectedValid && err != nil {
				t.Errorf("Expected valid email creation but got error: %v", err)
			}

			if !tc.ExpectedValid {
				if err == nil {
					t.Errorf("Expected creation error but got none")
				} else if tc.ErrorMessage != "" && err.Error() != tc.ErrorMessage {
					t.Errorf("Error message mismatch. Got: %q, Want: %q", err.Error(), tc.ErrorMessage)
				}
			}
		})
	}
}

// TestNewPasswordFromTable tests NewPassword constructor using table-driven tests
func TestNewPasswordFromTable(t *testing.T) {
	testCases := loadEntityValidationCases(t)

	for _, tc := range testCases {
		if tc.EntityType != "new_password" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			_, err := entity.NewPassword(tc.Input)

			if tc.ExpectedValid && err != nil {
				t.Errorf("Expected valid password creation but got error: %v", err)
			}

			if !tc.ExpectedValid {
				if err == nil {
					t.Errorf("Expected creation error but got none")
				} else if tc.ErrorMessage != "" && err.Error() != tc.ErrorMessage {
					t.Errorf("Error message mismatch. Got: %q, Want: %q", err.Error(), tc.ErrorMessage)
				}
			}
		})
	}
}

// TestNewPhoneNumberFromTable tests NewPhoneNumber constructor using table-driven tests
func TestNewPhoneNumberFromTable(t *testing.T) {
	testCases := loadEntityValidationCases(t)

	for _, tc := range testCases {
		if tc.EntityType != "new_phone_number" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			_, err := entity.NewPhoneNumber(tc.CountryCode, tc.Number)

			if tc.ExpectedValid && err != nil {
				t.Errorf("Expected valid phone number creation but got error: %v", err)
			}

			if !tc.ExpectedValid {
				if err == nil {
					t.Errorf("Expected creation error but got none")
				} else if tc.ErrorMessage != "" && err.Error() != tc.ErrorMessage {
					t.Errorf("Error message mismatch. Got: %q, Want: %q", err.Error(), tc.ErrorMessage)
				}
			}
		})
	}
}
