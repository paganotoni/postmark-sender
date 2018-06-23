package sender

import (
	"errors"
	"strings"
	"sync"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/keighl/postmark"
)

//PostmarkSender implements the Sender interface to be used
//within buffalo mailer generated package.
type PostmarkSender struct {
	mutex  sync.RWMutex
	client *postmark.Client

	trackOpens bool
}

//Send sends an email to Postmark for delivery, it assumes
//bodies[0] is HTML body and bodies[1] is text.
func (ps *PostmarkSender) Send(m mail.Message) error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	if len(m.Bodies) < 2 {
		return errors.New("you must specify at least 2 bodies HTML and plain text")
	}

	email := postmark.Email{
		From:       m.From,
		To:         strings.Join(m.To, ","),
		Subject:    m.Subject,
		HtmlBody:   m.Bodies[0].Content,
		TextBody:   m.Bodies[1].Content,
		TrackOpens: ps.trackOpens,
	}

	_, err := ps.client.SendEmail(email)
	return err
}

// NewPostMarkSender creates a new postmarkSender with
// its own Postmark client inside
func NewPostMarkSender(serverToken, accountToken string, trackOpens bool) *PostmarkSender {
	return &PostmarkSender{
		client:     postmark.NewClient(serverToken, accountToken),
		trackOpens: trackOpens,
	}
}
