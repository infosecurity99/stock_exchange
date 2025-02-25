package models

type PrimaryKey struct {
	ID string `json:"id"`
}

type GetListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	UserID   string `json:"user_id"`
}
