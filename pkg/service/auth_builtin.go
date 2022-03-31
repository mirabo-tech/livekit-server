package service

import (
	"context"
	"github.com/livekit/protocol/auth"
	"net/http"
)

// authentication middleware
type APIKeyAuthMiddleware struct {
	provider auth.KeyProvider
}

func NewAPIKeyAuthMiddleware(provider auth.KeyProvider) *APIKeyAuthMiddleware {
	return &APIKeyAuthMiddleware{
		provider: provider,
	}
}

func (m *APIKeyAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authToken, err := extractBearToken(w, r)
	if err != nil {
		return
	}
	v, err := auth.ParseAPIToken(authToken)
	if err != nil {
		handleError(w, http.StatusUnauthorized, "invalid authorization token")
		return
	}

	secret := m.provider.GetSecret(v.APIKey())
	if secret == "" {
		handleError(w, http.StatusUnauthorized, "invalid API key")
		return
	}

	grants, err := v.Verify(secret)
	if err != nil {
		handleError(w, http.StatusUnauthorized, "invalid token: "+authToken+", error: "+err.Error())
		return
	}

	// set grants in context
	ctx := r.Context()
	r = r.WithContext(context.WithValue(ctx, grantsKey, grants))

	next.ServeHTTP(w, r)
}
