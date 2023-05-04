package go_validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
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

// validation basic collection
func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string  `validate:"required,dive,required,min=3"`
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
		Hobbies: []string{
			"Gaming",
			"Coding",
			"X",
			"",
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// validation map
func TestValidationMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}
	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"X",
			"",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// validation basic map
func TestValidationBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
	}
	validate := validator.New()
	request := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"X",
			"",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA":     1000000,
			"MANDIRI": 0,
			"":        1001,
		},
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "12345",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}
	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "CUSTOM",
		Password: "",
	}
	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "0890129312",
		Pin:   "123123",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err)
	}
}
