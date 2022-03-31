package service

import (
	"context"
	"github.com/Nerzal/gocloak/v11"
	"net/http"
)

// authentication middleware
type KeycloakAuthMiddleware struct {
	keycloak gocloak.GoCloak
}

func NewKeycloakAuthMiddleware(keycloak gocloak.GoCloak) *KeycloakAuthMiddleware {
	return &KeycloakAuthMiddleware{
		keycloak: keycloak,
	}
}

func (m *KeycloakAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authToken, err := extractBearToken(w, r)
	if err != nil {
		return
	}

	ctx := context.Background()
	v, err := m.keycloak.RetrospectToken(ctx, authToken)
	if err != nil {
		handleError(w, http.StatusUnauthorized, "invalid authorization token")
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
