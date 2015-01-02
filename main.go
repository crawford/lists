package main

import (
	"flag"
	"log"

	"github.com/crawford/lists/models"
	"github.com/crawford/lists/api"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	flags = struct {
		port    int
		address string
	}{}
)

func init() {
	flag.StringVar(&flags.address, "address", "0.0.0.0", "address to listen on")
	flag.IntVar(&flags.port, "port", 80, "port to bind on")
}

func main() {
	flag.Parse()

	db, err := gorm.Open("postgres", "user=lists_rw database=lists host=/var/run/postgresql sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.Debug().CreateTable(&models.Item{})
	db.Debug().CreateTable(&models.List{})
	db.Debug().CreateTable(&models.User{})
	//db.Debug().AutoMigrate(&models.Item{})
	db.Close()

	log.Fatal(api.NewServer(flags.address, flags.port).ListenAndServe())
}
