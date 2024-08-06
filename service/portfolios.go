package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type portfoliosService struct {
	storage storage.IPortfoliosStorage
	log     logger.ILogger
}

func NewportfoliosService(storage storage.IPortfoliosStorage, log logger.ILogger) portfoliosService {
	return portfoliosService{
		storage: storage,
		log:     log,
	}
}

func (s portfoliosService) Create(ctx context.Context, portfolio models.CreatePortfolio) (models.Portfolio, error) {
	s.log.Info("portfolios create service layer", logger.Any("portfolio", portfolio))
	id, err := s.storage.Create(ctx, portfolio)
	if err != nil {
		s.log.Error("error in service layer while creating portfolio", logger.Error(err))
		return models.Portfolio{}, err
	}

	createdPortfolio, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting portfolio by id", logger.Error(err))
		return models.Portfolio{}, err
	}

	return createdPortfolio, nil
}

func (s portfoliosService) Get(ctx context.Context, id string) (models.Portfolio, error) {
	portfolio, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error in service layer while getting portfolio by id", logger.Error(err))
		return models.Portfolio{}, err
	}

	return portfolio, nil
}

func (s portfoliosService) GetList(ctx context.Context, request models.GetListRequest) (models.PortfolioResponse, error) {
	s.log.Info("portfolios get list service layer", logger.Any("request", request))

	portfolios, err := s.storage.GetList(ctx, request)
	if err != nil {
		s.log.Error("error in service layer while getting list of portfolios", logger.Error(err))
		return models.PortfolioResponse{}, err
	}

	return portfolios, nil
}

func (s portfoliosService) Update(ctx context.Context, portfolio models.UpdatePortfolio) (models.Portfolio, error) {
	id, err := s.storage.Update(ctx, portfolio)
	if err != nil {
		s.log.Error("error in service layer while updating portfolio", logger.Error(err))
		return models.Portfolio{}, err
	}

	updatedPortfolio, err := s.storage.GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		s.log.Error("error while getting updated portfolio by id", logger.Error(err))
		return models.Portfolio{}, err
	}

	return updatedPortfolio, nil
}

func (s portfoliosService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := s.storage.Delete(ctx, key)
	if err != nil {
		s.log.Error("error in service layer while deleting portfolio", logger.Error(err))
	}
	return err
}
