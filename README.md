# VSlack - Go
### Version 1.4.0
Send messages to slack using a slack incoming web hook in Go.

## Import

```
import "github.com/taylorwiebe/vslack-go"
```

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


### @-tagging Users

``` go
func tagging() {
	s := vslack.NewVSlack(incomingWebHook)
	err := s.SetChannel("#random").
		SetIconEmoji(":laughing:").
		SetUsername("VSlack").
		SetMessage("With great power comes great responsibility, @channel").
        SetLinkNames(true).
		Send()
	if err != nil {
		panic(err)
	}
}
```

## Change Log

### 1.4.0
- Add LinkNames parameter, for enabling @-tagging of users without having to know their slack userid.

### 1.3.0
- Add mocks to Interface and AttachmentInterface

### 1.2.0
- Fix name in Interface of `SetIncomingWebhookURI`

### 1.1.0
- Add implementation for SendAsync

### 1.0.0
- Initial Release

## Copyright

Copyright 2017 Vendasta Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
