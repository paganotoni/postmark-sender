### Postmark Buffalo Sender

This is a buffalo sender for the [postmark](https://postmarkapp.com/) transactional email service.

#### How to use

In your `mailers.go`

```
...

var sender mail.Sender
var hcomposer hermes.Hermes

func init() {
	serverToken := envy.Get("POSTMARK_SERVER_TOKEN", "")
	accountToken := envy.Get("POSTMARK_ACCOUNT_TOKEN", "")
	sender = psender.NewPostMarkSender(serverToken, accountToken, false)
}
```

And then in your mailers you would do the same `sender.Send(m)` as this sender matches the mail.