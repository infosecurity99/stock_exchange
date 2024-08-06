package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type fundTransfersService struct {
	storage storage.IFundTransfersStorage
	log     logger.ILogger
}

func NewfundTransfersService(storage storage.IFundTransfersStorage, log logger.ILogger) fundTransfersService {
	return fundTransfersService{
		storage: storage,
		log:     log,
	}
}

func (s fundTransfersService) Create(ctx context.Context, fundTransfer models.CreateFundTransfer) (models.FundTransfer, error) {
	s.log.Info("fund transfer create service layer", logger.Any("fundTransfer", fundTransfer))
	id, err := s.storage.Create(ctx, fundTransfer)
	if err != nil {
		s.log.Error("error in service layer while creating fund transfer", logger.Error(err))
		return models.FundTransfer{}, err
	}

	createdFundTransfer, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting fund transfer by id", logger.Error(err))
		return models.FundTransfer{}, err
	}

	return createdFundTransfer, nil
}

func (s fundTransfersService) Get(ctx context.Context, id string) (models.FundTransfer, error) {
	fundTransfer, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting fund transfer by id", logger.Error(err))
		return models.FundTransfer{}, err
	}

	return fundTransfer, nil
}

func (s fundTransfersService) GetList(ctx context.Context, request models.GetListRequest) (models.FundTransferResponse, error) {
	s.log.Info("fund transfer get list service layer", logger.Any("request", request))

	fundTransfers, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of fund transfers", logger.Error(err))
		return models.FundTransferResponse{}, err
	}

	return fundTransfers, nil
}

func (s fundTransfersService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting fund transfer", logger.Error(err))
	}
	return err
}
