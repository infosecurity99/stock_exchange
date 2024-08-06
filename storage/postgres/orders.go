package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ordersRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewOrdersRepo(db *pgxpool.Pool, log logger.ILogger) storage.IOrdersStorage {
	return &ordersRepo{
		db:  db,
		log: log,
	}
}

func (o *ordersRepo) Create(ctx context.Context, order models.CreateOrder) (string, error) {
	query := `INSERT INTO orders(user_id, stock_id, order_type, quantity, price, status) 
	          VALUES($1, $2, $3, $4, $5, $6) RETURNING order_id`
	var orderID int
	err := o.db.QueryRow(ctx, query, order.UserID, order.StockID, order.OrderType, order.Quantity, order.Price).Scan(&orderID)
	if err != nil {
		o.log.Error("error is while inserting order", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(orderID), nil
}

func (o *ordersRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Order, error) {
	var order models.Order
	query := `SELECT order_id, user_id, stock_id, order_type, quantity, price, status, created_at 
	          FROM orders WHERE order_id = $1`
	err := o.db.QueryRow(ctx, query, key.ID).Scan(&order.OrderID, &order.UserID, &order.StockID, &order.OrderType, &order.Quantity, &order.Price, &order.Status, &order.CreatedAt)
	if err != nil {
		o.log.Error("error is while selecting order", logger.Error(err))
		return models.Order{}, err
	}
	return order, nil
}

func (o *ordersRepo) GetList(ctx context.Context, req models.GetListRequest) (models.OrderResponse, error) {
	var (
		orders     = []models.Order{}
		count      = 0
		query      = `SELECT order_id, user_id, stock_id, order_type, quantity, price, status, created_at FROM orders`
		countQuery = `SELECT COUNT(1) FROM orders`
		filter     = ""
		page       = req.Page
		offset     = (page - 1) * req.Limit
		search     = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE order_type ILIKE '%%%s%%' OR status ILIKE '%%%s%%'", search, search)
	}

	countQuery += filter
	err := o.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		o.log.Error("error is while selecting count", logger.Error(err))
		return models.OrderResponse{}, err
	}

	query += filter
	query += ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := o.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		o.log.Error("error is while selecting orders", logger.Error(err))
		return models.OrderResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.StockID, &order.OrderType, &order.Quantity, &order.Price, &order.Status, &order.CreatedAt); err != nil {
			o.log.Error("error is while scanning data", logger.Error(err))
			return models.OrderResponse{}, err
		}
		orders = append(orders, order)
	}

	return models.OrderResponse{
		Orders: orders,
		Count:  count,
	}, nil
}

func (o *ordersRepo) Update(ctx context.Context, order models.UpdateOrder) (string, error) {
	query := `UPDATE orders SET user_id = $1, stock_id = $2, order_type = $3, quantity = $4, price = $5, status = $6 
	          WHERE order_id = $7 RETURNING order_id`
	var orderID int
	err := o.db.QueryRow(ctx, query, order.UserID, order.StockID, order.OrderType, order.Quantity, order.Price, order.Status, order.OrderID).Scan(&orderID)
	if err != nil {
		o.log.Error("error is while updating order", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(orderID), nil
}

func (o *ordersRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM orders WHERE order_id = $1`
	_, err := o.db.Exec(ctx, query, key.ID)
	if err != nil {
		o.log.Error("error is while deleting order", logger.Error(err))
		return err
	}
	return nil
}
