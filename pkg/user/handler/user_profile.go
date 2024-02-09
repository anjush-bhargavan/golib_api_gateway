package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	dom "github.com/anjush-bhargavan/golib_api_gateway/pkg/DOM"
	userpb "github.com/anjush-bhargavan/golib_api_gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func ViewProfileHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Printf("error getting id from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "id not in context",
		})
		return
	}
	userID,ok := id.(uint64)
	if !ok {
		log.Printf("error parsing id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "id parsing issue ",
		})
		return
	}

	response, err := client.UserProfile(ctx, &userpb.UserID{Id: uint32(userID)})
	if err != nil {
		log.Printf("error getting profile of user %v err: %v", userID, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v profile fetched successfully", userID),
		"data":    response,
	})
}

func UpdateProfileHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Printf("error getting id from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "id not in context",
		})
		return
	}
	userID,ok := id.(uint64)
	if !ok {
		log.Printf("error parsing id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "id parsing issue ",
		})
		return
	}

	var user dom.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response,err := client.UserProfileUpdate(ctx,&userpb.ProfileUpdate{
		Userid: uint32(userID),
		Firstname: user.FirstName,
		Lastname: user.LastName,
		Username: user.UserName,
		Dob: user.DoB,
		Gender: user.Gender,
		Phone: user.Phone,
		Address: user.Address,
	})
	if err != nil {
		log.Printf("error updating user %v err: %v", user.UserName, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v user updated successfully", user.UserName),
		"data":    response,
	})


}
