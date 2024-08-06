package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type marketUpdatesService struct {
	storage storage.IMarketUpdatesStorage
	log     logger.ILogger
}

func NewmarketUpdatesService(storage storage.IMarketNewsStorage, log logger.ILogger) marketUpdatesService {
	return marketUpdatesService{
		storage: storage,
		log:     log,
	}
}

func (s marketUpdatesService) Create(ctx context.Context, update models.CreateMarketUpdate) (models.MarketUpdate, error) {
	s.log.Info("market updates create service layer", logger.Any("update", update))
	id, err := s.storage.Create(ctx, update)
	if err != nil {
		s.log.Error("error in service layer while creating market update", logger.Error(err))
		return models.MarketUpdate{}, err
	}

	createdUpdate, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting market update by id", logger.Error(err))
		return models.MarketUpdate{}, err
	}

	return createdUpdate, nil
}

func (s marketUpdatesService) Get(ctx context.Context, id string) (models.MarketUpdate, error) {
	update, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting market update by id", logger.Error(err))
		return models.MarketUpdate{}, err
	}

	return update, nil
}

func (s marketUpdatesService) GetList(ctx context.Context, request models.GetListRequest) (models.MarketUpdateResponse, error) {
	s.log.Info("market updates get list service layer", logger.Any("request", request))

	updates, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting market updates list", logger.Error(err))
		return models.MarketUpdateResponse{}, err
	}

	return updates, nil
}

func (s marketUpdatesService) Update(ctx context.Context, update models.UpdateMarketUpdate) (models.MarketUpdate, error) {
	id, err := s.storage.Update(ctx, update)
	if err != nil {
		s.log.Error("error in service layer while updating market update", logger.Error(err))
		return models.MarketUpdate{}, err
	}

	updatedUpdate, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting updated market update by id", logger.Error(err))
		return models.MarketUpdate{}, err
	}

	return updatedUpdate, nil
}

func (s marketUpdatesService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting market update", logger.Error(err))
	}
	return err
}
