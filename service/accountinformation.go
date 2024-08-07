package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type accountInforamtionService struct {
	storage storage.IAccountInforamtionStorage
	log     logger.ILogger
}

func NewAccountInforamtionServiceService(storage storage.IAccountInforamtionStorage, log logger.ILogger) accountInforamtionService {
	return accountInforamtionService{
		storage: storage,
		log:     log,
	}
}

func (s accountInforamtionService) Create(ctx context.Context, accountInfo models.CreateAccountInformation) (models.AccountInformation, error) {
	s.log.Info("account information create service layer", logger.Any("accountInfo", accountInfo))
	id, err := s.storage.Create(ctx, accountInfo)
	if err != nil {
		s.log.Error("error in service layer while creating account information", logger.Error(err))
		return models.AccountInformation{}, err
	}

	createdAccountInfo, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting account information by id", logger.Error(err))
		return models.AccountInformation{}, err
	}

	return createdAccountInfo, nil
}

func (s accountInforamtionService) Get(ctx context.Context, id string) (models.AccountInformation, error) {
	accountInfo, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting account information by id", logger.Error(err))
		return models.AccountInformation{}, err
	}

	return accountInfo, nil
}

func (s accountInforamtionService) GetList(ctx context.Context, request models.GetListRequest) (models.AccountInformationResponse, error) {
	s.log.Info("account information get list service layer", logger.Any("request", request))

	accountInfos, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of account information", logger.Error(err))
		return models.AccountInformationResponse{}, err
	}

	return accountInfos, nil
}

func (s accountInforamtionService) Update(ctx context.Context, accountInfo models.UpdateAccountInformation) (models.AccountInformation, error) {
	id, err := s.storage.Update(ctx, accountInfo)
	if err != nil {
		s.log.Error("error in service layer while updating account information", logger.Error(err))
		return models.AccountInformation{}, err
	}

	updatedAccountInfo, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting updated account information by id", logger.Error(err))
		return models.AccountInformation{}, err
	}

	return updatedAccountInfo, nil
}

func (s accountInforamtionService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting account information", logger.Error(err))
	}
	return err
}
