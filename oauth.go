package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type OAuth interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Claims(ctx context.Context, token *oauth2.Token) (Claims, error)
}
