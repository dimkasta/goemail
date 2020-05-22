package goemail

import (
	"github.com/dimkasta/gologger"
	"testing"
)

func TestEmailer(t *testing.T) {
	server  := "wrong:25"
	logger := gologger.NewLoggerService()
	mailer := NewMailer(logger, server)

	if mailer.Server != server{
		t.Errorf("There was an error setting the server")
	}

	email := NewHtmlMail()

	success := mailer.Send(email)

	if success {
		t.Errorf("Sending an empty email should have failed")
	}

	email.SetFrom("dimkasta@yahoo.gr", "Name")
	email.AddTo("to@iconic.gr", "Recipient Name")
	email.SetSubject("Test Subject")
	email.SetBody("body goes here")

	success = mailer.Send(email)

	if success {
		t.Errorf("Sending a valid email to the wrong server should fail")
	}

	mailer = NewMailer(logger, "localhost:1025")
	success = mailer.Send(email)

	if !success {
		t.Errorf("Sending should have been successful. Check mailhog container")
	}
}
