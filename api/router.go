package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/rates", func(w http.ResponseWriter, r *http.Request) {
		GetAllRates(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/rates/date", func(w http.ResponseWriter, r *http.Request) {
		GetRateByDate(w, r, db)
	}).Methods("GET")

	return router
}
