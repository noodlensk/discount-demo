package auth

import (
	"context"

	commonerrors "github.com/noodlensk/discount-demo/internal/common/errors"
)

type User struct {
	ID              string
	Email           string
	Role            string
	AllowedBrandIDs []string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

// NoUserInContextError if we expect that the user of the function may be interested with concrete error,
// it's a good idea to provide variable with this error
var NoUserInContextError = commonerrors.NewAuthorizationError("no user in context", "no-user-found")

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}
