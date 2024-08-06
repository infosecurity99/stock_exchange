package models

type FundTransfer struct {
    TransferID   int     `json:"transfer_id"`
    FromUserID   int     `json:"from_user_id"`
    ToUserID     int     `json:"to_user_id"`
    Amount       float64 `json:"amount"`
    Timestamp    string  `json:"timestamp"`
}

type CreateFundTransfer struct {
    FromUserID int     `json:"from_user_id"`
    ToUserID   int     `json:"to_user_id"`
    Amount     float64 `json:"amount"`
}

type FundTransferResponse struct {
    Transfers []FundTransfer `json:"transfers"`
    Count     int            `json:"count"`
}
