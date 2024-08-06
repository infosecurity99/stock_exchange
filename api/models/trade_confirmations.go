package models

type TradeConfirmation struct {
    ConfirmationID int     `json:"confirmation_id"`
    OrderID        int     `json:"order_id"`
    UserID         int     `json:"user_id"`
    StockID        int     `json:"stock_id"`
    Quantity       int     `json:"quantity"`
    Price          float64 `json:"price"`
    Timestamp      string  `json:"timestamp"`
}

type CreateTradeConfirmation struct {
    OrderID   int     `json:"order_id"`
    UserID    int     `json:"user_id"`
    StockID   int     `json:"stock_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type UpdateTradeConfirmation struct {
    ConfirmationID int     `json:"-"`
    OrderID        int     `json:"order_id"`
    UserID         int     `json:"user_id"`
    StockID        int     `json:"stock_id"`
    Quantity       int     `json:"quantity"`
    Price          float64 `json:"price"`
}

type TradeConfirmationResponse struct {
    Confirmations []TradeConfirmation `json:"confirmations"`
    Count         int                 `json:"count"`
}
