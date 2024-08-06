package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type stocksRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewStocksRepo(db *pgxpool.Pool, log logger.ILogger) storage.IStocksStorage {
	return &stocksRepo{
		db:  db,
		log: log,
	}
}

func (s *stocksRepo) Create(ctx context.Context, stock models.CreateStock) (string, error) {
	query := `INSERT INTO stocks(symbol, company_name, current_price) VALUES($1, $2, $3) RETURNING stock_id`
	var stockID int
	err := s.db.QueryRow(ctx, query, stock.Symbol, stock.CompanyName, stock.CurrentPrice).Scan(&stockID)
	if err != nil {
		s.log.Error("error is while inserting stock data", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(stockID), nil
}

func (s *stocksRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Stock, error) {
	var stock models.Stock
	query := `SELECT stock_id, symbol, company_name, current_price, created_at FROM stocks WHERE stock_id = $1`
	err := s.db.QueryRow(ctx, query, key.ID).Scan(&stock.StockID, &stock.Symbol, &stock.CompanyName, &stock.CurrentPrice, &stock.CreatedAt)
	if err != nil {
		s.log.Error("error is while selecting stock", logger.Error(err))
		return models.Stock{}, err
	}
	return stock, nil
}

func (s *stocksRepo) GetList(ctx context.Context, req models.GetListRequest) (models.StockResponse, error) {
	var (
		stocks     = []models.Stock{}
		count      = 0
		query      = `SELECT stock_id, symbol, company_name, current_price, created_at FROM stocks`
		countQuery = `SELECT COUNT(1) FROM stocks`
		filter     = ""
		page       = req.Page
		offset     = (page - 1) * req.Limit
		search     = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE symbol ILIKE '%%%s%%' OR company_name ILIKE '%%%s%%'", search, search)
	}

	countQuery += filter
	err := s.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		s.log.Error("error is while selecting count", logger.Error(err))
		return models.StockResponse{}, err
	}

	query += filter
	query += ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		s.log.Error("error is while selecting stocks", logger.Error(err))
		return models.StockResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		if err := rows.Scan(&stock.StockID, &stock.Symbol, &stock.CompanyName, &stock.CurrentPrice, &stock.CreatedAt); err != nil {
			s.log.Error("error is while scanning data", logger.Error(err))
			return models.StockResponse{}, err
		}
		stocks = append(stocks, stock)
	}

	return models.StockResponse{
		Stocks: stocks,
		Count:  count,
	}, nil
}

func (s *stocksRepo) Update(ctx context.Context, stock models.UpdateStock) (string, error) {
	query := `UPDATE stocks SET symbol = $1, company_name = $2, current_price = $3 WHERE stock_id = $4 RETURNING stock_id`
	var stockID int
	err := s.db.QueryRow(ctx, query, stock.Symbol, stock.CompanyName, stock.CurrentPrice, stock.StockID).Scan(&stockID)
	if err != nil {
		s.log.Error("error is while updating stock", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(stockID), nil
}

func (s *stocksRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM stocks WHERE stock_id = $1`
	_, err := s.db.Exec(ctx, query, key.ID)
	if err != nil {
		s.log.Error("error is while deleting stock", logger.Error(err))
		return err
	}
	return nil
}
