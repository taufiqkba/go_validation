package go_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

// Validation variable
func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "name"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation two variable
func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "secret"
	confirmPassword := "secret"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}
