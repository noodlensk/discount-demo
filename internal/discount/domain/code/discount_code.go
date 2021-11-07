package code

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type DiscountCode struct {
	code           string
	brandID        string
	assignedUserID string
	createdAt      time.Time
	updatedAt      time.Time
}

func (d *DiscountCode) Code() string {
	return d.code
}

func (d *DiscountCode) BrandID() string {
	return d.brandID
}

func (d *DiscountCode) AssignedUserID() string {
	return d.assignedUserID
}

func NewDiscountCode(brandID string) (*DiscountCode, error) {
	if brandID == "" {
		return nil, errors.New("empty brand id")
	}

	return &DiscountCode{
		brandID:   brandID,
		code:      generateCodeID(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// let's assume that it will uniq in 100%
func generateCodeID() string {
	return uuid.New().String()
}
