package api

import (
	"net/http"

	"github.com/crawford/nap"
	"github.com/jinzhu/gorm"
)

type action struct {
	method string
	fn     func(*http.Request, *gorm.DB) (interface{}, nap.Status)
}

var routes = []struct {
	path   string
	actions []action
}{
	{
		"/user/{UserId}",
		[]action{
			{"GET", userGet},
			//{"PUT", userPut},
			//{"DELETE", userDelete},
		},
	},
	{
		"/user/{UserId}/list/{ListId}",
		[]action{
			{"GET", listGet},
			//{"PUT", listPut},
			//{"DELETE", listDelete},
		},
	},
	{
		"/user/{UserId}/list/{ListId}/item/{ItemId}",
		[]action{
			{"GET", itemGet},
			//{"PUT", itemPut},
			//{"DELETE", itemDelete},
		},
	},
}
