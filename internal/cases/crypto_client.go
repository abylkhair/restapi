package cases

import (
	"context"
	"restapi/internal/entities"
)

type CryptoClient interface {
	GetActualValue(ctx context.Context, title []string) (map[string]entities.Coin, error)
	GetCurrency(ctx context.Context, currencies []string) (string, []string, error)
	GetCurrencyStatistics(ctx context.Context, currencies []string) (string, []string, error)
}

//func (s *Service) GetActualValueFromAPI(ctx context.Context, newTitle string) error {
//
//}
