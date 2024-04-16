package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"restapi/dto"
	"restapi/internal/entities"
	"strings"
)

type CryptoClient struct {
	logger *logrus.Logger
	Client *http.Client
}

func NewCryptoCLient(logger *logrus.Logger, client *http.Client) (*CryptoClient, error) {
	if logger == nil {
		return nil, errors.Wrapf(entities.ErrInvalidParam, "logger init failed: %v", logger)
	}
	if client == nil {
		return nil, errors.Wrapf(entities.ErrInvalidParam, "client init failed: %v", client)
	}
	return &CryptoClient{
		logger: logger,
		Client: client,
	}, nil
}

func (c *CryptoClient) GetCurrency(ctx context.Context, currencies []string) (string, []string, error) {
	endpoint := "api/crypto/{rates}"
	url := fmt.Sprintf("http://cryptoapi:8080%s", endpoint)

	cryptoReq := dto.CryptoRequest{
		Coin: currencies,
	}

	reqBody, err := json.Marshal(cryptoReq)
	if err != nil {
		c.logger.Errorf("Failed to Marshal JSON request body : %v , err")
		return "", nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(reqBody))
	if err != nil {
		c.logger.Errorf("Failed to create HTTP request : %v", err)
		return "", nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		c.logger.Errorf("Failed to send HTTP request : %v", err)
		return "", nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.logger.Errorf("Failed to decode JSON response %v", err)
		return "", nil, err
	}
	var invalidCurrencies []string
	var messages []string

	errorsMap, ok := data["errors"].(map[string]interface{})
	if ok && len(errorsMap) > 0 {
		for currency, errMsg := range errorsMap {
			invalidCurrencies = append(invalidCurrencies, currency)
			message := fmt.Sprintf("Криптовалюта %s не является валидной: %v", currency, errMsg)
			c.logger.Info(message)
			messages = append(messages, message)
		}
	}

	if rates, ok := data["rates"].([]interface{}); ok {
		for _, rate := range rates {
			rateMap := rate.(map[string]interface{})
			title := rateMap["title"].(string)
			price := rateMap["price"].(float64)
			message := fmt.Sprintf("Курс %s составляет %.2f", title, price)
			c.logger.Info(message)
			messages = append(messages, message)
		}
	}
	return strings.Join(messages, "\n"), invalidCurrencies, nil
}
func (c *CryptoClient) GetCurrencyStatistics(ctx context.Context, currencies []string) (string, []string, error) {
	endpoint := "/api/crypto/{stats}"
	url := fmt.Sprintf("http://cryptoapi:8080%s", endpoint)

	// Создаем экземпляр структуры CryptoRequest и заполняем криптовалютами
	cryptoReq := dto.CryptoRequest{
		Coin: currencies,
	}
	// Преобразуем структуру в JSON
	reqBody, err := json.Marshal(cryptoReq)
	if err != nil {
		c.logger.Errorf("Failed to marshal JSON request body: %v", err)
		return "", nil, err
	}

	// Создаем HTTP-запрос с указанием тела в формате JSON
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(reqBody))
	if err != nil {
		c.logger.Errorf("Failed to create HTTP request: %v", err)
		return "", nil, err
	}

	// Устанавливаем заголовок Content-Type для указания формата JSON
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		c.logger.Errorf("Failed to send HTTP request: %v", err)
		return "", nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.logger.Errorf("Failed to decode JSON response: %v", err)
		return "", nil, err
	}

	var invalidCurrencies []string
	var messages []string

	// Обработка ошибок
	errorsMap, ok := data["errors"].(map[string]interface{})
	if ok && len(errorsMap) > 0 {
		for currency, errMsg := range errorsMap {
			invalidCurrencies = append(invalidCurrencies, currency)
			message := fmt.Sprintf("Криптовалюта %s не является валидной: %v", currency, errMsg)
			c.logger.Info(message)
			messages = append(messages, message)
		}
	}

	// Обработка валидных данных
	if stats, ok := data["stats"].([]interface{}); ok {
		for _, stat := range stats {
			statMap := stat.(map[string]interface{})
			title := statMap["title"].(string)
			maxValue := statMap["max_value"].(float64)
			minValue := statMap["min_value"].(float64)
			percentChange := statMap["percent_change"].(string)
			message := fmt.Sprintf("Статистика для %s:\nМаксимальное значение: %.2f\nМинимальное значение: %.2f\nИзменение в процентах: %s", title, maxValue, minValue, percentChange)
			c.logger.Info(message)
			messages = append(messages, message)

		}
	}

	return strings.Join(messages, "\n"), invalidCurrencies, nil
}
