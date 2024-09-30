package models

type Rate struct {
	Cur_ID           int     `json:"cur_id"`
	Date             string  `json:"date"`
	Cur_Abbreviation string  `json:"cur_abbreviation"`
	Cur_Scale        int     `json:"cur_scale"`
	Cur_Name         string  `json:"cur_name"`
	Cur_OfficialRate float64 `json:"cur_official_rate"`
}
