package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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

	id := r.FormValue("post_id")
	idInt, _ := strconv.ParseInt(id, 10, 64)

	newpost := Post{
		Id:     idInt,
		Title:  r.FormValue("title"),
		Body:   r.FormValue("body"),
		UserId: 1,
	}

	jsonValue, err := json.Marshal(newpost)
	if err != nil {
		logger.Println(err)
	}

	buffer := bytes.NewBuffer(jsonValue)

	var req *http.Request

	if id != "" {
		// Update Data
		fmt.Println("Update Process")
		req, err = http.NewRequest(http.MethodPut, basedOnURL+"/posts/"+id, buffer)
		if err != nil {
			logger.Println(err)
		}
	} else {
		// create data
		fmt.Println("Create Process")
		req, err = http.NewRequest(http.MethodPost, basedOnURL+"/posts", buffer)
		if err != nil {
			logger.Println(err)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	// must pointer
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
	if res.StatusCode == 201 || res.StatusCode == 200 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	req, err := http.NewRequest(http.MethodDelete, basedOnURL+"/posts/"+id, nil)
	if err != nil {
		logger.Println(err)
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		logger.Println(err)
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}
}
