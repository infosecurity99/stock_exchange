package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountInformationRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewAccountInforamtionRepo(db *pgxpool.Pool, log logger.ILogger) storage.IAccountINforamtionStorage {
	return &accountInformationRepo{
		db:  db,
		log: log,
	}
}

func (a *accountInformationRepo) Create(ctx context.Context, info models.CreateAccountInformation) (string, error) {
	query := `INSERT INTO account_information(user_id, balance) 
	          VALUES($1, $2) RETURNING account_id`
	var accountID int
	err := a.db.QueryRow(ctx, query, info.UserID, info.Balance).Scan(&accountID)
	if err != nil {
		a.log.Error("error is while inserting account information", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(accountID), nil
}

func (a *accountInformationRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.AccountInformation, error) {
	var info models.AccountInformation
	query := `SELECT account_id, user_id, balance, created_at 
	          FROM account_information WHERE account_id = $1`
	err := a.db.QueryRow(ctx, query, key.ID).Scan(&info.AccountID, &info.UserID, &info.Balance, &info.CreatedAt)
	if err != nil {
		a.log.Error("error is while selecting account information", logger.Error(err))
		return models.AccountInformation{}, err
	}
	return info, nil
}

func (a *accountInformationRepo) GetList(ctx context.Context, req models.GetListRequest) (models.AccountInformationResponse, error) {
	var (
		informations = []models.AccountInformation{}
		count        = 0
		query        = `SELECT account_id, user_id, balance, created_at FROM account_information`
		countQuery   = `SELECT COUNT(1) FROM account_information`
		filter       = ""
		page         = req.Page
		offset       = (page - 1) * req.Limit
		search       = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE user_id::text ILIKE '%%%s%%' OR balance::text ILIKE '%%%s%%'", search, search)
	}

	countQuery += filter
	err := a.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		a.log.Error("error is while selecting count", logger.Error(err))
		return models.AccountInformationResponse{}, err
	}

	query += filter
	query += ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := a.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		a.log.Error("error is while selecting account information", logger.Error(err))
		return models.AccountInformationResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var info models.AccountInformation
		if err := rows.Scan(&info.AccountID, &info.UserID, &info.Balance, &info.CreatedAt); err != nil {
			a.log.Error("error is while scanning data", logger.Error(err))
			return models.AccountInformationResponse{}, err
		}
		informations = append(informations, info)
	}

	return models.AccountInformationResponse{
		Accounts: informations,
		Count:               count,
	}, nil
}

func (a *accountInformationRepo) Update(ctx context.Context, info models.UpdateAccountInformation) (string, error) {
	query := `UPDATE account_information SET account_id = $1, balance = $2 
	          WHERE account_id = $3 RETURNING account_id`
	var accountID int
	err := a.db.QueryRow(ctx, query, info.AccountID, info.Balance, info.AccountID).Scan(&accountID)
	if err != nil {
		a.log.Error("error is while updating account information", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(accountID), nil
}

func (a *accountInformationRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM account_information WHERE account_id = $1`
	_, err := a.db.Exec(ctx, query, key.ID)
	if err != nil {
		a.log.Error("error is while deleting account information", logger.Error(err))
		return err
	}
	return nil
}
