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

// here we will have basic logic that works with other services storage and other stuff
// мы получаем абсолютно все значения из базы и обновляем монеты
// нужно получить все монеты
// нужно получить монеты по названию
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

// every 5 min send data
// max min average sending in one query
//adapter for telegram config it there
//two scenarios with work in telegram
//cryptocompare.com  //
// min max in service not outer
//
// /start-auto {minutes_count} (пример /start-auto 10, что значит отправлять каждые 10 минут)
//cron
