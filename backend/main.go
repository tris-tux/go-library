package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-library/backend/database"
	"github.com/tris-tux/go-library/backend/handler"
	"github.com/tris-tux/go-library/backend/middleware"
	"github.com/tris-tux/go-library/backend/repository"
	"github.com/tris-tux/go-library/backend/schema"
	"github.com/tris-tux/go-library/backend/service"
)

func main() {

	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&schema.Visitor{})

	visitorRepository := repository.NewVisitorRepository(db)
	jwtService := service.NewJWTService()
	visitorService := service.NewVisitorService(visitorRepository)
	authService := service.NewAuthService(visitorRepository)
	visitorHandler := handler.NewVisitorHandler(visitorService, jwtService)
	authHandler := handler.NewAuthHandler(authService, jwtService)

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	visitorRoutes := r.Group("api/visitor", middleware.AuthorizeJWT(jwtService))
	{
		visitorRoutes.GET("/profile", visitorHandler.Profile)
		// visitorRoutes.PUT("/profile", visitorHandler.Update)
	}
	r.Run()
}
