package main

import (
	"BlackWings/internal/services/apps"
	"BlackWings/internal/services/integrations"
	"log"
)

func main() {
	// Initialize Gmail service
	googleService, err := integrations.InitializeGoogleService()
	if err != nil {
		log.Fatalf("Unable to initialize Gmail service: %v", err)
	}

	gmailService, err := apps.InitializeGmailService(googleService)

	// Search for "hello" in Gmail
	query := "hello"
	apps.SearchGmail(query, gmailService)
}
