package schemavalidators

import (
	"errors"
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateEmail(email string) (bool, error) {
	emailSchema := `{
	"type":"object",
	"properties":{
	"email":{
	"type":"string",
	"format":"email"
			}
		},
	"required":["email"]
	}`
	emailData := fmt.Sprintf(`{"email": "%s"}`, email)
	schemaLoader := gojsonschema.NewStringLoader(emailSchema)
	dataLoader := gojsonschema.NewStringLoader(emailData)
	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		log.Println("Invalid email format", err.Error())
		return false, nil
	}
	if !result.Valid() {
		err := errors.New("invalid email format")
		log.Println("Invalid email")
		return false, err
	}

	log.Println("Valid schema ")
	return true, nil
}
