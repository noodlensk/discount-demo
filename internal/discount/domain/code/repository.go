package code

import (
	"context"

	"github.com/pkg/errors"
)

var ErrNoCodesAvailable = errors.New("no codes available")

type Repository interface {
	AddDiscountCode(ctx context.Context, d *DiscountCode) error
	GetFirstAvailableDiscountCode(ctx context.Context, brandID string) (*DiscountCode, error)
	UpdateDiscountCode(
		ctx context.Context,
		brandID string,
		id string,
		updateFn func(ctx context.Context, d *DiscountCode) (*DiscountCode, error),
	) error
}
