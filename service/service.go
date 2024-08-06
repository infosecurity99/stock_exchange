package service

import (
	"test/pkg/logger"
	"test/storage"
)

type IServiceManager interface {
	User() userService

	Stocks() stocksService
	MarketUpdates() marketUpdatesService
	Orders() ordersService
	OrderBook() orderBookService
	Porfolios() portfoliosService
	TradeConfirmations() tradeConfirmationsService
	MarketNews() marketNewsService
	AccountInforamtion() accountInforamtionService
	FundTransfers() fundTransfersService

	AuthService() authService
}

type Service struct {
	userService               userService
	stocksService             stocksService
	marketUpdatesService      marketUpdatesService
	ordersService             ordersService
	orderBookService          orderBookService
	portfoliosService         portfoliosService
	tradeConfirmationsService tradeConfirmationsService
	marketNewsService         marketNewsService
	accountInforamtionService accountInforamtionService
	fundTransfersService      fundTransfersService
	authService               authService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}

	services.userService = NewUserService(storage, log)
	services.stocksService = NewStocksService(storage, log)
	services.marketUpdatesService = NewmarketUpdatesService(storage, log)
	services.ordersService = NewordersService(storage, log)
	services.orderBookService = NeworderBookService(storage, log)
	services.portfoliosService = NewportfoliosService(storage, log)
	services.tradeConfirmationsService = NewtradeConfirmationsService(storage, log)
	services.marketNewsService = NewmarketNewsService(storage, log)
	services.accountInforamtionService = NewaccountInforamtionServiceService(storage, log)
	services.fundTransfersService = NewfundTransfersService(storage, log)
	services.authService = NewAuthService(storage, log)

	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Stocks() stocksService {
	return s.stocksService
}

func (s Service) MarketUpdates() marketUpdatesService {
	return s.marketUpdatesService
}

func (s Service) Order() ordersService {
	return s.ordersService
}

func (s Service) OrderBook() orderBookService {
	return s.orderBookService
}

func (s Service) Portfolio() portfoliosService {
	return s.portfoliosService
}

func (s Service) TradeConfirmations() tradeConfirmationsService {
	return s.tradeConfirmationsService
}

func (s Service) MarketNews() marketNewsService {
	return s.marketNewsService
}

func (s Service) AccountInformation() accountInforamtionService {
	return s.accountInforamtionService
}


func (s Service) FundTransfer() fundTransfersService {
	return s.fundTransfersService
}

func (s Service) AuthService() authService {
	return s.authService
}
