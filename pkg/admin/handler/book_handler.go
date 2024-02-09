package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	dom "github.com/anjush-bhargavan/golib_api_gateway/pkg/DOM"
	adminpb "github.com/anjush-bhargavan/golib_api_gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func CreateBookHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var book dom.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.UserCreateBook(ctxt, &adminpb.ABookModel{
		BookName:     book.BookName,
		Author:       book.Author,
		NoOfCopies:   uint32(book.NoOfCopies),
		Year:         book.Year,
		Publications: book.Publications,
		Category:     book.Category,
		Description:  book.Description,
	})
	if err != nil {
		log.Printf("error creating book %v err: %v", book.BookName, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v created successfully", book.BookName),
		"data":    response,
	})
}


func FindAllBooksHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	response,err := client.USerFetchAllBooks(ctxt,&adminpb.AdminNoParam{})
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

func FindBookHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")
	response := &adminpb.ABookModel{}
	var err error
	if id == "" && name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid query",
		})
		return
	}else if id != "" {
		bookID,err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "invalid id",
			})
			return
		}
		response,err = client.UserFetchBookByID(ctxt,&adminpb.ABookID{Id: uint32(bookID)})
		if err != nil {
			log.Printf("error finding  book by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	}else if name != "" {
		response,err = client.UserFetchBookByName(ctxt,&adminpb.ABookName{Name: name})
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