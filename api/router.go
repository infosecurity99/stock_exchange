package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "test/api/docs"
	"test/api/handler"
	"test/pkg/logger"
	"test/service"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.New(services, log)

	r := gin.New()

	//r.Use(authenticateMiddleware)
	r.Use(gin.Logger())

	{

		r.POST("/auth/admin/login", h.AdminLogin)

		// user endpoints
		r.POST("/user", h.CreateUser)
		r.GET("/user/:id", h.GetUser)
		r.GET("/users", h.GetUserList)
		r.PUT("/user/:id", h.UpdateUser)
		r.DELETE("/user/:id", h.DeleteUser)
		r.PATCH("/user/:id", h.UpdateUserPassword)

		// basket endpoints
		r.POST("/basket", h.CreateBasket)
		r.GET("/basket/:id", h.GetBasket)
		r.GET("/baskets", h.GetBasketList)
		r.PUT("basket/:id", h.UpdateBasket)
		r.DELETE("basket/:id", h.DeleteBasket)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}

func authenticateMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
	} else {
		c.Next()
	}
}

func traceRequest(c *gin.Context) {
	beforeRequest(c)

	c.Next()

	afterRequest(c)
}

func beforeRequest(c *gin.Context) {
	startTime := time.Now()

	c.Set("start_time", startTime)

	log.Println("start time:", startTime.Format("2006-01-02 15:04:05.0000"), "path:", c.Request.URL.Path)
}

func afterRequest(c *gin.Context) {
	startTime, exists := c.Get("start_time")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time)).Seconds()

	log.Println("end time:", time.Now().Format("2006-01-02 15:04:05.0000"), "duration:", duration, "method:", c.Request.Method)
	fmt.Println()
}
