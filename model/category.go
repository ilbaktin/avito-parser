package model

type Category struct {
	Id				string			`json:"id,omitempty"`
	Name			string			`json:"name"`
	Count			*int		`json:"count,omitempty"`
	Subcategories	[]*Category		`json:"subcategories,omitempty"`
}

func NewCategory() *Category {
	return &Category{
		Subcategories: make([]*Category, 0, 10),
	}
}
