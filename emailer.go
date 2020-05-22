package goemail

import (
	"github.com/dimkasta/gologger"
	"go.uber.org/zap"
	"net/smtp"
)

type Mailer struct {
	Server string
	Logger *gologger.LoggerService
}

func NewMailer(logger *gologger.LoggerService, server string) *Mailer {
	// Choose auth method and set it up
	//auth := smtp.PlainAuth("", "piotr@mailtrap.io", "extremely_secret_pass", "smtp.mailtrap.io")
	logger.Log.Debug("Creating New Mailer")
	return &Mailer{
		Logger: logger,
		Server: server,
	}
}

func (mailer Mailer) Send(email *Mail) bool {
	mailer.Logger.Log.Debug("Checking Email")
	if ! email.IsValid() {
		mailer.Logger.Log.Debug("Email is not Valid. Aborting")
		return false
	}
	mailer.Logger.Log.Debug("Sending Email")
	mailer.Logger.Log.Debug("Subject: " + email.Subject)
	mailer.Logger.Log.Debug("From: " + email.From.ToString())
	mailer.Logger.Log.Debug("To: " + email.To[0].ToString())
	err := smtp.SendMail(mailer.Server, nil, email.From.ToString(),  email.ToList,
		[]byte(createMessage(email.From.ToString(), email.To[0].ToString(), email.Subject, email.Mime, email.Body)))

	if err != nil {
		mailer.Logger.Log.Error(err.Error())
		return false
	}
	mailer.Logger.Log.Info("Mail sent", zap.String("From",email.From.ToString()), zap.String("To", email.To[0].ToString()), zap.String("Subject", email.Subject))
	return true
}



func createMessage(from string, to string, title string, mime string, body string) string {
	return "From: " + from + "\n" +
		"To: " + to + "\n" + "Subject: " + title + "\n" + mime + body
}
