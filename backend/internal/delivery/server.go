package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"react-go-template/docs"
	"react-go-template/internal/delivery/handlers"
	"react-go-template/internal/delivery/middleware"
	"react-go-template/internal/repository/user"
	"react-go-template/internal/service"
	"react-go-template/pkg/log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, logger *log.Logs) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRepo := user.InitUserRepo(db)
	userService := service.InitUserService(userRepo, logger)
	userHandler := handlers.InitUserHandler(userService)

	userRouter := r.Group("/user")

	userRouter.POST("/create", userHandler.Create)
	userRouter.GET("/get/:id", userHandler.GetUser)
	userRouter.DELETE("/delete/:id", userHandler.Delete)
	userRouter.POST("/login", userHandler.Login)

	mdw := middleware.InitMiddleware(logger)
	r.Use(mdw.CORSMiddleware())

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
