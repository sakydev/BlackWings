package apps

import (
	"black-wings/internal/types"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/guregu/null/v5"
	"github.com/samber/do"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const User = "me"

func InjectGmailService(i *do.Injector) (*GmailService, error) {
	return &GmailService{}, nil
}

type GmailService struct{}

func (s GmailService) Search(ctx context.Context, client *http.Client, options types.SearchFlags) ([]types.EmailResponse, error) {
	var results []types.EmailResponse

	srv, err := initialize(ctx, client)
	if err != nil {
		return results, fmt.Errorf("error initializing Gmail service: %v", err)
	}

	query := s.buildQuery(srv, options)
	messages, err := query.Do()
	if err != nil {
		return results, fmt.Errorf("error retrieving messages: %v", err)
	}

	if len(messages.Messages) == 0 {
		return results, fmt.Errorf("no messages found")
	}

	for _, message := range messages.Messages {
		messageDetails, err := s.getMessageDetails(srv, User, message.Id)
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

func (s GmailService) getMessageDetails(srv *gmail.Service, user string, messageID string) (types.EmailResponse, error) {
	message, err := srv.Users.Messages.Get(user, messageID).Do()
	if err != nil {
		return types.EmailResponse{}, fmt.Errorf("error retrieving message %s: %v", messageID, err)
	}

	msg, err := srv.Users.Messages.Get(user, message.Id).Do()
	if err != nil {
		return types.EmailResponse{}, fmt.Errorf("error retrieving message %s: %v", message.Id, err)
	}

	var senderName, senderEmail, subject string
	for _, header := range msg.Payload.Headers {
		if header.Name == "Subject" {
			subject = header.Value
		} else if header.Name == "From" {
			senderName, senderEmail = s.extractSender(header.Value)
		}
	}

	readableTime, err := s.extractTime(msg.Payload.Headers[1].Value)
	if err != nil {
		return types.EmailResponse{}, fmt.Errorf("error parsing date: %v", err)
	}

	return types.EmailResponse{
		Subject:     subject,
		SenderName:  null.StringFrom(senderName),
		SenderEmail: senderEmail,
		Date:        readableTime,
		Snippet:     msg.Snippet,
	}, nil
}

func (s GmailService) extractSender(fromHeader string) (name, email string) {
	start := strings.Index(fromHeader, "<")
	end := strings.Index(fromHeader, ">")

	if start != -1 && end != -1 {
		name = strings.TrimSpace(fromHeader[:start])
		email = strings.TrimSpace(fromHeader[start+1 : end])

		return name, email
	}

	return name, fromHeader
}

func (s GmailService) extractTime(timestamp string) (time.Time, error) {
	timeLayout := "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"

	parts := strings.Split(timestamp, ";")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("invalid timestamp format")
	}

	date := strings.TrimSpace(parts[1])

	return time.Parse(timeLayout, date)
}
