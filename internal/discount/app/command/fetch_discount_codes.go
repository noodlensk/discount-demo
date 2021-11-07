package command

import (
	"context"

	"github.com/pkg/errors"

	commonerrors "github.com/noodlensk/discount-demo/internal/common/errors"
	"github.com/noodlensk/discount-demo/internal/discount/domain/code"
)

type FetchDiscountCode struct {
	BrandID string
	UserID  string
}

type FetchDiscountCodeHandler struct {
	repo         code.Repository
	brandService BrandService
}

func NewFetchDiscountCodesHandler(repo code.Repository, brandService BrandService) FetchDiscountCodeHandler {
	if repo == nil {
		panic("nil repo")
	}

	if brandService == nil {
		panic("nil brandService")
	}

	return FetchDiscountCodeHandler{repo: repo, brandService: brandService}
}

func (h FetchDiscountCodeHandler) Handle(ctx context.Context, cmd FetchDiscountCode) (*code.DiscountCode, error) {
	if cmd.BrandID == "" {
		return nil, commonerrors.NewIncorrectInputError("brand can't be empty", "incorrect-brand")
	}

	if cmd.UserID == "" {
		return nil, commonerrors.NewIncorrectInputError("user id can't be empty", "incorrect-user")
	}

	brandExist, err := h.brandService.BrandExist(ctx, cmd.BrandID)
	if err != nil {
		return nil, errors.Wrapf(err, "check brand existence")
	}

	if !brandExist {
		return nil, commonerrors.NewIncorrectInputError("brand does not exist", "brand-does-not-exist")
	}

	dc, err := h.repo.GetFirstAvailableDiscountCode(ctx, cmd.BrandID)
	if err != nil {
		if errors.Cause(err) == code.ErrNoCodesAvailable {
			return nil, commonerrors.NewIncorrectInputError("no codes available for selected brand", "no-codes-available")
		}

		return nil, err
	}

	var discountCode *code.DiscountCode

	err = h.repo.UpdateDiscountCode(ctx, dc.BrandID(), dc.Code(), func(ctx context.Context, d *code.DiscountCode) (*code.DiscountCode, error) {
		if err := d.AssignToUser(cmd.UserID); err != nil {
			return nil, err
		}

		discountCode = d

		return d, nil
	})

	// inform brand that user fetch a code
	defer func() {
		h.brandService.UserFetchCode(ctx, discountCode.BrandID(), discountCode.AssignedUserID(), discountCode.Code())
	}()

	return discountCode, err
}
