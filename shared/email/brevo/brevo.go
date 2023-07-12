package brevo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/campbel/terpcatalog/shared/types"
	"github.com/campbel/terpcatalog/util/config"
	"github.com/campbel/terpcatalog/util/log"
)

type Client struct {
	APIKey string
	http   *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		http:   &http.Client{},
	}
}

// curl -H 'api-key:YOUR_API_V3_KEY'
// -X POST -d '{
// # Define the campaign settings
// "name":"Campaign sent via the API",
// "subject":"My subject",
// "sender": { "name": "From name", "email":"myfromemail@mycompany.com" },
// "type": "classic",
// # Content that will be sent
// "htmlContent": "Congratulations! You successfully sent this example campaign via the Brevo API.",
// # Select the recipients
// "recipients": { "listIds": [2,7] },
// # Schedule the sending in one hour
// "scheduledAt": "2018-01-01 00:00:01",
// }'

type Email struct {
	Sender      types.EmailAddress   `json:"sender"`
	To          []types.EmailAddress `json:"to"`
	Cc          []types.EmailAddress `json:"cc"`
	Bcc         []types.EmailAddress `json:"bcc"`
	ReplyTo     types.EmailAddress   `json:"replyTo"`
	Subject     string               `json:"subject"`
	HtmlContent string               `json:"htmlContent"`
	Headers     map[string]string    `json:"headers"`
	Params      map[string]string    `json:"params"`
}

func (c *Client) SendEmail(to []types.EmailAddress, subject, content string) error {
	data, err := json.Marshal(Email{
		Sender: types.EmailAddress{
			Name:  "TerpScout",
			Email: config.EmailSender(),
		},
		ReplyTo: types.EmailAddress{
			Name:  "TerpScout",
			Email: config.EmailSender(),
		},
		To:          to,
		Subject:     subject,
		HtmlContent: content,
	})
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Add("api-key", c.APIKey)

	response, err := c.http.Do(request)
	if err != nil {
		log.Error("brevo do error", err)
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error("brevo read body error", err)
		return err
	}

	if response.StatusCode != http.StatusCreated {
		err := errors.New("brevo bad status: " + response.Status)
		log.Error("brevo status error", err, "body", string(body))
		return err
	}

	log.Debug("brevo api response", "to", to, "subject", subject, "body", string(body))

	return nil
}
