package models

type MarketNews struct {
    NewsID     int     `json:"news_id"`
    Headline   string  `json:"headline"`
    Content    string  `json:"content"`
    Timestamp  string  `json:"timestamp"`
}

type CreateMarketNews struct {
    Headline  string `json:"headline"`
    Content   string `json:"content"`
}

type UpdateMarketNews struct {
    NewsID    int    `json:"-"`
    Headline  string `json:"headline"`
    Content   string `json:"content"`
}

type MarketNewsResponse struct {
    NewsItems []MarketNews `json:"news_items"`
    Count     int          `json:"count"`
}
