package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type Rate struct {
	Cur_ID           int
	Date             string
	Cur_Abbreviation string
	Cur_Scale        int
	Cur_Name         string
	Cur_OfficialRate float64
}

func FetchRates() ([]Rate, error) {
	resp, err := http.Get("https://api.nbrb.by/exrates/rates?periodicity=0")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rates []Rate
	if err := json.Unmarshal(body, &rates); err != nil {
		return nil, err
	}
	return rates, nil
}
