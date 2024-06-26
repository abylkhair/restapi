package dto

import "time"

type CryptoRequest struct {
	Coin []string `json:"coin"`
}

// dto шаблон проектирования почитать

type Coin struct {
	Id               string    `json:"id"`     //old as world
	Title            string    `json:"title"`  //tenge, ruble, dollar
	Symbol           string    `json:"symbol"` //KZT,RUB,USD
	Price            float64   `json:"price"`
	MaxChangePercent float64   `json:"maxChangePercent"`
	MinChangePercent float64   `json:"minChangePercent"`
	CreatedAt        time.Time `json:"createdAt" `
	UpdatedAt        time.Time `json:"updatedAt"` // dto шаблон
}

func NewCoin(id, title, symbol string, price float64) *Coin {
	return &Coin{
		Id:               id,
		Title:            title,
		Symbol:           symbol,
		Price:            price,
		MaxChangePercent: 0,
		MinChangePercent: 0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}
