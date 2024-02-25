package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"restapi/internal/entities"
)

type CoinRepository struct {
	db *sqlx.DB
}

func NewCoinRepository() *CoinRepository {
	return &CoinRepository{}
}

func (r *CoinRepository) getCoins(ctx context.Context) ([]entities.Coin, error) {

}

func (r *CoinRepository) GetCoinByName(ctx context.Context, name string) (entities.Coin, error) {

}

func (r *CoinRepository) PostCoin(ctx context.Context, coin entities.Coin) error {

}

func (r *CoinRepository) DeleteCoin(ctx context.Context, coin entities.Coin) error {

}
