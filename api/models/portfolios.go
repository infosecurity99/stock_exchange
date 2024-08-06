package models

type Portfolio struct {
	ID          string `json:"id"`
	PortfolioID int    `json:"portfolio_id"`
	UserID      int    `json:"user_id"`
	StockID     int    `json:"stock_id"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"created_at"`
}

type CreatePortfolio struct {
	UserID   int `json:"user_id"`
	StockID  int `json:"stock_id"`
	Quantity int `json:"quantity"`
}

type UpdatePortfolio struct {
	PortfolioID int `json:"-"`
	StockID     int `json:"stock_id"`
	Quantity    int `json:"quantity"`
}

type PortfolioResponse struct {
	Portfolios []Portfolio `json:"portfolios"`
	Count      int         `json:"count"`
}
