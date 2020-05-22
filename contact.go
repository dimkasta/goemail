package goemail

import (
	"fmt"
)

type Contact struct {
	Email string
	Name  string
}

func NewContact(email string, name string) *Contact {
	return &Contact{
		Email: email,
		Name: name,
	}
}

func (contact *Contact)ToString() string {
	if contact.Name != ""{
		return fmt.Sprintf("%s <%s>", contact.Name, contact.Email)
	}

	return contact.Email
}