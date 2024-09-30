package db

import (
	"currency_service/models"
	"currency_service/utils"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/currencyDB")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func FetchAndSaveRates(db *sql.DB) error {
	rates, err := utils.FetchRates()
	if err != nil {
		return err
	}

	today := rates[0].Date[:10]

	existingRates := make(map[int]struct{})

	rows, err := db.Query("SELECT cur_id FROM rates WHERE date = ?", today)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var curID int
		if err := rows.Scan(&curID); err != nil {
			return err
		}
		existingRates[curID] = struct{}{}
	}

	uniqueRates := make(map[int]models.Rate)

	for _, rate := range rates {
		curID := rate.Cur_ID

		if _, exists := existingRates[curID]; exists {
			continue
		}

		uniqueRates[curID] = rate
	}

	_, err = db.Exec(`CREATE TEMPORARY TABLE temp_rates AS
                    SELECT MIN(id) AS id 
                    FROM rates 
                    WHERE date = ?
                    GROUP BY cur_id, date`, today)

	if err != nil {
		log.Println("Error creating temp table:", err)
		return err
	}

	_, err = db.Exec(`
        DELETE FROM rates 
        WHERE id NOT IN (SELECT id FROM temp_rates) 
        AND date = ?`, today)

	if err != nil {
		log.Println("Error deleting duplicates:", err)
	}

	_, err = db.Exec("DROP TEMPORARY TABLE IF EXISTS temp_rates")
	if err != nil {
		log.Println("Error dropping temp table:", err)
	}

	for _, rate := range uniqueRates {
		_, err := db.Exec(
			"INSERT INTO rates (cur_id, abbreviation, scale, name, official_rate, date) VALUES (?, ?, ?, ?, ?, ?)",
			rate.Cur_ID, rate.Cur_Abbreviation, rate.Cur_Scale, rate.Cur_Name, rate.Cur_OfficialRate, today,
		)
		if err != nil {
			log.Println("Error inserting new rate for currency:", rate.Cur_ID, err)
		} else {
			log.Println("Inserted new rate for currency:", rate.Cur_ID)
		}
	}

	return nil
}
