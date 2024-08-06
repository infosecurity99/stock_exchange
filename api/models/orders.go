package models

type Order struct {
    OrderID    int     `json:"order_id"`
    UserID     int     `json:"user_id"`
    StockID    int     `json:"stock_id"`
    OrderType  string  `json:"order_type"`
    Quantity   int     `json:"quantity"`
    Price      float64 `json:"price"`
    Status     string  `json:"status"`
    CreatedAt  string  `json:"created_at"`
}

type CreateOrder struct {
    UserID    int     `json:"user_id"`
    StockID   int     `json:"stock_id"`
    OrderType string  `json:"order_type"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type UpdateOrder struct {
    OrderID   int     `json:"-"`
    UserID    int     `json:"user_id"`
    StockID   int     `json:"stock_id"`
    OrderType string  `json:"order_type"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
    Status    string  `json:"status"`
}

type OrderResponse struct {
    Orders []Order `json:"orders"`
    Count  int     `json:"count"`
}
