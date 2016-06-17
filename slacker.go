package slacker

// "http"
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// Webhook is the URL for the webhook
type Webhook struct {
	Url *url.URL
}

// Message - foo
type Message struct {
	Username string `json:"username"`
	Icon     string `json:"icon_emoji"`
	Text     string `json:"text"`
}

// NewSlackEndpoint -
func NewSlackEndpoint(slack string) (Webhook, error) {
	u, err := url.Parse(slack)
	if err != nil {
		return Webhook{}, err
	}
	s := Webhook{
		Url: u,
	}
	return s, nil
}

// Send -
func (webhook *Webhook) Send(message Message) (int, error) {
	m, err := json.Marshal(message)
	if err != nil {
		return 0, err
	}
	u := webhook.Url.String()
	buf := bytes.NewReader(m)
	r, err := http.Post(u, "application/json", buf)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()
	return r.StatusCode, nil
}
