package services

import (
	"BlackWings/internal/services/apps"
	"BlackWings/internal/services/integrations"
	"BlackWings/internal/types"
	"fmt"
)

type SearchService interface {
	Search() string
}

func Search(options types.SearchFlags) ([]apps.GmailMessageResponse, error) {
	var results []apps.GmailMessageResponse
	googleService, err := integrations.InitializeGoogleService()
	if err != nil {
		return results, fmt.Errorf("error initializing Google service: %v", err)
	}

	gmailService, err := apps.InitializeGmailService(googleService)

	return apps.SearchGmail(options, gmailService)
}
