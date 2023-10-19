package main

import (
	"log"

	"github.com/iki-rumondor/project2-grup9/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/project2-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project2-grup9/internal/application"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
	"github.com/iki-rumondor/project2-grup9/internal/routes"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	repo := repository.NewRepository(gormDB)
	service := application.NewService(repo)
	handler := customHTTP.NewHandler(service)

	var PORT = ":8080"
	routes.StartServer(handler).Run(PORT)
}
