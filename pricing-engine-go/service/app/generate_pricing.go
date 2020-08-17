package app

import (
	"context"
	"errors"

	"pricingengine"
)

type App struct{}

// GeneratePricing will calculate how much a 'risk' be priced or if they should
// be denied.
func (a *App) GeneratePricing(ctx context.Context, input *pricingengine.GeneratePricingRequest) (*pricingengine.GeneratePricingResponse, error) {
	return nil, errors.New("not implemented")
}
