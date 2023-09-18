package models

type Product struct {
    ID       uint   `gorm:"primary_key" json:"id"`
	    Name     string `json:"name"`
		    Category string `json:"category"`
			}
			