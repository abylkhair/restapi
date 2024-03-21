package adapters

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"restapi/internal/entities"
	"time"
)

// что нужно и что будет использовать
type PGStorage struct {
	pgx    *pgxpool.Pool
	logger *logrus.Logger
}

// конструктор
func NewPGStorage(pgx *pgxpool.Pool, logger *logrus.Logger) (*PGStorage, error) {
	if pgx == nil {
		return nil, errors.Wrapf(entities.ErrInitFail, "expected init db : %v", pgx)
	}
	if logger == nil {
		return nil, errors.Wrapf(entities.ErrInitFail, "expected logger : %v", pgx)
	}
	return &PGStorage{
		pgx: pgx, logger: logger,
	}, nil
}

// GetCoins,GetTitles
func (p *PGStorage) GetCoins(ctx context.Context, coins []*entities.Coin) error {
	q := `SELECT id, title, price, created_at FROM rates;`
	rows, err := p.pgx.Query(ctx, q)
	if err != nil {
		p.logger.Errorf("failed to execute query: %v", err)
		return errors.Wrap(entities.ErrStorageRead, "failed to execute query")
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var title string
		var price float64
		var createdAt time.Time

		if err := rows.Scan(&id, &title, &price, &createdAt); err != nil {
			p.logger.Errorf("failed to scan row: %v", err)
			return errors.Wrap(entities.ErrStorageRead, "failed to scan row")
		}

		crypto := entities.NewCoin(id, title, price, createdAt)
		coins = append(coins, crypto)
	}

	if err := rows.Err(); err != nil {
		p.logger.Errorf("error iterating over rows: %v", err)
		return errors.Wrap(entities.ErrStorageRead, "error iterating over rows")
	}

	return nil
}
