package compaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email"`
}

type Compaign struct {
	ID        string
	Name      string `validate:"required,min=3,max=100"`
	Content   string `validate:"required,min=3,max=1000"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Contacts  []Contact `validate:"required,min=1,max=100"`
}

func NewCompaign(name, content string, contacts []string) (*Compaign, error) { // create a new compaign

	if name == "" {
		return nil, errors.New("name is required")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(contacts) == 0 {
		return nil, errors.New("contacts are required")
	}

	contactEmails := make([]Contact, len(contacts)) // create a new slice of contacts
	for i, contact := range contacts {
		contactEmails[i].Email = contact // add the email to the slice
	}

	return &Compaign{
		ID:        xid.New().String(), // generate a new UUID for the compaign
		Name:      name,
		Content:   content,
		Contacts:  contactEmails,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
