package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewUserRepo(db *pgxpool.Pool, log logger.ILogger) storage.IUserStorage {
	return &userRepo{
		db:  db,
		log: log,
	}
}

func (u *userRepo) Create(ctx context.Context, createUser models.CreateUser) (string, error) {

	uid := uuid.New()

	if _, err := u.db.Exec(ctx, `insert into 
			users (id, name, email, password_hash, login) values ($1, $2, $3, $4, $5)
			`,
		uid,
		createUser.Name,
		createUser.Phone,
		createUser.Email,
		createUser.PasswordHash,
		createUser.Login,
	); err != nil {
		u.log.Error("error while inserting data", logger.Error(err))
		return "", err
	}

	return uid.String(), nil
}

func (u *userRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	var createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	user := models.User{}

	query := `
		select id, name, email, password_hash, created_at, updated_at 
						from users where id = $1 and deleted_at = 0
`
	if err := u.db.QueryRow(ctx, query, pKey.ID).Scan(
		&user.ID,           //0
		&user.Name,         //1
		&user.Email,        //2
		&user.PasswordHash, //3
		&createdAt,         //4
		&updatedAt,         //5
	); err != nil {
		u.log.Error("error while scanning user", logger.Error(err))
		return models.User{}, err
	}

	if createdAt.Valid {
		user.CreatedAt = createdAt.Time
	}

	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.String
	}

	return user, nil
}

func (u *userRepo) GetList(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error) {
	var (
		users                = []models.User{}
		count                = 0
		countQuery, query    string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullTime{}, sql.NullString{}
	)

	countQuery = `
		SELECT count(1) from users where deleted_at = 0 `

	if search != "" {
		countQuery += fmt.Sprintf(` and (phone ilike '%s' or full_name ilike '%s')`, search, search)
	}

	if err := u.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of users", err.Error())
		return models.UsersResponse{}, err
	}

	query = `
		id, name, email, password_hash, created_at, updated_at 
						from users where id = $1 and deleted_at = 0
			    `

	if search != "" {
		query += fmt.Sprintf(` and (phone ilike '%s' or full_name ilike '%s') `, search, search)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2`

	rows, err := u.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.UsersResponse{}, err
	}

	for rows.Next() {
		user := models.User{}

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PasswordHash,
			&createdAt,
			&updatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.UsersResponse{}, err
		}
		if createdAt.Valid {
			user.CreatedAt = createdAt.Time
		}

		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.String
		}
		users = append(users, user)
	}

	return models.UsersResponse{
		Users: users,
		Count: count,
	}, nil
}

func (u *userRepo) Update(ctx context.Context, request models.UpdateUser) (string, error) {
	query := `
		update users 
			set full_name = $1, phone = $2, cash = $3, updated_at = now()
				where user_role = 'customer' and id = $4`

	if _, err := u.db.Exec(ctx, query, request.Name, request.Email,request.ID); err != nil {
		fmt.Println("error while updating user data", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (u *userRepo) Delete(ctx context.Context, request models.PrimaryKey) error {
	query := `update users set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if _, err := u.db.Exec(ctx, query, request.ID); err != nil {
		fmt.Println("error while deleting user by id", err.Error())
		return err
	}

	return nil
}

func (u *userRepo) GetPassword(ctx context.Context, id string) (string, error) {
	password := ""

	query := `
		select password from users 
		                where  id = $1`

	if err := u.db.QueryRow(ctx, query, id).Scan(&password); err != nil {
		fmt.Println("Error while scanning password from users", err.Error())
		return "", err
	}

	return password, nil
}

func (u *userRepo) UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error {
	query := `
		update users 
				set password = $1, updated_at = now()
					where id = $2 `

	if _, err := u.db.Exec(ctx, query, request.NewPassword, request.ID); err != nil {
		fmt.Println("error while updating password for user", err.Error())
		return err
	}

	return nil
}

func (u *userRepo) GetAdminCredentialsByLogin(ctx context.Context, login string) (models.User, error) {
	user := models.User{}

	query := `
		select id, password from users 
		                where  login = $1`

	if err := u.db.QueryRow(ctx, query, login).Scan(&user.ID, &user.PasswordHash); err != nil {
		fmt.Println("Error while scanning password from users", err.Error())
		return models.User{}, err
	}

	return user, nil
}
