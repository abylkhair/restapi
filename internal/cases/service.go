package cases

import "context"

type CurrencyUseCase interface {
	UpdateCurrencyRates(ctx context.Context) error
}
