package cases

import (
	"context"
	"restapi/internal/entities"
)

type Coins interface {
	GetCoins(ctx context.Context) ([]entities.Coin, error)
	GetCoinByName(ctx context.Context, coin string) (entities.Coin, error)
	StoreCoin(ctx context.Context, coins entities.Coin) error
	GetTitles(ctx context.Context) ([]string, error)
}

//mocks for interface
