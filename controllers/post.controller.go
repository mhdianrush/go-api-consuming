package controllers

import (
	"bytes"
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
	// check id to edit
	var post Post

	id := r.URL.Query().Get("id")
	if id != "" {
		res, err := http.Get(basedOnURL + "/posts/" + id)
		if err != nil {
			logger.Println(err)
		}
		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&post)
		if err != nil {
			logger.Println(err)
		}
	}
	data := map[string]any{
		"post": post,
	}
	temp, err := template.ParseFiles("views/create.html")
	if err != nil {
		logger.Println(err)
	}
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	newpost := Post{
		Id: 0,
		// Id = 0 is just an example, the real value will follow next when Click Submit
		Title:  r.FormValue("title"),
		Body:   r.FormValue("body"),
		UserId: 1,
	}

	jsonValue, err := json.Marshal(newpost)
	if err != nil {
		logger.Println(err)
	}

	buffer := bytes.NewBuffer(jsonValue)
	req, err := http.NewRequest(http.MethodPost, basedOnURL+"/posts", buffer)
	if err != nil {
		logger.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		logger.Println(err)
	}
	defer res.Body.Close()

	var postResponse Post
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&postResponse)
	if err != nil {
		logger.Println(err)
	}
	// fmt.Println(res.StatusCode)
	// fmt.Println(res.Status)
	// fmt.Println(postResponse)
	if res.StatusCode == 201 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

}
