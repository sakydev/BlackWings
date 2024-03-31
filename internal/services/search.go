package services

import (
	"BlackWings/internal/services/apps"
	"BlackWings/internal/services/integrations"
	"BlackWings/internal/types"
	"fmt"

	"github.com/samber/do"
)

func InjectSearchService(i *do.Injector) (*SearchService, error) {
	return &SearchService{
		gmailService: do.MustInvoke[*apps.GmailService](i),
	}, nil
}

type SearchService struct {
	gmailService *apps.GmailService
}

func (s SearchService) Search(options types.SearchFlags) ([]apps.GmailMessageResponse, error) {
	var results []apps.GmailMessageResponse
	googleService, err := integrations.InitializeGoogleService()
	if err != nil {
		return results, fmt.Errorf("error initializing Google service: %v", err)
	}

	results, err = s.gmailService.Search(options, googleService)
	if err != nil {
		return results, fmt.Errorf("error searching Gmail: %v", err)
	}

	return results, nil
}
