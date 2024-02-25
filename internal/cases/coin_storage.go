package cases

import (
	"context"
	"restapi/internal/entities"
)

type CoinStorage interface {
	getCoins(context.Context) ([]entities.Coin, error)
	GetCoinByName(context.Context, string) (entities.Coin, error)
	PostCoin(context.Context, entities.Coin) error
	DeleteCoin(context.Context, entities.Coin) error
}
