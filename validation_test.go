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

// Validation multiple tags
func TestValidationMultipleTags(t *testing.T) {
	validate := validator.New()

	name := "123123"

	err := validate.Var(name, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Tag parameter
func TestValidationTagParameter(t *testing.T) {
	validate := validator.New()
	num := "994444"

	err := validate.Var(num, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}
