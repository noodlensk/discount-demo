package adapters

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/noodlensk/discount-demo/internal/discount/domain/code"
)

func NewDiscountCodesInMemoryRepository() DiscountCodesInMemoryRepository {
	return DiscountCodesInMemoryRepository{codes: map[string][]*code.DiscountCode{}}
}

type DiscountCodesInMemoryRepository struct {
	codes map[string][]*code.DiscountCode

	sync.RWMutex
}

func (r *DiscountCodesInMemoryRepository) AddDiscountCode(ctx context.Context, d *code.DiscountCode) error {
	r.Lock()
	defer r.Unlock()

	r.codes[d.BrandID()] = append(r.codes[d.BrandID()], d)

	return nil
}

func (r *DiscountCodesInMemoryRepository) GetFirstAvailableDiscountCode(ctx context.Context, brandID string) (*code.DiscountCode, error) {
	r.RLock()
	defer r.RUnlock()

	for _, discountCode := range r.codes[brandID] {
		if discountCode.AssignedUserID() == "" {
			return discountCode, nil
		}
	}

	return nil, code.ErrNoCodesAvailable
}

func (r *DiscountCodesInMemoryRepository) UpdateDiscountCode(ctx context.Context, brandID, id string, updateFn func(ctx context.Context, d *code.DiscountCode) (*code.DiscountCode, error)) error {
	r.RLock()
	defer r.RUnlock()

	for i, discountCode := range r.codes[brandID] {
		if discountCode.Code() != id {
			continue
		}

		newCode, err := updateFn(ctx, discountCode)
		if err != nil {
			return err
		}

		r.codes[brandID][i] = newCode

		return nil
	}

	return errors.New("no code found")
}
