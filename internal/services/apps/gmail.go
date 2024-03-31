package apps

import (
	"BlackWings/internal/types"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/api/gmail/v1"
)

type GmailMessageResponse struct {
	Subject string
	Sender  string
	Date    time.Time
	Snippet string
}

func InitializeGmailService(googleService *http.Client) (*gmail.Service, error) {
	// Create Gmail service
	srv, err := gmail.New(googleService)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmail client: %v", err)
	}
	return srv, nil
}

func SearchGmail(options types.SearchFlags, srv *gmail.Service) ([]GmailMessageResponse, error) {
	var results []GmailMessageResponse

	user := "me"
	messages, err := srv.Users.Messages.List(user).Q(options.Query).MaxResults(options.Limit).Do()
	if err != nil {
		return results, fmt.Errorf("error retrieving messages: %v", err)
	}

	if len(messages.Messages) == 0 {
		return results, fmt.Errorf("no messages found")
	}

	for _, message := range messages.Messages {
		messageDetails, err := getMessageDetails(srv, user, message.Id)
		if err != nil {
			return results, err
		}

		results = append(results, messageDetails)
	}

	return results, nil
}

func getMessageDetails(srv *gmail.Service, user string, messageID string) (GmailMessageResponse, error) {
	message, err := srv.Users.Messages.Get(user, messageID).Do()
	if err != nil {
		return GmailMessageResponse{}, fmt.Errorf("error retrieving message %s: %v", messageID, err)
	}

	msg, err := srv.Users.Messages.Get(user, message.Id).Do()
	if err != nil {
		return GmailMessageResponse{}, fmt.Errorf("error retrieving message %s: %v", message.Id, err)
	}

	var subject string
	var sender string
	for _, header := range msg.Payload.Headers {
		if header.Name == "Subject" {
			subject = header.Value
			break
		} else if header.Name == "From" {
			sender = header.Value
		}
	}

	// Extract date
	/*date, err := time.Parse(time.RFC822, msg.Payload.Headers[1].Value)
	if err != nil {
		return GmailMessageResponse{}, fmt.Errorf("error parsing date: %v", err)
	}*/

	return GmailMessageResponse{
		Subject: subject,
		Sender:  sender,
		//Date:    message.Payload.Headers[1].Value,
		Snippet: msg.Snippet,
	}, nil
}
