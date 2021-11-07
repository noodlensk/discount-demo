package command

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	commonerrors "github.com/noodlensk/discount-demo/internal/common/errors"
	"github.com/noodlensk/discount-demo/internal/discount/domain/code"
)

const MaxAvailableCodes = 200

type GenerateDiscountCodes struct {
	BrandID       string
	DesiredAmount int
}

type GenerateDiscountCodesHandler struct {
	repo         code.Repository
	brandService BrandService
}

func NewGenerateDiscountCodesHandler(repo code.Repository, brandService BrandService) GenerateDiscountCodesHandler {
	if repo == nil {
		panic("nil repo")
	}

	if brandService == nil {
		panic("nil brandService")
	}

	return GenerateDiscountCodesHandler{repo: repo, brandService: brandService}
}

func (h GenerateDiscountCodesHandler) Handle(ctx context.Context, cmd GenerateDiscountCodes) (err error) {
	if cmd.DesiredAmount < 1 {
		return commonerrors.NewIncorrectInputError("desired amount should be >= 1", "incorrect-desired-amount")
	}

	if cmd.DesiredAmount > MaxAvailableCodes {
		return commonerrors.NewIncorrectInputError(fmt.Sprintf("desired amount should be <= %d", MaxAvailableCodes), "incorrect-desired-amount")
	}

	if cmd.BrandID == "" {
		return commonerrors.NewIncorrectInputError("brand can't be empty", "incorrect-brand")
	}

	brandExist, err := h.brandService.BrandExist(ctx, cmd.BrandID)
	if err != nil {
		return errors.Wrapf(err, "check brand existence")
	}

	if !brandExist {
		return commonerrors.NewIncorrectInputError("brand does not exist", "brand-does-not-exist")
	}

	// TODO: make it inside transaction?
	for i := 0; i < cmd.DesiredAmount; i++ {
		dc, err := code.NewDiscountCode(cmd.BrandID)
		if err != nil {
			return errors.Wrapf(err, "generate %d/%d discount code", i+1, cmd.DesiredAmount)
		}

		if err := h.repo.AddDiscountCode(ctx, dc); err != nil {
			return errors.Wrapf(err, "save %d/%d discount code", i+1, cmd.DesiredAmount)
		}
	}

	return nil
}
