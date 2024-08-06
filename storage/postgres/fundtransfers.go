package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type fundTransfersRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewFundTransfersRepo(db *pgxpool.Pool, log logger.ILogger) storage.IFundTransfersStorage {
	return &fundTransfersRepo{
		db:  db,
		log: log,
	}
}

func (f *fundTransfersRepo) Create(ctx context.Context, transfer models.CreateFundTransfer) (string, error) {
	query := `INSERT INTO fund_transfers(from_user_id, to_user_id, amount) 
	          VALUES($1, $2, $3) RETURNING transfer_id`
	var transferID int
	err := f.db.QueryRow(ctx, query, transfer.FromUserID, transfer.ToUserID, transfer.Amount).Scan(&transferID)
	if err != nil {
		f.log.Error("error is while inserting fund transfer", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(transferID), nil
}

func (f *fundTransfersRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.FundTransfer, error) {
	var transfer models.FundTransfer
	query := `SELECT transfer_id, from_user_id, to_user_id, amount, timestamp 
	          FROM fund_transfers WHERE transfer_id = $1`
	err := f.db.QueryRow(ctx, query, key.ID).Scan(&transfer.TransferID, &transfer.FromUserID, &transfer.ToUserID, &transfer.Amount, &transfer.Timestamp)
	if err != nil {
		f.log.Error("error is while selecting fund transfer", logger.Error(err))
		return models.FundTransfer{}, err
	}
	return transfer, nil
}

func (f *fundTransfersRepo) GetList(ctx context.Context, req models.GetListRequest) (models.FundTransferResponse, error) {
	var (
		transfers  = []models.FundTransfer{}
		count      = 0
		query      = `SELECT transfer_id, from_user_id, to_user_id, amount, timestamp FROM fund_transfers`
		countQuery = `SELECT COUNT(1) FROM fund_transfers`
		filter     = ""
		page       = req.Page
		offset     = (page - 1) * req.Limit
		search     = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE from_user_id::text ILIKE '%%%s%%' OR to_user_id::text ILIKE '%%%s%%'", search, search)
	}

	countQuery += filter
	err := f.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		f.log.Error("error is while selecting count", logger.Error(err))
		return models.FundTransferResponse{}, err
	}

	query += filter
	query += ` ORDER BY timestamp DESC LIMIT $1 OFFSET $2`
	rows, err := f.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		f.log.Error("error is while selecting fund transfers", logger.Error(err))
		return models.FundTransferResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var transfer models.FundTransfer
		if err := rows.Scan(&transfer.TransferID, &transfer.FromUserID, &transfer.ToUserID, &transfer.Amount, &transfer.Timestamp); err != nil {
			f.log.Error("error is while scanning data", logger.Error(err))
			return models.FundTransferResponse{}, err
		}
		transfers = append(transfers, transfer)
	}

	return models.FundTransferResponse{
		Transfers:   transfers,
		Count: count,
	}, nil
}

func (f *fundTransfersRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM fund_transfers WHERE transfer_id = $1`
	_, err := f.db.Exec(ctx, query, key.ID)
	if err != nil {
		f.log.Error("error is while deleting fund transfer", logger.Error(err))
		return err
	}
	return nil
}
