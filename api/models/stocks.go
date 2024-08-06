package models

type Stock struct {
    StockID      int     `json:"stock_id"`
    Symbol       string  `json:"symbol"`
    CompanyName  string  `json:"company_name"`
    CurrentPrice float64 `json:"current_price"`
    CreatedAt    string  `json:"created_at"`
}

type CreateStock struct {
    Symbol       string  `json:"symbol"`
    CompanyName  string  `json:"company_name"`
    CurrentPrice float64 `json:"current_price"`
}

type UpdateStock struct {
    StockID      int     `json:"-"`
    Symbol       string  `json:"symbol"`
    CompanyName  string  `json:"company_name"`
    CurrentPrice float64 `json:"current_price"`
}

type StockResponse struct {
    Stocks []Stock `json:"stocks"`
    Count  int     `json:"count"`
}
