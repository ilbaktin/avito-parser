package model

type CategoryView struct {
	Region		*Region		`json:"region"`
	Categories	[]*Category	`json:"categories"`
}
