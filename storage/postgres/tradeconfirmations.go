package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type tradeConfirmationsRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewTradeConfirmationsRepo(db *pgxpool.Pool, log logger.ILogger) storage.ITradeConfirmationsStorage {
	return &tradeConfirmationsRepo{
		db:  db,
		log: log,
	}
}

func (t *tradeConfirmationsRepo) Create(ctx context.Context, confirmation models.CreateTradeConfirmation) (string, error) {
	query := `INSERT INTO trade_confirmations(order_id, user_id, stock_id, quantity, price) 
	          VALUES($1, $2, $3, $4, $5) RETURNING confirmation_id`
	var confirmationID int
	err := t.db.QueryRow(ctx, query, confirmation.OrderID, confirmation.UserID, confirmation.StockID, confirmation.Quantity, confirmation.Price).Scan(&confirmationID)
	if err != nil {
		t.log.Error("error is while inserting trade confirmation", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(confirmationID), nil
}

func (t *tradeConfirmationsRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.TradeConfirmation, error) {
	var confirmation models.TradeConfirmation
	query := `SELECT confirmation_id, order_id, user_id, stock_id, quantity, price, timestamp 
	          FROM trade_confirmations WHERE confirmation_id = $1`
	err := t.db.QueryRow(ctx, query, key.ID).Scan(&confirmation.ConfirmationID, &confirmation.OrderID, &confirmation.UserID, &confirmation.StockID, &confirmation.Quantity, &confirmation.Price, &confirmation.Timestamp)
	if err != nil {
		t.log.Error("error is while selecting trade confirmation", logger.Error(err))
		return models.TradeConfirmation{}, err
	}
	return confirmation, nil
}

func (t *tradeConfirmationsRepo) GetList(ctx context.Context, req models.GetListRequest) (models.TradeConfirmationResponse, error) {
	var (
		confirmations = []models.TradeConfirmation{}
		count         = 0
		query         = `SELECT confirmation_id, order_id, user_id, stock_id, quantity, price, timestamp FROM trade_confirmations`
		countQuery    = `SELECT COUNT(1) FROM trade_confirmations`
		filter        = ""
		page          = req.Page
		offset        = (page - 1) * req.Limit
		search        = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE order_id::text ILIKE '%%%s%%' OR user_id::text ILIKE '%%%s%%' OR stock_id::text ILIKE '%%%s%%'", search, search, search)
	}

	countQuery += filter
	err := t.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		t.log.Error("error is while selecting count", logger.Error(err))
		return models.TradeConfirmationResponse{}, err
	}

	query += filter
	query += ` ORDER BY timestamp DESC LIMIT $1 OFFSET $2`
	rows, err := t.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		t.log.Error("error is while selecting trade confirmations", logger.Error(err))
		return models.TradeConfirmationResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var confirmation models.TradeConfirmation
		if err := rows.Scan(&confirmation.ConfirmationID, &confirmation.OrderID, &confirmation.UserID, &confirmation.StockID, &confirmation.Quantity, &confirmation.Price, &confirmation.Timestamp); err != nil {
			t.log.Error("error is while scanning data", logger.Error(err))
			return models.TradeConfirmationResponse{}, err
		}
		confirmations = append(confirmations, confirmation)
	}

	return models.TradeConfirmationResponse{
		Confirmations: confirmations,
		Count:              count,
	}, nil
}

func (t *tradeConfirmationsRepo) Update(ctx context.Context, confirmation models.UpdateTradeConfirmation) (string, error) {
	query := `UPDATE trade_confirmations SET order_id = $1, user_id = $2, stock_id = $3, quantity = $4, price = $5 
	          WHERE confirmation_id = $6 RETURNING confirmation_id`
	var confirmationID int
	err := t.db.QueryRow(ctx, query, confirmation.OrderID, confirmation.UserID, confirmation.StockID, confirmation.Quantity, confirmation.Price, confirmation.ConfirmationID).Scan(&confirmationID)
	if err != nil {
		t.log.Error("error is while updating trade confirmation", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(confirmationID), nil
}

func (t *tradeConfirmationsRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM trade_confirmations WHERE confirmation_id = $1`
	_, err := t.db.Exec(ctx, query, key.ID)
	if err != nil {
		t.log.Error("error is while deleting trade confirmation", logger.Error(err))
		return err
	}
	return nil
}
