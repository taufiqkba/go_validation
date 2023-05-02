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

// validation struct
func TestStruct(t *testing.T) {
	type LoginStruct struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()
	loginRequest := LoginStruct{
		Username: "masuk@admin.com",
		Password: "masuk",
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// validation errors
func TestValidationErrors(t *testing.T) {
	type LoginStruct struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()
	loginRequest := LoginStruct{
		Username: "masuk",
		Password: "mas",
	}
	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

// validation cross field
func TestValidationCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}
	validate := validator.New()
	registerUser := RegisterUser{
		Username:        "test@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}
	err := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// validation nested struct
func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}
	validate := validator.New()
	request := User{
		Id:   "1",
		Name: "myName",
		Address: Address{
			City:    "Semarang",
			Country: "Indonesia",
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// validation collection
func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
	}
	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
