package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-api-consuming/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/", controllers.Index)
	routes.HandleFunc("/posts", controllers.Index)
	routes.HandleFunc("/post/create", controllers.Create)
	routes.HandleFunc("/post/store", controllers.Store)
	routes.HandleFunc("/post/delete", controllers.Delete)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: routes,
	}
	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}
	logger.Printf("server running on port %s", os.Getenv("SERVER_PORT"))
}
