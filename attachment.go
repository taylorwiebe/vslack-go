package vslack

// AttachmentInterface is the interface for an attachment
//
//go:generate mockery --inpackage --name=AttachmentInterface
type AttachmentInterface interface {
	SetText(t string) Attachment
	SetColor(c string) Attachment
	SetTitle(t string) Attachment
	SetTitleLink(l string) Attachment
	SetFields(f ...Field) Attachment
	SetMarkdown(m bool) Attachment
	SetMarkdownIn(opts ...MarkdownOption) Attachment
}

// Attachment is a slack attachment
type Attachment struct {
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	Fallback   string   `json:"fallback"`
	Markdown   bool     `json:"mrkdwn"`
	MarkdownIn []string `json:"mrkdwn_in,omitempty"`
	Color      string   `json:"color"`
	Fields     []Field  `json:"fields"`
}

// Field is an a slack attachment
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// SetText is an attachments text
func (a Attachment) SetText(t string) Attachment {
	a.Text = t
	return a
}

// SetColor sets the message colour
func (a Attachment) SetColor(c string) Attachment {
	a.Color = c
	return a
}

// SetTitle of the message
func (a Attachment) SetTitle(t string) Attachment {
	a.Title = t
	return a
}

// SetTitleLink of the message
func (a Attachment) SetTitleLink(l string) Attachment {
	a.TitleLink = l
	return a
}

// SetFields is the list of attachment fields
func (a Attachment) SetFields(f ...Field) Attachment {
	a.Fields = append(a.Fields, f...)
	return a
}

// SetMarkdown sets whether or not to use markdown in a message
func (a Attachment) SetMarkdown(m bool) Attachment {
	a.Markdown = m
	return a
}

// MarkdownConfiguration are options for message markdown
type MarkdownConfiguration struct {
	options map[string]struct{}
}

// MarkdownOption is an option for markdown in a slack message
type MarkdownOption func(m *MarkdownConfiguration)

// Fields set the fields option
func Fields() MarkdownOption {
	return func(m *MarkdownConfiguration) {
		m.options["fields"] = struct{}{}
	}
}

// Text set the text option
func Text() MarkdownOption {
	return func(m *MarkdownConfiguration) {
		m.options["text"] = struct{}{}
	}
}

// Pretext set the pretext option
func Pretext() MarkdownOption {
	return func(m *MarkdownConfiguration) {
		m.options["pretext"] = struct{}{}
	}
}

// SetMarkdownIn takes in supported markdown types
func (a Attachment) SetMarkdownIn(opts ...MarkdownOption) Attachment {
	m := &MarkdownConfiguration{
		options: map[string]struct{}{},
	}

	for _, o := range opts {
		o(m)
	}

	for k := range m.options {
		a.MarkdownIn = append(a.MarkdownIn, k)
	}

	return a
}
