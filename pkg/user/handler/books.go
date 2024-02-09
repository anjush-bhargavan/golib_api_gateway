package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	userpb "github.com/anjush-bhargavan/golib_api_gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func FindAllBooksHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	response, err := client.USerFetchAllBooks(ctxt, &userpb.UNoParam{})
	if err != nil {
		log.Printf("error finding all books  err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all books successfully",
		"data":    response,
	})

}

func FindBookHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")
	response := &userpb.UBookModel{}
	var err error
	if id == "" && name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid query",
		})
		return
	} else if id != "" {
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "invalid id",
			})
			return
		}
		response, err = client.UserFetchBookByID(ctxt, &userpb.UBookID{Id: uint32(bookID)})
		if err != nil {
			log.Printf("error finding  book by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	} else if name != "" {
		response, err = client.UserFetchBookByName(ctxt, &userpb.UBookName{Name: name})
		if err != nil {
			log.Printf("error finding  book by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all books successfully",
		"data":    response,
	})

}
