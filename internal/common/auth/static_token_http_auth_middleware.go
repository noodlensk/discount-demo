package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/noodlensk/discount-demo/internal/common/server/httperr"
)

func NewStaticTokenHTTPMiddleware() StaticTokenHTTPMiddleware {
	return StaticTokenHTTPMiddleware{tokens: map[string]User{
		"user": {
			ID:    "user",
			Email: "user@example.com",
			Role:  "user",
		},
		"brandA": {
			ID:              "brandA",
			Email:           "brandA@example.com",
			Role:            "brand",
			AllowedBrandIDs: []string{"brandA"},
		},
		"brandB": {
			ID:              "brandB",
			Email:           "brandB@example.com",
			Role:            "brand",
			AllowedBrandIDs: []string{"brandB"},
		},
		"admin": {
			ID:              "admin",
			Email:           "admin@example.com",
			Role:            "brand",
			AllowedBrandIDs: []string{"brandA", "brandB"},
		},
	}}
}

type StaticTokenHTTPMiddleware struct {
	tokens map[string]User
}

func (a StaticTokenHTTPMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			httperr.Unauthorized("empty-bearer-token", nil, w, r)

			return
		}

		user, ok := a.tokens[bearerToken]
		if !ok {
			httperr.Unauthorized("unable-to-verify-jwt", errors.New("unknown token"), w, r)

			return
		}
		// it's always a good idea to use custom type as context value (in this case ctxKey)
		// because nobody from the outside of the package will be able to override/read this value
		ctx = context.WithValue(ctx, userContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (a StaticTokenHTTPMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.EqualFold(headerValue[0:6], "bearer") {
		return headerValue[7:]
	}

	return ""
}
