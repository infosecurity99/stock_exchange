package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type marketNewsService struct {
	storage storage.IMarketNewsStorage
	log     logger.ILogger
}

func NewmarketNewsService(storage storage.IMarketNewsStorage, log logger.ILogger) marketNewsService {
	return marketNewsService{
		storage: storage,
		log:     log,
	}
}

func (s marketNewsService) Create(ctx context.Context, marketNews models.CreateMarketNews) (models.MarketNews, error) {
	s.log.Info("market news create service layer", logger.Any("marketNews", marketNews))
	id, err := s.storage.Create(ctx, marketNews)
	if err != nil {
		s.log.Error("error in service layer while creating market news", logger.Error(err))
		return models.MarketNews{}, err
	}

	createdMarketNews, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting market news by id", logger.Error(err))
		return models.MarketNews{}, err
	}

	return createdMarketNews, nil
}

func (s marketNewsService) Get(ctx context.Context, id string) (models.MarketNews, error) {
	marketNews, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting market news by id", logger.Error(err))
		return models.MarketNews{}, err
	}

	return marketNews, nil
}

func (s marketNewsService) GetList(ctx context.Context, request models.GetListRequest) (models.MarketNewsResponse, error) {
	s.log.Info("market news get list service layer", logger.Any("request", request))

	marketNewsList, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of market news", logger.Error(err))
		return models.MarketNewsResponse{}, err
	}

	return marketNewsList, nil
}

func (s marketNewsService) Update(ctx context.Context, marketNews models.UpdateMarketNews) (models.MarketNews, error) {
	id, err := s.storage.Update(ctx, marketNews)
	if err != nil {
		s.log.Error("error in service layer while updating market news", logger.Error(err))
		return models.MarketNews{}, err
	}

	updatedMarketNews, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting updated market news by id", logger.Error(err))
		return models.MarketNews{}, err
	}

	return updatedMarketNews, nil
}

func (s marketNewsService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting market news", logger.Error(err))
	}
	return err
}
