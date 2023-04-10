package builder

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (e * EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	e.email.from = from
	return e
}

func (e * EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("email should contain @")
	}
	e.email.to = to
	return e
}

func (e * EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}

func (e * EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

func sendEmailInt(email *email) {
	fmt.Println("BUILDER -> Parameter", email)
}

type build func(*EmailBuilder) 
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailInt(&builder.email)
}
