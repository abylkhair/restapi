package cases

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"restapi/internal/entities"
)

type Service struct {
	storage Coins
	client  CryptoClient
	logger  *logrus.Logger
}

// мы получаем абсолютно все значения из базы и обновляем монеты
func (s *Service) GetActualValue(ctx context.Context) ([]entities.Coin, error) {
	var rates []entities.Coin
	coins, err := s.storage.GetCoins(ctx)
	if err != nil {
		s.logger.Errorf("Failed received data %v", err)
		return nil, errors.Wrapf(entities.ErrInternal, "failed receiving coins from storage : %v", err)
	}

	if coins != nil {
		s.logger.Warningf("Failed to find data %v", coins) // level of err
		return nil, errors.Wrapf(entities.ErrNotFound, "expected got nil result : %v", coins)
	}
	rates = append(rates, coins...)
	return rates, nil
}
func (s *Service) GetCoin(ctx context.Context, coin string) ([]string, error) {

}

func (s *Service) GetValuesFromSrc(ctx context.Context, newTitle string) error {
	titles, err := s.storage.GetTitles(ctx) //можно вынести
	if err != nil {
		s.logger.Warningf("Failed get titles from storage  %v", err)
		return errors.Wrapf(entities.ErrNotFound, "expected not nil result : %v", err)
	}
	if titles == nil {
		s.logger.Warningf("Failed to get from storage title %v", titles)
		return errors.Wrapf(entities.ErrNotFound, "expected not nil result : %v", titles)
	}
	if newTitle != "" {
		_, err := s.client.GetActualValue(ctx, []string{newTitle})
		if err != nil {
			s.logger.Warningf("Failed to get actual value  %v", err)
			return errors.Wrapf(entities.ErrNotFound, "expected not nil result : %v", err)
		}
		titles = append(titles, newTitle)
	}

	return nil
}

//GetACtualRates->Storage(last value of coin string)
//GetActualRates->Storage(string coin)->CryptoClient(coin (if coin == ""))
//StoreActualRates->CryptoClient()->Storage(save)

// two adapters needed
//work with storage
//work with cryptoapi
