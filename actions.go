package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type PostData struct {
	Target string `json:"target"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var links []Link
	json.NewEncoder(w).Encode(db.Find(&links).Value)
}

func store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	var data PostData
	json.NewDecoder(r.Body).Decode(&data)

	if data.Target == "" {
		w.WriteHeader(400)
		enc.Encode(map[string]string{"error": "Target attribute missing"})
		return
	}

	link := Link{
		Slug:   uniqueSlug(slugLen),
		Target: data.Target,
	}
	db.Create(&link)
	enc.Encode(link)
}

func resolve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var link Link
	db.Where("slug = ?", params["slug"]).First(&link)
	log.Println(link.Slug, link.Target)
	http.Redirect(w, r, link.Target, 302)
}