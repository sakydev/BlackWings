package apps

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/gmail/v1"
)

func InitializeGmailService(googleService *http.Client) (*gmail.Service, error) {
	// Create Gmail service
	srv, err := gmail.New(googleService)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Gmail client: %v", err)
	}
	return srv, nil
}

func SearchGmail(query string, srv *gmail.Service) {
	user := "me"
	messages, err := srv.Users.Messages.List(user).Q(query).Do()
	if err != nil {
		log.Fatalf("Error searching Gmail: %v", err)
	}

	if len(messages.Messages) == 0 {
		fmt.Println("No messages found.")
		return
	}

	fmt.Println("Messages:")
	for _, message := range messages.Messages {
		fmt.Printf("- Email ID: %v\n", message.Id)
	}
}
