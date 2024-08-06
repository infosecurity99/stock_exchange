package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type tradeConfirmationsService struct {
	storage storage.ITradeConfirmationsStorage
	log     logger.ILogger
}

func NewtradeConfirmationsService(storage storage.ITradeConfirmationsStorage, log logger.ILogger) tradeConfirmationsService {
	return tradeConfirmationsService{
		storage: storage,
		log:     log,
	}
}

func (s tradeConfirmationsService) Create(ctx context.Context, tradeConfirmation models.CreateTradeConfirmation) (models.TradeConfirmation, error) {
	s.log.Info("trade confirmation create service layer", logger.Any("tradeConfirmation", tradeConfirmation))
	id, err := s.storage.Create(ctx, tradeConfirmation)
	if err != nil {
		s.log.Error("error in service layer while creating trade confirmation", logger.Error(err))
		return models.TradeConfirmation{}, err
	}

	createdTradeConfirmation, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting trade confirmation by id", logger.Error(err))
		return models.TradeConfirmation{}, err
	}

	return createdTradeConfirmation, nil
}

func (s tradeConfirmationsService) Get(ctx context.Context, id string) (models.TradeConfirmation, error) {
	tradeConfirmation, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting trade confirmation by id", logger.Error(err))
		return models.TradeConfirmation{}, err
	}

	return tradeConfirmation, nil
}

func (s tradeConfirmationsService) GetList(ctx context.Context, request models.GetListRequest) (models.TradeConfirmationResponse, error) {
	s.log.Info("trade confirmation get list service layer", logger.Any("request", request))

	tradeConfirmations, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of trade confirmations", logger.Error(err))
		return models.TradeConfirmationResponse{}, err
	}

	return tradeConfirmations, nil
}

func (s tradeConfirmationsService) Update(ctx context.Context, tradeConfirmation models.UpdateTradeConfirmation) (models.TradeConfirmation, error) {
	id, err := s.storage.Update(ctx, tradeConfirmation)
	if err != nil {
		s.log.Error("error in service layer while updating trade confirmation", logger.Error(err))
		return models.TradeConfirmation{}, err
	}

	updatedTradeConfirmation, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting trade confirmation by id", logger.Error(err))
		return models.TradeConfirmation{}, err
	}

	return updatedTradeConfirmation, nil
}

func (s tradeConfirmationsService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting trade confirmation", logger.Error(err))
	}
	return err
}
