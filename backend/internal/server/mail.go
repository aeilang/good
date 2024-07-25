package server

import "gopkg.in/gomail.v2"

type Mail struct {
	dialer *gomail.Dialer
}

type Mialer interface {
	Send(message Message) error
}

var _ Mialer = (*Mail)(nil)

type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
	Attach  []string
}

func (m *Mail) Send(message Message) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", message.From)
	msg.SetHeader("To", message.To...)
	msg.SetHeader("Subject", message.Subject)
	msg.SetBody("text/html", message.Body)

	for _, name := range message.Attach {
		msg.Attach(name)
	}

	return m.dialer.DialAndSend(msg)
}
