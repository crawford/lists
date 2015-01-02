package api

import (
	"net/http"

	"github.com/crawford/lists/models"

	"github.com/crawford/nap"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func listGet(req *http.Request, db *gorm.DB) (interface{}, nap.Status) {
	list := models.List{}
	id := mux.Vars(req)["ListId"]
	if db.First(&list, id).RecordNotFound() {
		return nil, nap.NotFound{}
	}
	db.Model(&list).Related(&list.Items)
	return list, nap.OK{}
}
