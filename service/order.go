package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type ordersService struct {
	storage storage.IOrdersStorage
	log     logger.ILogger
}

func NewordersService(storage storage.IOrdersStorage, log logger.ILogger) ordersService {
	return ordersService{
		storage: storage,
		log:     log,
	}
}

func (s ordersService) Create(ctx context.Context, order models.CreateOrder) (models.Order, error) {
	s.log.Info("orders create service layer", logger.Any("order", order))
	id, err := s.storage.Create(ctx, order)
	if err != nil {
		s.log.Error("error in service layer while creating order", logger.Error(err))
		return models.Order{}, err
	}

	createdOrder, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting order by id", logger.Error(err))
		return models.Order{}, err
	}

	return createdOrder, nil
}

func (s ordersService) Get(ctx context.Context, id string) (models.Order, error) {
	order, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting order by id", logger.Error(err))
		return models.Order{}, err
	}

	return order, nil
}

func (s ordersService) GetList(ctx context.Context, request models.GetListRequest) (models.OrderResponse, error) {
	s.log.Info("orders get list service layer", logger.Any("request", request))

	orders, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of orders", logger.Error(err))
		return models.OrderResponse{}, err
	}

	return orders, nil
}

func (s ordersService) Update(ctx context.Context, order models.UpdateOrder) (models.Order, error) {
	id, err := s.storage.Update(ctx, order)
	if err != nil {
		s.log.Error("error in service layer while updating order", logger.Error(err))
		return models.Order{}, err
	}

	updatedOrder, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting updated order by id", logger.Error(err))
		return models.Order{}, err
	}

	return updatedOrder, nil
}

func (s ordersService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting order", logger.Error(err))
	}
	return err
}
