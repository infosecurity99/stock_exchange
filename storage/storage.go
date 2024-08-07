package storage

import (
	"context"
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Stocks() IStocksStorage
	MarketUpdates() IMarketUpdatesStorage
	Orders() IOrdersStorage
	OrderBook() IOrderBookStorage
	Porfolios() IPortfoliosStorage
	TradeConfirmations() ITradeConfirmationsStorage
	MarketNews() IMarketNewsStorage
	AccountInforamtion() IAccountInforamtionStorage
	FundTransfers() IFundTransfersStorage
}
//users 1
type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	GetList(context.Context, models.GetListRequest) (models.UsersResponse, error)
	Update(context.Context, models.UpdateUser) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	GetPassword(context.Context, string) (string, error)
	UpdatePassword(context.Context, models.UpdateUserPassword) error
	GetAdminCredentialsByLogin(context.Context, string) (models.User, error)
}



// stocks2
type IStocksStorage interface {
	Create(context.Context, models.CreateStock) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Stock, error)
	GetList(context.Context, models.GetListRequest) (models.StockResponse, error)
	Update(context.Context, models.UpdateStock) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// markets_updates3
type IMarketUpdatesStorage interface {
	Create(context.Context, models.CreateMarketUpdate) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.MarketUpdate, error)
	GetList(context.Context, models.GetListRequest) (models.MarketUpdateResponse, error)
	Update(context.Context, models.UpdateMarketUpdate) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// orders4
type IOrdersStorage interface {
	Create(context.Context, models.CreateOrder) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Order, error)
	GetList(context.Context, models.GetListRequest) (models.OrderResponse, error)
	Update(context.Context, models.UpdateOrder) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// order_book5
type IOrderBookStorage interface {
	Create(context.Context, models.CreateOrderBook) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.OrderBook, error)
	GetList(context.Context, models.GetListRequest) (models.OrderBookResponse, error)
	Update(context.Context, models.UpdateOrderBook) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// portfolios6
type IPortfoliosStorage interface {
	Create(context.Context, models.CreatePortfolio) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Portfolio, error)
	GetList(context.Context, models.GetListRequest) (models.PortfolioResponse, error)
	Update(context.Context, models.UpdatePortfolio) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// trade_confirmations7
type ITradeConfirmationsStorage interface {
	Create(context.Context, models.CreateTradeConfirmation) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.TradeConfirmation, error)
	GetList(context.Context, models.GetListRequest) (models.TradeConfirmationResponse, error)
	Update(context.Context, models.UpdateTradeConfirmation) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// market_news8
type IMarketNewsStorage interface {
	Create(context.Context, models.CreateMarketNews) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.MarketNews, error)
	GetList(context.Context, models.GetListRequest) (models.MarketNewsResponse, error)
	Update(context.Context, models.UpdateMarketNews) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

// accountinformation9
type IAccountInforamtionStorage interface {
	Create(context.Context, models.CreateAccountInformation) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.AccountInformation, error)
	GetList(context.Context, models.GetListRequest) (models.AccountInformationResponse, error)
	Update(context.Context, models.UpdateAccountInformation) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

//fund_transffers10
type IFundTransfersStorage interface {
	Create(context.Context, models.CreateFundTransfer) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.FundTransfer, error)
	GetList(context.Context, models.GetListRequest) (models.FundTransferResponse, error)
	Delete(context.Context, models.PrimaryKey) error
}
