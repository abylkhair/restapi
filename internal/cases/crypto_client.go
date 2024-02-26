package cases

import (
	"context"
	"restapi/internal/entities"
)

type CryptoClient interface {
	GetActualValue(context.Context) (entities.Coin, error)
	GetActualValueByName(context.Context, string) (string, error)
}
