# VSlack - go
### Version 1.0.0
Send slack messages to a slack incoming web hook in go

## Examples
### synchronous
```
func synchronous() {
	s := vslack.NewVSlack(incomingWebHook)
	err := s.SetChannel("#random").
		SetIconEmoji(":laughing:").
		SetUsername("VSlack").
		SetMessage("Test message").
		Send()
	if err != nil {
		panic(err)
	}
}
```

### Concurrent Example
```
func concurrent(c chan error) {
	s := vslack.NewVSlack(incomingWebHook)
	err := s.SetChannel("#random").
		SetIconEmoji(":laughing:").
		SetUsername("VSlack").
		SetMessage("Test message").
		Send()

	c <- err
}

func main() {
	c := make(chan error)
	go concurrent(c)
	if err := <-c; err != nil {
		panic(err)
	}
}
```


### Attachments 
```
func attach() {
	s := vslack.NewVSlack(incomingWebHook)
	err := s.SetChannel("#general").
		SetIconEmoji(":laughing:").
		SetUsername("VSlack").
		SetAttachments(
			vslack.NewVSlackAttachment().
				SetText("*test*").
				SetTitle("title").
				SetMarkdownIn(
					vslack.Text(),
					vslack.Fields(),
					vslack.Pretext())).
		Send()

	if err != nil {
		panic(err)
	}
}
```
