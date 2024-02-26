package cases

import (
	"context"
	"restapi/internal/entities"
)

type CoinStorage interface {
	GetCoins(context.Context) ([]entities.Coin, error)
	GetCoinByName(context.Context, string) (entities.Coin, error)
	StoreCoin(context.Context, entities.Coin) error
}

//mocks for interface
