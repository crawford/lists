package api

import (
	"net/http"

	"github.com/crawford/lists/models"

	"github.com/crawford/nap"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func itemGet(req *http.Request, db *gorm.DB) (interface{}, nap.Status) {
	item := models.Item{}
	id := mux.Vars(req)["ItemId"]
	if db.First(&item, id).RecordNotFound() {
		return nil, nap.NotFound{}
	}
	return item, nap.OK{}
}
