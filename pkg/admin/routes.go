package admin

import (
	"log"

	"github.com/anjush-bhargavan/golib_api_gateway/middleware"
	"github.com/gin-gonic/gin"
)

func NewAdminRoute(c *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	adminHandler := AdminRoutes{
		client: client,
	}

	apiAdmin := c.Group("/api/admin")
	{
		apiAdmin.POST("/login", adminHandler.Login)

	}
	apiAdminAuth := c.Group("/api/admin/auth")
	apiAdminAuth.Use(middleware.Authorization("admin"))
	{
		apiAdminAuth.PUT("/edit/user", adminHandler.EditUser)
		apiAdminAuth.DELETE("delete/user", adminHandler.DeleteUser)
		apiAdminAuth.GET("/users", adminHandler.FindAllUsers)
		apiAdminAuth.GET("/id/user", adminHandler.FindUserByID)
		apiAdminAuth.GET("/email/user", adminHandler.FindUserByEmail)
		apiAdminAuth.POST("/book",adminHandler.CreateBook)
		apiAdminAuth.GET("/books",adminHandler.FindAllBooks)
		apiAdminAuth.GET("/book",adminHandler.FindBook)
	}
}
