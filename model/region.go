package model

type Region struct {
	Name		string		`json:"name"`
	Url			string		`json:"url"`
	LocId		float64		`json:"lid"`
	Counts		float64		`json:"cnt"`
}
