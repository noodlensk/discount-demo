package command

import "context"

type BrandService interface {
	BrandExist(ctx context.Context, brandID string) (bool, error)
	// UserFetchCode should use fire and forget approach (i.e. non-blocking call)
	UserFetchCode(ctx context.Context, brandID string, userID string, codeID string)
}
