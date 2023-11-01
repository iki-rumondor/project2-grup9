package main

import (
	"log"

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
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// migration(gormDB)

	userRepo := repository.NewUserRepository(gormDB)
	userService := application.NewUserService(userRepo)
	userHandler := customHTTP.NewUserHandler(userService)

	photoRepo := repository.NewPhotoRepository(gormDB)
	photoService := application.NewPhotoService(photoRepo)
	photoHandler := customHTTP.NewPhotoHandler(photoService)
  
  commentRepo := repository.NewRepository(gormDB)
	commentService := application.NewCommentService(commentRepo)
	commentHandler := customHTTP.NewCommentHandler(commentService)

	socialmediaRepo := repository.NewRepository(gormDB)
	socialmediaService := application.NewSocialMediaService(socialmediaRepo)
	socialmediaHandler := customHTTP.NewSocialMediaHandler(socialmediaService)

	handlers := customHTTP.Handlers{
    CommentHandler:     commentHandler,
		SocialMediaHandler: socialmediaHandler,
		UserHandler:  userHandler,
		PhotoHandler: photoHandler,
	}

	utils.NewCustomValidator(gormDB)

	var PORT = ":8080"
	routes.StartServer(&handlers).Run(PORT)
}

func migration(db *gorm.DB) {
	migrate := db.Debug().Migrator()
	migrate.DropTable(domain.User{}, domain.Photo{})
	migrate.CreateTable(domain.User{}, domain.Photo{})
	// db.Create(domain.User{
	// 	Age:       12,
	// 	Email:     "iki@gmail.id",
	// 	Password:  "123456",
	// 	Username:  "ilham",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })
}
