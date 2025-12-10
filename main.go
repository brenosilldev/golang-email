package main

import (
	"emailn/internal/domain/compaign"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	compaignService := compaign.Compaign{
		Name: "Test Compaign",
		Content: "This is a test compaign",
		Contacts: []compaign.Contact{
			{Email: "test@example.com"},
			{Email: "test2@example.com"},
			{Email: "test3@example.com"},
		},
	}

	validate := validator.New()

	err := validate.Struct(compaignService)

	if err != nil {
		fmt.Println(err)
	} 
}
