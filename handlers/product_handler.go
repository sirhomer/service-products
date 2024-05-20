package handlers

import (
    "net/http"
	"github.com/gorilla/mux"
	 "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
    "encoding/json"
    "io/ioutil"
)
type ProductHandler struct {
	    db *gorm.DB
}
func NewProductHandler(db *gorm.DB) *ProductHandler {
	    return &ProductHandler{
	        db: db,
	    }
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	
}

												
			