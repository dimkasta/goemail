package goemail

import "testing"

func TestEmail(t *testing.T) {
	email := NewHtmlMail()

	valid := email.IsValid()

	if valid {
		t.Errorf("A new email should not be valid")
	}

	valid = email.IsFromValid()
	if valid {
		t.Errorf("An empty From should not be valid")
	}

	senderAddress := "sender@email.com"
	senderName := "Sender Name"
	email.SetFrom(senderAddress, senderName)
	if email.From.Email != senderAddress || email.From.Name != senderName{
		t.Errorf("The sender was not set correctly")
	}

	valid = email.IsFromValid()

	if !valid {
		t.Errorf("A non empty From is valid")
	}

	valid = email.IsToValid()

	if valid {
		t.Errorf("An empty To is not valid")
	}

	address := "valid@email.com"
	name := "Valid Name"

	email.AddTo(address, name)

	if len(email.To) == 0 || len(email.ToList) == 0{
		t.Errorf("To not properly added")
	}

	if email.To[0].Name != name || email.To[0].Email != address || email.ToList[0] != address {
		t.Errorf("To Name or Email was not added")
	}

	valid = email.IsToValid()

	if !valid{
		t.Errorf("A non emtpy To should be valid")
	}

	valid = email.IsSubjectValid()

	if valid {
		t.Errorf("An empty subject should not be valid")
	}

	subject := "title"
	email.SetSubject(subject)

	if email.Subject != subject {
		t.Errorf("Subject was not set correctly")
	}

	valid = email.IsSubjectValid()

	if !valid {
		t.Errorf("A non-empty subject should be valid")
	}

	valid = email.IsBodyValid()

	if valid {
		t.Errorf("An empty body should not be valid")
	}

	body := "test"
	email.SetBody(body)

	if email.Body != body{
		t.Errorf("Body was not set correctly")
	}

	valid = email.IsBodyValid()

	if !valid {
		t.Errorf("A non-empty body should be valid")
	}

	valid = email.IsValid()

	if !valid {
		t.Errorf("A complete email should be valid")
	}




}
