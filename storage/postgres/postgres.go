package postgres

import (
	"context"
	"fmt"
	"strings"
	"test/config"
	"test/pkg/logger"
	"test/storage"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	_ "github.com/lib/pq"
)

type Store struct {
	pool *pgxpool.Pool
	log  logger.ILogger
	cfg  config.Config
}

func New(ctx context.Context, cfg config.Config, log logger.ILogger) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Error("error while parsing config", logger.Error(err))
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Error("error while connecting to db", logger.Error(err))
		return nil, err
	}

	//migration
	m, err := migrate.New("file://migrations/postgres/", url)
	if err != nil {
		log.Error("error while migrating", logger.Error(err))
		return nil, err
	}

	log.Info("???? came")

	if err = m.Up(); err != nil {
		log.Warning("migration up", logger.Error(err))
		if !strings.Contains(err.Error(), "no change") {
			fmt.Println("entered")
			version, dirty, err := m.Version()
			log.Info("version and dirty", logger.Any("version", version), logger.Any("dirty", dirty))
			if err != nil {
				log.Error("err in checking version and dirty", logger.Error(err))
				return nil, err
			}

			if dirty {
				version--
				if err = m.Force(int(version)); err != nil {
					log.Error("ERR in making force", logger.Error(err))
					return nil, err
				}
			}
			log.Warning("WARNING in migrating", logger.Error(err))
			return nil, err
		}
	}

	log.Info("!!!!! came here")

	return Store{
		pool: pool,
		log:  log,
		cfg:  cfg,
	}, nil
}

func (s Store) Close() {
	s.pool.Close()
}

//user1
func (s Store) User() storage.IUserStorage {
	return NewUserRepo(s.pool, s.log)
}

//stocks2
func (s Store) Stocks() storage.IStocksStorage {
	return NewStocksRepo(s.pool, s.log)
}

//marketupdates3
func (s Store) MarketUpdates() storage.IMarketUpdatesStorage {
	return NewMarketUpdatesRepo(s.pool, s.log)
}

//Orders4
func (s Store) Orders() storage.IOrdersStorage {
	return NewOrdersRepo(s.pool, s.log)
}

//orderbooks5
func (s Store) OrderBook() storage.IOrderBookStorage {
	return NewOrderBookRepo(s.pool, s.log)
}

//porfolios6
func (s Store) Porfolios() storage.IPortfoliosStorage {
	return NewPorfoliosRepo(s.pool, s.log)
}

//tradeconfimations7
func (s Store) TradeConfirmations() storage.ITradeConfirmationsStorage{
	return NewTradeConfirmationsRepo(s.pool, s.log)
}

//Marketnews8
func (s Store) MarketNews() storage.IMarketNewsStorage{
	return NewMarketNewsRepo(s.pool, s.log)
}

//AccountInforamtion9
func (s Store) AccountInforamtion() storage.IAccountInforamtionStorage{
	return NewAccountInforamtionRepo(s.pool, s.log)
}

//FundTransfers10
func (s Store) FundTransfers() storage.IFundTransfersStorage{
	return NewFundTransfersRepo(s.pool, s.log)
}