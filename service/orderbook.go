package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type orderBookService struct {
	storage storage.IOrderBookStorage
	log     logger.ILogger
}

func NeworderBookService(storage storage.IOrderBookStorage, log logger.ILogger) orderBookService {
	return orderBookService{
		storage: storage,
		log:     log,
	}
}

func (s orderBookService) Create(ctx context.Context, orderBook models.CreateOrderBook) (models.OrderBook, error) {
	s.log.Info("orderBook create service layer", logger.Any("orderBook", orderBook))
	id, err := s.storage.Create(ctx, orderBook)
	if err != nil {
		s.log.Error("error in service layer while creating orderBook", logger.Error(err))
		return models.OrderBook{}, err
	}

	createdOrderBook, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting orderBook by id", logger.Error(err))
		return models.OrderBook{}, err
	}

	return createdOrderBook, nil
}

func (s orderBookService) Get(ctx context.Context, id string) (models.OrderBook, error) {
	orderBook, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting orderBook by id", logger.Error(err))
		return models.OrderBook{}, err
	}

	return orderBook, nil
}

func (s orderBookService) GetList(ctx context.Context, request models.GetListRequest) (models.OrderBookResponse, error) {
	s.log.Info("orderBook get list service layer", logger.Any("request", request))

	orderBooks, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of orderBooks", logger.Error(err))
		return models.OrderBookResponse{}, err
	}

	return orderBooks, nil
}

func (s orderBookService) Update(ctx context.Context, orderBook models.UpdateOrderBook) (models.OrderBook, error) {
	id, err := s.storage.Update(ctx, orderBook)
	if err != nil {
		s.log.Error("error in service layer while updating orderBook", logger.Error(err))
		return models.OrderBook{}, err
	}

	updatedOrderBook, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting updated orderBook by id", logger.Error(err))
		return models.OrderBook{}, err
	}

	return updatedOrderBook, nil
}

func (s orderBookService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting orderBook", logger.Error(err))
	}
	return err
}
