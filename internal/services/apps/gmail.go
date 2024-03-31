package apps

import (
	"BlackWings/internal/types"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/samber/do"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailMessageResponse struct {
	Subject string
	Sender  string
	Date    time.Time
	Snippet string
}

func InjectGmailService(i *do.Injector) (*GmailService, error) {
	return &GmailService{}, nil
}

type GmailService struct{}

func (s GmailService) Search(ctx context.Context, options types.SearchFlags, client *http.Client) ([]GmailMessageResponse, error) {
	var results []GmailMessageResponse

	srv, err := initialize(ctx, client)
	if err != nil {
		return results, fmt.Errorf("error initializing Gmail service: %v", err)
	}

	user := "me"
	query := s.buildQuery(srv, options)
	messages, err := query.Do()
	if err != nil {
		return results, fmt.Errorf("error retrieving messages: %v", err)
	}

	if len(messages.Messages) == 0 {
		return results, fmt.Errorf("no messages found")
	}

	for _, message := range messages.Messages {
		messageDetails, err := s.getMessageDetails(srv, user, message.Id)
		if err != nil {
			return results, err
		}

		results = append(results, messageDetails)
	}

	return results, nil
}

func initialize(ctx context.Context, googleService *http.Client) (*gmail.Service, error) {
	srv, err := gmail.NewService(ctx, option.WithHTTPClient(googleService))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmail client: %v", err)
	}
	return srv, nil
}

func (s GmailService) buildQuery(srv *gmail.Service, options types.SearchFlags) *gmail.UsersMessagesListCall {
	user := "me"
	query := srv.Users.Messages.List(user).Q(options.Query)

	if options.Limit > 0 {
		query.MaxResults(options.Limit)
	}

	return query
}

func (s GmailService) getMessageDetails(srv *gmail.Service, user string, messageID string) (GmailMessageResponse, error) {
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
