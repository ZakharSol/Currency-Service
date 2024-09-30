package main

import (
	"currency_service/api"
	"currency_service/db"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	startScheduler(database)

	if err := db.FetchAndSaveRates(database); err != nil {
		log.Println("Failed to fetch and save rates:", err)
	}

	router := api.SetupRouter(database)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func startScheduler(database *sql.DB) {
	c := cron.New()

	_, err := c.AddFunc("0 3 * * *", func() {
		if err := db.FetchAndSaveRates(database); err != nil {
			log.Println("Failed to fetch and save rates in cron job:", err)
		} else {
			log.Println("Successfully fetched and saved rates.")
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	defer c.Stop()
}
