package models

type Category struct {
    ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	

}
		