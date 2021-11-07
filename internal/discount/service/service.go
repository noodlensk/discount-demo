package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/noodlensk/discount-demo/internal/discount/adapters"
	"github.com/noodlensk/discount-demo/internal/discount/app"
	"github.com/noodlensk/discount-demo/internal/discount/app/command"
)

func NewApplication(ctx context.Context, logger *zap.SugaredLogger) app.Application {
	return newApplication(ctx, logger)
}

func newApplication(ctx context.Context, logger *zap.SugaredLogger) app.Application { //nolint:unparam
	discountCodesRepository := adapters.NewDiscountCodesInMemoryRepository()
	brandService := adapters.NewBrandServiceMock(logger)

	return app.Application{
		Commands: app.Commands{
			GenerateDiscountCodes: command.NewGenerateDiscountCodesHandler(&discountCodesRepository, brandService),
			FetchDiscount:         command.NewFetchDiscountCodesHandler(&discountCodesRepository, brandService),
		},
	}
}
