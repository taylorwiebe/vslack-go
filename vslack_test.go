package vslack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVSlack(t *testing.T) {
	t.Run("Test the uri is set", func(t *testing.T) {
		actual := NewVSlack("uri")
		expected := &VSlack{IncomingWebhookURI: "uri"}
		assert.Equal(t, expected, actual)
	})
}

func TestValidate(t *testing.T) {
	tests := []struct {
		Description   string
		In            *VSlack
		ExpectedError error
	}{
		{
			Description: "Should succeed with valid arguments",
			In: &VSlack{
				IncomingWebhookURI: "uri",
				Message:            "A message",
			},
			ExpectedError: nil,
		},
		{
			Description: "Should succeed with attachments",
			In: &VSlack{
				IncomingWebhookURI: "uri",
				Attachments: []Attachment{
					{
						Title:      "title",
						Text:       "text",
						Fallback:   "fallback",
						Markdown:   true,
						MarkdownIn: []string{},
						Color:      "BLACK",
						Fields: []Field{
							{
								Title: "title",
								Value: "value",
							},
						},
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Description: "Should fail with missing webhook",
			In: &VSlack{
				IncomingWebhookURI: "",
				Message:            "A message",
			},
			ExpectedError: errors.New("VSlack needs an incoming web hook"),
		},
		{
			Description: "Should fail with missing message and attachments",
			In: &VSlack{
				IncomingWebhookURI: "uri",
			},
			ExpectedError: errors.New("VSlack needs a message, or attachments"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			err := tt.In.validate()
			assert.Equal(t, tt.ExpectedError, err)
		})
	}
}
