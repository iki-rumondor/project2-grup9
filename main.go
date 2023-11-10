package main

import (
	"log"
	"os"

	"github.com/iki-rumondor/project2-grup9/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/project2-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project2-grup9/internal/application"
	"github.com/iki-rumondor/project2-grup9/internal/domain"
	"github.com/iki-rumondor/project2-grup9/internal/repository"
	"github.com/iki-rumondor/project2-grup9/internal/routes"
	"github.com/iki-rumondor/project2-grup9/internal/utils"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// go migration(gormDB)

	userRepo := repository.NewUserRepository(gormDB)
	userService := application.NewUserService(userRepo)
	userHandler := customHTTP.NewUserHandler(userService)

	photoRepo := repository.NewPhotoRepository(gormDB)
	photoService := application.NewPhotoService(photoRepo)
	photoHandler := customHTTP.NewPhotoHandler(photoService)

	commentRepo := repository.NewCommentRepository(gormDB)
	commentService := application.NewCommentService(commentRepo)
	commentHandler := customHTTP.NewCommentHandler(commentService)

	socialmediaRepo := repository.NewSocialMediaRepository(gormDB)
	socialmediaService := application.NewSocialMediaService(socialmediaRepo)
	socialmediaHandler := customHTTP.NewSocialMediaHandler(socialmediaService)

	handlers := customHTTP.Handlers{
		CommentHandler:     commentHandler,
		SocialMediaHandler: socialmediaHandler,
		UserHandler:        userHandler,
		PhotoHandler:       photoHandler,
	}

	utils.NewCustomValidator(gormDB)

	var PORT = envPortOr("3000")
	routes.StartServer(&handlers).Run(PORT)
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}

func migration(db *gorm.DB) {
	migrate := db.Debug().Migrator()
	migrate.DropTable(domain.User{})
	migrate.DropTable(domain.Photo{})
	migrate.DropTable(domain.Comment{})
	migrate.DropTable(domain.SocialMedia{})
	migrate.CreateTable(domain.User{})
	migrate.CreateTable(domain.Photo{})
	migrate.CreateTable(domain.Comment{})
	migrate.CreateTable(domain.SocialMedia{})
}
