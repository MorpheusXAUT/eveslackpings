package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Ping struct {
	ID          int64
	Sender      string
	Message     string
	Destination string
	Timestamp   time.Time
}

func NewPing(sender string, message string, destination string) *Ping {
	ping := &Ping{
		Sender:      sender,
		Message:     message,
		Destination: destination,
		Timestamp:   time.Now(),
	}

	return ping
}

func (ping *Ping) Send(webhookURL string) error {
	payload := SlackPayload{
		Text:      "@channel Turn your head and COF! @group",
		Channel:   ping.Destination,
		ParseMode: "full",
		Attachments: []SlackAttachment{
			SlackAttachment{
				Fallback:   fmt.Sprintf("@channel Turn your head and COF! @group --- %s", ping.Message),
				Color:      "danger",
				MarkdownIn: []string{"pretext", "text", "title", "fields", "fallback"},
				Fields: []SlackAttachmentField{
					SlackAttachmentField{
						Title: "Message:",
						Value: ping.Message,
						Short: false,
					},
					SlackAttachmentField{
						Title: "Sender:",
						Value: ping.Sender,
						Short: false,
					},
					SlackAttachmentField{
						Title: "Timestamp:",
						Value: ping.Timestamp.String(),
						Short: false,
					},
				},
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("Failed to send ping: [%s]", string(response))
	}

	return nil
}
