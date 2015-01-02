package api

import (
	"net/http"

	"github.com/crawford/lists/models"

	"github.com/crawford/nap"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func userGet(req *http.Request, db *gorm.DB) (interface{}, nap.Status) {
	user := models.User{}
	id := mux.Vars(req)["UserId"]
	if db.First(&user, id).RecordNotFound() {
		return nil, nap.NotFound{}
	}
	return user, nap.OK{}
}
