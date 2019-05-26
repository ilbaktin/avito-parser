package model

type Product struct {
	Name		string		`json:"name"`
	Price		int			`json:"price"`
	Currency	string		`json:"currency"`
	Category	string		`json:"category"`
	Location	string		`json:"location"`
	Date		string		`json:"date"`
}
