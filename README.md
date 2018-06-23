### Postmark Buffalo Sender

This is a [buffalo](github.com/gobuffalo/buffalo) sender for the [postmark](https://postmarkapp.com/) transactional email service.

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

And then in your mailers you would do the same `sender.Send(m)` as this sender matches buffalos [`mail.Sender`](https://github.com/gobuffalo/buffalo/blob/master/mail/mail.go#L4) interface.

#### Test mode

Whenever the GO_ENV variable is set to be `test` this sender will use [mocksmtp](https://github.com/stanislas-m/mocksmtp) sender to send messages, you can read values in your tests within the property `TestSender` of the PostmarkSender.