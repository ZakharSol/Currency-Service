package db

type Rate struct {
	ID           int
	Cur_ID       int
	Abbreviation string
	Scale        int
	Name         string
	OfficialRate float64
	Date         string
}
