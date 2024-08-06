package models

type AccountInformation struct {
    AccountID  int     `json:"account_id"`
    UserID     int     `json:"user_id"`
    Balance    float64 `json:"balance"`
    CreatedAt  string  `json:"created_at"`
}

type CreateAccountInformation struct {
    UserID  int `json:"user_id"`
    Balance float64 `json:"balance"`
}

type UpdateAccountInformation struct {
    AccountID int     `json:"-"`
    Balance   float64 `json:"balance"`
}

type AccountInformationResponse struct {
    Accounts []AccountInformation `json:"accounts"`
    Count    int                  `json:"count"`
}
