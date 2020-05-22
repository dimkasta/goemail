package goemail

type Mail struct {
	From    *Contact
	To      []*Contact
	ToList  []string
	Subject string
	Body    string
	Mime    string
}

func NewHtmlMail() *Mail {
	return &Mail{
		To: []*Contact{},
		Mime: "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n",
	}
}

func (mail *Mail) AddTo(email string, name string) {
	mail.To = append(mail.To, NewContact(email, name))
	mail.ToList = append(mail.ToList, email)
}

func (mail *Mail) SetBody(body string) {
	mail.Body = body
}

func (mail *Mail) SetFrom(email string, name string) {
	mail.From = NewContact(email, name)
}

func (mail *Mail) SetSubject(subject string) {
	mail.Subject = subject
}

func (email *Mail) IsFromValid() bool {
	return email.From != nil
}

func (email *Mail) IsToValid() bool {
	return len(email.To) > 0 && len(email.ToList) > 0
}

func (email *Mail) IsSubjectValid() bool{
	return email.Subject != ""
}

func (email *Mail) IsBodyValid() bool {
	return email.Body != ""
}

func (email *Mail) IsValid() bool {
	return email.IsFromValid() && email.IsToValid() && email.IsSubjectValid() && email.IsBodyValid()

}