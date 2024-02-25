package entities

import "time"

type Coin struct {
	ID        int64
	Name      string
	Price     float64
	CreatedAt time.Time
}

func NewCoin(ID int64, name string, price float64, createdAt time.Time) *Coin {
	return &Coin{ID: ID, Name: name, Price: price, CreatedAt: createdAt}
}

// all need data to currencies time and min max price of changing for hour
// we can scrap all data or calculate it later
func calculate() {
	//take data from api and calculate
	//orr take already calculated data
}
получить все монеты
получить одну монету
t->start->rates->api(getRATES)->return(telega)
№№API

GET /rates //вернет все монеты со всей фулл инфой с криптобиржи
GET /rates/{cryptocurrency} -\\- но по одной монете

№№Команды в Telegram боте

/start
/rates
/rates {cryptocurrency}
/start-auto {minutes_count} (пример /start-auto 10, что значит отправлять каждые 10 минут)
/stop-auto