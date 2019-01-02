### Postmark Buffalo Sender

This is a [buffalo](github.com/gobuffalo/buffalo) sender for the [postmark](https://postmarkapp.com/) transactional email service.

#### How to use

In your `mailers.go`

```go
import psender "github.com/paganotoni/postmark-sender"
...

var sender mail.Sender
var hcomposer hermes.Hermes

func init() {
	sender = psender.NewPostMarkSender(envy.Get("POSTMARK_SERVER_TOKEN", ""), envy.Get("POSTMARK_ACCOUNT_TOKEN", ""), false)
}
```

And then in your mailers you would do the same `sender.Send(m)` as this sender matches buffalos [`mail.Sender`](https://github.com/gobuffalo/buffalo/blob/master/mail/mail.go#L4) interface.

#### Test mode

Whenever the GO_ENV variable is set to be `test` this sender will use [mocksmtp](https://github.com/stanislas-m/mocksmtp) sender to send messages, you can read values in your tests within the property `TestSender` of the PostmarkSender.