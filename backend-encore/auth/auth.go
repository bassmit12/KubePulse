package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type VerifyRequest struct {
	Token string `json:"token"`
}

type Claims struct {
	Subject string   `json:"sub"`
	Email   string   `json:"email,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	Issuer  string   `json:"iss"`
}

type VerifyResponse struct {
	Valid  bool    `json:"valid"`
	Claims *Claims `json:"claims,omitempty"`
}

var (
	jwksMu sync.Mutex
	jwks   keyfunc.Keyfunc
)

func getIssuer() string {
	return strings.TrimSuffix(os.Getenv("KEYCLOAK_ISSUER"), "/")
}

func getJWKS() (keyfunc.Keyfunc, error) {
	jwksMu.Lock()
	defer jwksMu.Unlock()
	if jwks != nil {
		return jwks, nil
	}

	issuer := getIssuer()
	if issuer == "" {
		return nil, errors.New("KEYCLOAK_ISSUER is required")
	}

	url := fmt.Sprintf("%s/protocol/openid-connect/certs", issuer)
	k, err := keyfunc.NewDefaultCtx(context.Background(), []string{url})
	if err != nil {
		return nil, err
	}
	jwks = k
	return jwks, nil
}

func extractRoles(claims jwt.MapClaims) []string {
	roles := []string{}
	realm, ok := claims["realm_access"].(map[string]any)
	if !ok {
		return roles
	}
	rawRoles, ok := realm["roles"].([]any)
	if !ok {
		return roles
	}
	for _, r := range rawRoles {
		if s, ok := r.(string); ok {
			roles = append(roles, s)
		}
	}
	return roles
}

// Verify validates Keycloak JWTs using JWKS and returns basic claims.
//encore:api public method=POST path=/auth/verify
func Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	if req == nil || strings.TrimSpace(req.Token) == "" {
		return &VerifyResponse{Valid: false}, nil
	}

	k, err := getJWKS()
	if err != nil {
		return nil, err
	}

	issuer := getIssuer()
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(req.Token, claims, k.Keyfunc, jwt.WithIssuer(issuer), jwt.WithLeeway(30*time.Second))
	if err != nil {
		return &VerifyResponse{Valid: false}, nil
	}

	sub, _ := claims["sub"].(string)
	email, _ := claims["email"].(string)
	iss, _ := claims["iss"].(string)

	return &VerifyResponse{
		Valid: true,
		Claims: &Claims{
			Subject: sub,
			Email:   email,
			Roles:   extractRoles(claims),
			Issuer:  iss,
		},
	}, nil
}
