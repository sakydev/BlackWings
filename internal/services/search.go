package services

import (
	"black-wings/internal/repositories"
	"black-wings/internal/services/apps"
	"black-wings/internal/services/integrations"
	"black-wings/internal/types"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/samber/do"
)

func InjectSearchService(i *do.Injector) (*SearchService, error) {
	return &SearchService{
		googleService: do.MustInvoke[*integrations.GoogleService](i),
		gmailService:  do.MustInvoke[*apps.GmailService](i),
		accountRepo:   do.MustInvoke[repositories.AccountRepository](i),
	}, nil
}

type ProviderService interface {
	Init(ctx context.Context, database *sql.DB, accountId int64, credentials, token string) (*http.Client, error)
}

type AccountSearchService interface {
	Search(ctx context.Context, client *http.Client, options types.SearchFlags) ([]types.EmailResponse, error)
}

type SearchService struct {
	// Service providers e.g Google
	googleService *integrations.GoogleService

	// Account providers e.g Gmail
	gmailService *apps.GmailService

	// repositories
	accountRepo repositories.AccountRepository
}

func (s SearchService) Search(ctx context.Context, database *sql.DB, options types.SearchFlags) ([]types.EmailResponse, error) {
	var results []types.EmailResponse

	accounts, err := s.accountRepo.List(ctx, database)
	if err != nil {
		return nil, fmt.Errorf("failed to list accounts: %w", err)
	}

	if len(accounts) == 0 {
		return results, fmt.Errorf("no accounts found")
	}

	for _, account := range accounts {
		providerService, err := s.getProviderService(account.App.Provider)
		if err != nil {
			return nil, fmt.Errorf("failed to get provider service: %w", err)
		}

		provider, err := providerService.Init(ctx, database, account.ID, account.CredentialsJSON, account.TokenJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to init provider %s: %w", account.App.Provider, err)
		}

		accountSearchService, err := s.getAccountSearchService(account.App.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to get account search service %s: %w", account.App.Name, err)
		}

		currentResults, err := accountSearchService.Search(ctx, provider, options)
		if err != nil {
			return nil, fmt.Errorf("failed to search account: %w", err)
		}

		results = append(results, currentResults...)
	}

	return results, nil
}

func (s SearchService) getProviderService(name string) (ProviderService, error) {
	name = strings.ToLower(name)

	switch name {
	case "google":
		return s.googleService, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", name)
	}
}

func (s SearchService) getAccountSearchService(name string) (AccountSearchService, error) {
	name = strings.ToLower(name)

	switch name {
	case "gmail":
		return s.gmailService, nil
	default:
		return nil, fmt.Errorf("unsupported account service: %s", name)
	}
}
