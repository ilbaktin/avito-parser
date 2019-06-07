package model

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Link     string `json:"link"`
	FullText string `json:"full_text"`
	Currency string `json:"currency"`
	Category string `json:"category"`
	Location string `json:"location,omitempty"`
	Date     string `json:"date"`
}
