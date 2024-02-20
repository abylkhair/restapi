package entities

import "time"

type CurrencyRate struct {
	Id               string    `json:"id"`     //old as world
	Title            string    `json:"title"`  //tenge, ruble, dollar
	Symbol           string    `json:"symbol"` //KZT,RUB,USD
	Price            float64   `json:"price"`
	MaxChangePercent float64   `json:"maxChangePercent"`
	MinChangePercent float64   `json:"minChangePercent"`
	CreatedAt        time.Time `json:"createdAt" `
	UpdatedAt        time.Time `json:"updatedAt"`
}

func NewCurrencyRate(id, title, symbol string, price float64) *CurrencyRate {
	return &CurrencyRate{
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

// all need data to currencies time and min max price of changing for hour
// we can scrap all data or calculate it later
func calculate() {
	//take data from api and calculate
	//orr take already calculated data
}
