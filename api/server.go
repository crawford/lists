package api

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"log"

	"github.com/crawford/nap"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func NewServer(address string, port int) *http.Server {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", address, port),
		Handler: router,
	}

	nap.PayloadWrapper = responseWrapper{}
	nap.PanicHandler = panicHandler{}
	router.NotFoundHandler = nap.NotFoundHandler

	for _, route := range routes {
		for _, action := range route.actions {
			router.Handle(route.path, apiHandler(action.fn)).Methods(action.method)
		}
		router.Handle(route.path, nap.MethodNotAllowedHandler)
	}

	return server
}

func apiHandler(fn func(*http.Request, *gorm.DB) (interface{}, nap.Status)) http.Handler {
	return nap.HandlerFunc(func(req *http.Request) (interface{}, nap.Status) {
		db, err := gorm.Open("postgres", "user=lists_rw database=lists host=/var/run/postgresql sslmode=disable")
		if err != nil {
			return nil, nap.InternalError{"An internal server error occured."}
		}
		defer db.Close()
		return fn(req, db.Debug())
	})
}

type responseWrapper struct{}

func (w responseWrapper) Wrap(payload interface{}, status nap.Status) (interface{}, int) {
	return map[string]interface{}{
		"result": payload,
		"status": map[string]interface{}{
			"code":    status.Code(),
			"message": status.Message(),
		},
	}, status.Code()
}

type panicHandler struct{}

func (h panicHandler) Handle(e interface{}) {
	log.Printf("%#v", e)
	debug.PrintStack()
}
