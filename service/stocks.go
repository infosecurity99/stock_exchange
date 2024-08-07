package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type stocksService struct {
	storage storage.IStocksStorage
	log     logger.ILogger
}


func NewStocksService(storage storage.IStocksStorage, log logger.ILogger) stocksService {
    return stocksService{
        storage: storage,
        log:     log,
    }
}

func (s stocksService) Create(ctx context.Context, stock models.CreateStock) (models.Stock, error) {
	s.log.Info("stocks create service layer", logger.Any("stock", stock))
	id, err := s.storage.Create(ctx, stock)
	if err != nil {
		s.log.Error("error in service layer while creating stock", logger.Error(err))
		return models.Stock{}, err
	}

	createdStock, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error is while getting by id", logger.Error(err))
		return models.Stock{}, err
	}

	return createdStock, nil
}

func (s stocksService) Get(ctx context.Context, id string) (models.Stock, error) {
	stock, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Stock{}, err
	}

	return stock, nil
}

func (s stocksService) GetList(ctx context.Context, request models.GetListRequest) (models.StockResponse, error) {
	s.log.Info("stocks get list service layer", logger.Any("request", request))

	stocks, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list", logger.Error(err))
		return models.StockResponse{}, err
	}

	return stocks, nil
}

func (s stocksService) Update(ctx context.Context, stock models.UpdateStock) (models.Stock, error) {
	id, err := s.storage.Update(ctx, stock)
	if err != nil {
		s.log.Error("error in service layer while updating stock", logger.Error(err))
		return models.Stock{}, err
	}

	updatedStock, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting stock by id", logger.Error(err))
		return models.Stock{}, err
	}

	return updatedStock, nil
}

func (s stocksService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting stock", logger.Error(err))
	}
	return err
}
