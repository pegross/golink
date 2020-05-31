package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const slugLen = 4

var db *gorm.DB

type Link struct {
	gorm.Model
	Slug   string `json:"slug" gorm:"type:varchar(4);unique"`
	Target string `json:"target" gorm:"size:255"`
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var err error
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Link{})
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/", store).Methods("POST")
	router.HandleFunc("/{id}", resolve).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))
}
