package models

type MarketUpdate struct {
    UpdateID   int     `json:"update_id"`
    StockID    int     `json:"stock_id"`
    Price      float64 `json:"price"`
    Timestamp  string  `json:"timestamp"`
}

type CreateMarketUpdate struct {
    StockID   int     `json:"stock_id"`
    Price     float64 `json:"price"`
}

type UpdateMarketUpdate struct {
    UpdateID  int     `json:"-"`
    StockID   int     `json:"stock_id"`
    Price     float64 `json:"price"`
}

type MarketUpdateResponse struct {
    MarketUpdates []MarketUpdate `json:"market_updates"`
    Count         int            `json:"count"`
}
