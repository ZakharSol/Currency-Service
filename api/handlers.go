package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Rate struct {
	Cur_ID       int
	Abbreviation string
	Scale        int
	Name         string
	OfficialRate float64
	Date         string
}

func GetAllRates(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	rows, err := database.Query("SELECT cur_id, abbreviation, scale, name, official_rate, date FROM rates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	rates := []Rate{}
	for rows.Next() {
		var rate Rate
		if err := rows.Scan(&rate.Cur_ID, &rate.Abbreviation, &rate.Scale, &rate.Name, &rate.OfficialRate, &rate.Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rates = append(rates, rate)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}

func GetRateByDate(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	date := r.URL.Query().Get("date")
	if date == "" {
		http.Error(w, "Date is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Requested date:", date)

	rows, err := database.Query("SELECT cur_id, abbreviation, scale, name, official_rate FROM rates WHERE date = ?", date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rates []Rate
	for rows.Next() {
		var rate Rate
		if err := rows.Scan(&rate.Cur_ID, &rate.Abbreviation, &rate.Scale, &rate.Name, &rate.OfficialRate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rates = append(rates, rate)
	}

	if len(rates) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "No data available for the requested date"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}
