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

	commentRepo := repository.NewRepository(gormDB)
	commentService := application.NewCommentService(commentRepo)
	commentHandler := customHTTP.NewCommentHandler(commentService)

	socialmediaRepo := repository.NewRepository(gormDB)
	socialmediaService := application.NewSocialMediaService(socialmediaRepo)
	socialmediaHandler := customHTTP.NewSocialMediaHandler(socialmediaService)

	handlers := customHTTP.Handlers{
		CommentHandler:     commentHandler,
		SocialMediaHandler: socialmediaHandler,
	}

	var PORT = ":8080"
	routes.StartServer(&handlers).Run(PORT)

}
