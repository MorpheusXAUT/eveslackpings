package models

type SlackPayload struct {
	Text        string            `json:"text,omitempty"`
	Channel     string            `json:"channel,omitempty"`
	LinkNames   int64             `json:"link_names,omitempty"`
	ParseMode   string            `json:"parse,omitempty"`
	Attachments []SlackAttachment `json:"attachments,omitempty"`
}

type SlackAttachment struct {
	Fallback   string                 `json:"fallback"`
	Text       string                 `json:"text"`
	PreText    string                 `json:"pretext,omitempty"`
	Color      string                 `json:"color,omitempty"`
	MarkdownIn []string               `json:"mrkdwn_in,omitempty"`
	Fields     []SlackAttachmentField `json:"fields,omitempty"`
}

type SlackAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
