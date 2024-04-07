package integrations

import (
	"black-wings/internal/repositories"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/samber/do"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func InjectGoogleService(i *do.Injector) (*GoogleService, error) {
	return &GoogleService{
		accountRepo: do.MustInvoke[repositories.AccountRepository](i),
	}, nil
}

type GoogleService struct {
	accountRepo repositories.AccountRepository
}

func (s *GoogleService) Init(ctx context.Context, database *sql.DB, accountId int64, credentials, savedToken string) (*http.Client, error) {
	config, err := google.ConfigFromJSON([]byte(credentials), gmail.GmailReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	token, err := jsonToToken(savedToken)
	if err != nil {
		return nil, err
	}

	if s.isTokenExpired(token) {
		token = s.getTokenFromWeb(config)

		tokenJSON, err := json.Marshal(token)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal token response: %v", err)
		}

		err = s.accountRepo.UpdateToken(ctx, database, accountId, string(tokenJSON))
		if err != nil {
			return nil, fmt.Errorf("unable to update token: %v", err)
		}
	}

	// Create OAuth2 HTTP client
	client := config.Client(ctx, token)

	return client, nil
}

func (s *GoogleService) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return token
}

func (s *GoogleService) isTokenExpired(token *oauth2.Token) bool {
	if token.Expiry.IsZero() {
		return false
	}

	return token.Expiry.Before(time.Now())
}

func jsonToToken(tokenJSON string) (*oauth2.Token, error) {
	var token oauth2.Token

	err := json.Unmarshal([]byte(tokenJSON), &token)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal token: %v", err)
	}

	return &token, nil
}
