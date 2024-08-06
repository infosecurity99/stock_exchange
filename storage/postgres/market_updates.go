package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type marketUpdatesRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewMarketUpdatesRepo(db *pgxpool.Pool, log logger.ILogger) storage.IMarketUpdatesStorage {
	return &marketUpdatesRepo{
		db:  db,
		log: log,
	}
}

func (m *marketUpdatesRepo) Create(ctx context.Context, update models.CreateMarketUpdate) (string, error) {
	query := `INSERT INTO market_updates(stock_id, price) VALUES($1, $2) RETURNING update_id`
	var updateID int
	err := m.db.QueryRow(ctx, query, update.StockID, update.Price).Scan(&updateID)
	if err != nil {
		m.log.Error("error is while inserting market update", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(updateID), nil
}

func (m *marketUpdatesRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.MarketUpdate, error) {
	var update models.MarketUpdate
	query := `SELECT update_id, stock_id, price, timestamp FROM market_updates WHERE update_id = $1`
	err := m.db.QueryRow(ctx, query, key.ID).Scan(&update.UpdateID, &update.StockID, &update.Price, &update.Timestamp)
	if err != nil {
		m.log.Error("error is while selecting market update", logger.Error(err))
		return models.MarketUpdate{}, err
	}
	return update, nil
}

func (m *marketUpdatesRepo) GetList(ctx context.Context, req models.GetListRequest) (models.MarketUpdateResponse, error) {
	var (
		updates    = []models.MarketUpdate{}
		count      = 0
		query      = `SELECT update_id, stock_id, price, timestamp FROM market_updates`
		countQuery = `SELECT COUNT(1) FROM market_updates`
		filter     = ""
		page       = req.Page
		offset     = (page - 1) * req.Limit
		search     = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE price::TEXT ILIKE '%%%s%%'", search)
	}

	countQuery += filter
	err := m.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		m.log.Error("error is while selecting count", logger.Error(err))
		return models.MarketUpdateResponse{}, err
	}

	query += filter
	query += ` ORDER BY timestamp DESC LIMIT $1 OFFSET $2`
	rows, err := m.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		m.log.Error("error is while selecting market updates", logger.Error(err))
		return models.MarketUpdateResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var update models.MarketUpdate
		if err := rows.Scan(&update.UpdateID, &update.StockID, &update.Price, &update.Timestamp); err != nil {
			m.log.Error("error is while scanning data", logger.Error(err))
			return models.MarketUpdateResponse{}, err
		}
		updates = append(updates, update)
	}

	return models.MarketUpdateResponse{
		MarketUpdates: updates,
		Count:         count,
	}, nil
}

func (m *marketUpdatesRepo) Update(ctx context.Context, update models.UpdateMarketUpdate) (string, error) {
	query := `UPDATE market_updates SET stock_id = $1, price = $2 WHERE update_id = $3 RETURNING update_id`
	var updateID int
	err := m.db.QueryRow(ctx, query, update.StockID, update.Price, update.UpdateID).Scan(&updateID)
	if err != nil {
		m.log.Error("error is while updating market update", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(updateID), nil
}

func (m *marketUpdatesRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM market_updates WHERE update_id = $1`
	_, err := m.db.Exec(ctx, query, key.ID)
	if err != nil {
		m.log.Error("error is while deleting market update", logger.Error(err))
		return err
	}
	return nil
}
