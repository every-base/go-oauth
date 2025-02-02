package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/every-base/go-oauth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	apiVersion = "2022-11-28"

	apiUser = "https://api.github.com/user"
)

func NewOAuth(config *oauth2.Config) *OAuth {
	config.Endpoint = github.Endpoint
	return &OAuth{config}
}

var _ oauth.OAuth = &OAuth{}

type OAuth struct {
	*oauth2.Config
}

func (c *OAuth) Claims(ctx context.Context, token *oauth2.Token) (oauth.Claims, error) {
	req, err := http.NewRequest(http.MethodGet, apiUser, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", apiVersion)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var claims Claims
	err = json.NewDecoder(res.Body).Decode(&claims)
	return claims, err
}

type Claims struct {
	ID int64 `json:"id"`
}

func (c Claims) UID() string {
	return fmt.Sprintf("%v", c.ID)
}
