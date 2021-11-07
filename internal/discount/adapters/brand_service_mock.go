package adapters

import (
	"context"

	"go.uber.org/zap"
)

func NewBrandServiceMock(logger *zap.SugaredLogger) BrandServiceMock {
	return BrandServiceMock{
		logger: logger,
		existingBrands: map[string]bool{
			"brandA": true,
			"brandB": true,
		},
	}
}

// BrandServiceMock is a mock representation of brand service. It is needed for checking brand existence in other APIs
type BrandServiceMock struct {
	logger         *zap.SugaredLogger
	existingBrands map[string]bool
}

// UserFetchCode could be used to put message into a queue (AMQP, Kafka, etc) for later processing
func (b BrandServiceMock) UserFetchCode(ctx context.Context, brandID, userID, codeID string) {
	b.logger.Infof("User %q successfully fetch code %q for brand %q", userID, codeID, brandID)
}

func (b BrandServiceMock) BrandExist(ctx context.Context, brandID string) (bool, error) {
	return b.existingBrands[brandID], nil
}
