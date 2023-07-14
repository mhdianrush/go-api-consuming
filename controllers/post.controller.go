package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// field based on JSON Placeholder
type Post struct {
	UserId int64  `json:"userId"`
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var basedOnURL = "https://jsonplaceholder.typicode.com"

func Index(w http.ResponseWriter, r *http.Request) {
	var posts []Post

	response, err := http.Get(basedOnURL + "/posts")
	if err != nil {
		logger.Println(err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&posts)
	if err != nil {
		logger.Println(err)
	}

	data := map[string]any{
		"post": posts,
	}

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		logger.Println(err)
	}
	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/create.html")
	if err != nil {
		logger.Println(err)
	}
	temp.Execute(w, nil)
}

func Store(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
