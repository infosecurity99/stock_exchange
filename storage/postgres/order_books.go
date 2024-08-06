package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type orderBookRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewOrderBookRepo(db *pgxpool.Pool, log logger.ILogger) storage.IOrderBookStorage {
	return &orderBookRepo{
		db:  db,
		log: log,
	}
}

func (o *orderBookRepo) Create(ctx context.Context, orderBook models.CreateOrderBook) (string, error) {
	query := `INSERT INTO order_book(order_id, order_type, quantity, price) 
	          VALUES($1, $2, $3, $4) RETURNING order_id`
	var orderID int
	err := o.db.QueryRow(ctx, query, orderBook.OrderID, orderBook.OrderType, orderBook.Quantity, orderBook.Price).Scan(&orderID)
	if err != nil {
		o.log.Error("error while inserting order book", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(orderID), nil
}

func (o *orderBookRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.OrderBook, error) {
	var orderBook models.OrderBook
	query := `SELECT order_id, order_type, quantity, price FROM order_book WHERE order_id = $1`
	err := o.db.QueryRow(ctx, query, key.ID).Scan(&orderBook.OrderID, &orderBook.OrderType, &orderBook.Quantity, &orderBook.Price)
	if err != nil {
		o.log.Error("error while selecting order book", logger.Error(err))
		return models.OrderBook{}, err
	}
	return orderBook, nil
}

func (o *orderBookRepo) GetList(ctx context.Context, req models.GetListRequest) (models.OrderBookResponse, error) {
	var (
		orderBooks = []models.OrderBook{}
		count      = 0
		query      = `SELECT order_id, order_type, quantity, price FROM order_book`
		countQuery = `SELECT COUNT(1) FROM order_book`
		filter     = ""
		page       = req.Page
		offset     = (page - 1) * req.Limit
		search     = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE order_type ILIKE '%%%s%%'", search)
	}

	countQuery += filter
	err := o.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		o.log.Error("error while selecting count", logger.Error(err))
		return models.OrderBookResponse{}, err
	}

	query += filter
	query += ` ORDER BY order_id DESC LIMIT $1 OFFSET $2`
	rows, err := o.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		o.log.Error("error while selecting order books", logger.Error(err))
		return models.OrderBookResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderBook models.OrderBook
		if err := rows.Scan(&orderBook.OrderID, &orderBook.OrderType, &orderBook.Quantity, &orderBook.Price); err != nil {
			o.log.Error("error while scanning data", logger.Error(err))
			return models.OrderBookResponse{}, err
		}
		orderBooks = append(orderBooks, orderBook)
	}

	return models.OrderBookResponse{
		Entries: orderBooks,
		Count:      count,
	}, nil
}

func (o *orderBookRepo) Update(ctx context.Context, orderBook models.UpdateOrderBook) (string, error) {
	query := `UPDATE order_book SET order_type = $1, quantity = $2, price = $3 WHERE order_id = $4 RETURNING order_id`
	var orderID int
	err := o.db.QueryRow(ctx, query, orderBook.OrderType, orderBook.Quantity, orderBook.Price, orderBook.OrderID).Scan(&orderID)
	if err != nil {
		o.log.Error("error while updating order book", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(orderID), nil
}

func (o *orderBookRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM order_book WHERE order_id = $1`
	_, err := o.db.Exec(ctx, query, key.ID)
	if err != nil {
		o.log.Error("error while deleting order book", logger.Error(err))
		return err
	}
	return nil
}
