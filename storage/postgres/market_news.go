package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type marketNewsRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewMarketNewsRepo(db *pgxpool.Pool, log logger.ILogger) storage.IMarketNewsStorage {
	return &marketNewsRepo{
		db:  db,
		log: log,
	}
}

func (m *marketNewsRepo) Create(ctx context.Context, news models.CreateMarketNews) (string, error) {
	query := `INSERT INTO market_news(headline, content) 
	          VALUES($1, $2) RETURNING news_id`
	var newsID int
	err := m.db.QueryRow(ctx, query, news.Headline, news.Content).Scan(&newsID)
	if err != nil {
		m.log.Error("error is while inserting market news", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(newsID), nil
}

func (m *marketNewsRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.MarketNews, error) {
	var marketNews models.MarketNews
	query := `SELECT news_id, headline, content, timestamp 
	          FROM market_news WHERE news_id = $1`
	err := m.db.QueryRow(ctx, query, key.ID).Scan(&marketNews.NewsID, &marketNews.Headline, &marketNews.Content, &marketNews.Timestamp)
	if err != nil {
		m.log.Error("error is while selecting market news", logger.Error(err))
		return models.MarketNews{}, err
	}
	return marketNews, nil
}

func (m *marketNewsRepo) GetList(ctx context.Context, req models.GetListRequest) (models.MarketNewsResponse, error) {
	var (
		marketNewsList = []models.MarketNews{}
		count          = 0
		query          = `SELECT news_id, headline, content, timestamp FROM market_news`
		countQuery     = `SELECT COUNT(1) FROM market_news`
		filter         = ""
		page           = req.Page
		offset         = (page - 1) * req.Limit
		search         = req.Search
	)

	if search != "" {
		filter = fmt.Sprintf(" WHERE headline ILIKE '%%%s%%' OR content ILIKE '%%%s%%'", search, search)
	}

	countQuery += filter
	err := m.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		m.log.Error("error is while selecting count", logger.Error(err))
		return models.MarketNewsResponse{}, err
	}

	query += filter
	query += ` ORDER BY timestamp DESC LIMIT $1 OFFSET $2`
	rows, err := m.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		m.log.Error("error is while selecting market news", logger.Error(err))
		return models.MarketNewsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var news models.MarketNews
		if err := rows.Scan(&news.NewsID, &news.Headline, &news.Content, &news.Timestamp); err != nil {
			m.log.Error("error is while scanning data", logger.Error(err))
			return models.MarketNewsResponse{}, err
		}
		marketNewsList = append(marketNewsList, news)
	}

	return models.MarketNewsResponse{
		NewsItems: marketNewsList,
		Count:     count,
	}, nil
}

func (m *marketNewsRepo) Update(ctx context.Context, news models.UpdateMarketNews) (string, error) {
	query := `UPDATE market_news SET headline = $1, content = $2 
	          WHERE news_id = $3 RETURNING news_id`
	var newsID int
	err := m.db.QueryRow(ctx, query, news.Headline, news.Content, news.NewsID).Scan(&newsID)
	if err != nil {
		m.log.Error("error is while updating market news", logger.Error(err))
		return "", err
	}
	return fmt.Sprint(newsID), nil
}

func (m *marketNewsRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM market_news WHERE news_id = $1`
	_, err := m.db.Exec(ctx, query, key.ID)
	if err != nil {
		m.log.Error("error is while deleting market news", logger.Error(err))
		return err
	}
	return nil
}
