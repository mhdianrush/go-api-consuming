package controllers

import (
	"net/http"
)

// field based on JSON Placeholder
type Post struct {
	UserId int64  `json:"userId"`
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Index(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Store(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
