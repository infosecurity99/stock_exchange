package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type portfoliosRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewPorfoliosRepo(db *pgxpool.Pool, log logger.ILogger) storage.IPortfoliosStorage {
	return &portfoliosRepo{
		db:  db,
		log: log,
	}
}

func (p *portfoliosRepo) Create(ctx context.Context, portfolio models.CreatePortfolio) (string, error) {
	query := `INSERT INTO portfolios(user_id, stock_id, quantity) 
	          VALUES($1, $2, $3) RETURNING portfolio_id`
	var portfolioID int
	err := p.db.QueryRow(ctx, query, portfolio.UserID, portfolio.StockID, portfolio.Quantity).Scan(&portfolioID)
	if err != nil {
		p.log.Error("error is while inserting portfolio", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(portfolioID), nil
}

func (p *portfoliosRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Portfolio, error) {
	var portfolio models.Portfolio
	query := `SELECT portfolio_id, user_id, stock_id, quantity, created_at 
	          FROM portfolios WHERE portfolio_id = $1`
	err := p.db.QueryRow(ctx, query, key.ID).Scan(&portfolio.PortfolioID, &portfolio.UserID, &portfolio.StockID, &portfolio.Quantity, &portfolio.CreatedAt)
	if err != nil {
		p.log.Error("error is while selecting portfolio", logger.Error(err))
		return models.Portfolio{}, err
	}
	return portfolio, nil
}

func (p *portfoliosRepo) GetList(ctx context.Context, req models.GetListRequest) (models.PortfolioResponse, error) {
	var (
		portfolios = []models.Portfolio{}
		count      = 0
		query      = `SELECT portfolio_id, user_id, stock_id, quantity, created_at FROM portfolios`
		countQuery = `SELECT COUNT(1) FROM portfolios`
		filter     = ""
		page       = req.Page
		offset     = (page - 1) * req.Limit
		search     = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE user_id::text ILIKE '%%%s%%' OR stock_id::text ILIKE '%%%s%%'", search, search)
	}

	countQuery += filter
	err := p.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		p.log.Error("error is while selecting count", logger.Error(err))
		return models.PortfolioResponse{}, err
	}

	query += filter
	query += ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := p.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		p.log.Error("error is while selecting portfolios", logger.Error(err))
		return models.PortfolioResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var portfolio models.Portfolio
		if err := rows.Scan(&portfolio.PortfolioID, &portfolio.UserID, &portfolio.StockID, &portfolio.Quantity, &portfolio.CreatedAt); err != nil {
			p.log.Error("error is while scanning data", logger.Error(err))
			return models.PortfolioResponse{}, err
		}
		portfolios = append(portfolios, portfolio)
	}

	return models.PortfolioResponse{
		Portfolios: portfolios,
		Count:      count,
	}, nil
}

func (p *portfoliosRepo) Update(ctx context.Context, portfolio models.UpdatePortfolio) (string, error) {
	query := `UPDATE portfolios SET  stock_id = $1, quantity = $2 
	          WHERE portfolio_id = $3 RETURNING portfolio_id`
	var portfolioID int
	err := p.db.QueryRow(ctx, query, portfolio.StockID, portfolio.Quantity, portfolio.PortfolioID).Scan(&portfolioID)
	if err != nil {
		p.log.Error("error is while updating portfolio", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(portfolioID), nil
}

func (p *portfoliosRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM portfolios WHERE portfolio_id = $1`
	_, err := p.db.Exec(ctx, query, key.ID)
	if err != nil {
		p.log.Error("error is while deleting portfolio", logger.Error(err))
		return err
	}
	return nil
}
