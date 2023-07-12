package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-api-consuming/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/posts", controllers.Index)
	r.HandleFunc("/post/create", controllers.Create)
	r.HandleFunc("/post/store", controllers.Store)
	r.HandleFunc("/post/delete", controllers.Delete)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
