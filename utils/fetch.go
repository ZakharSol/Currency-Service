package utils

import (
	"currency_service/models"
	"encoding/json"
	"io"
	"net/http"
)

func FetchRates() ([]models.Rate, error) {
	resp, err := http.Get("https://api.nbrb.by/exrates/rates?periodicity=0")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rates []models.Rate
	if err := json.Unmarshal(body, &rates); err != nil {
		return nil, err
	}
	return rates, nil
}
