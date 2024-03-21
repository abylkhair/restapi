package cases

import (
	"context"
	"restapi/internal/entities"
)

type CryptoClient interface {
	GetActualValue(ctx context.Context, title []string) (map[string]entities.Coin, error)
}
