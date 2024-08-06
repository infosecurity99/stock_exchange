package models

type OrderBook struct {
    OrderID   int     `json:"order_id"`
    OrderType string  `json:"order_type"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type CreateOrderBook struct {
    OrderID   int     `json:"order_id"`
    OrderType string  `json:"order_type"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type UpdateOrderBook struct {
    OrderID   int     `json:"-"`
    OrderType string  `json:"order_type"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type OrderBookResponse struct {
    Entries []OrderBook `json:"entries"`
    Count   int              `json:"count"`
}
