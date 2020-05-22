package goemail

import (
	"fmt"
	"testing"
)

func TestNewValidFullContact(t *testing.T) {
	email := "valid@email.com"
	name := "Valid Name"
	contact := NewContact(email, name)

	if contact.Email != email || contact.Name != name{
		t.Errorf("Email or Name are invalid %s - %s", email, name)
	}

	if contact.ToString() != fmt.Sprintf("%s <%s>", name, email){
		t.Errorf("Contact ToString format is not as expected")
	}
}

func TestValidContactWithoutName(t *testing.T) {
	email := "valid@email.com"
	name := ""
	contact := NewContact(email, name)

	if contact.Email != email || contact.Name != name{
		t.Errorf("Email or Name are invalid %s - %s", email, name)
	}

	if contact.ToString() != email {
		t.Errorf("Contact ToString format is not as expected")
	}
}
