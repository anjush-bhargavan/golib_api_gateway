package user

import (
	"log"

	"github.com/anjush-bhargavan/golib_api_gateway/middleware"
	"github.com/anjush-bhargavan/golib_api_gateway/pkg/services"
	"github.com/gin-gonic/gin"
)

func NewUserRoute(c *gin.Engine) {
	client, err := ClientDial()
	redisClient := services.NewRedisClient()
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v", err.Error())
	}

	userHandler := User{
		client: client,
		Redis:  redisClient,
	}
	apiUser := c.Group("/api/user")
	{
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/signup", userHandler.Signup)
	}

	apiUserAuth := c.Group("/api/user/auth")
	apiUserAuth.Use(middleware.Authorization("user"))
	{
		apiUserAuth.GET("/profile", userHandler.ViewProfile)
		apiUserAuth.PUT("/profile", userHandler.UpdateProfile)

		apiUserAuth.GET("/plans", userHandler.ShowPlans)
		apiUserAuth.POST("/plans", userHandler.GetPlan)
		apiUser.GET("/membership",userHandler.Membership)
		c.GET("/payment/success",userHandler.RazorpaySuccess)
		c.GET("/success", userHandler.SuccessPage)

		apiUserAuth.GET("/books",userHandler.FindAllBooks)
		apiUserAuth.GET("/book",userHandler.FindBook)
	}

	apiUserMemb := c.Group("/api/user/memb")
	apiUserMemb.Use(middleware.Authorization("user"))
	{
		apiUserMemb.POST("/checkout")
	}
}
