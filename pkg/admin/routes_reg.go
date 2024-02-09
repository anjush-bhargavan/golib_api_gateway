package admin

import (
	adminpb "github.com/anjush-bhargavan/golib_api_gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
	handler "github.com/anjush-bhargavan/golib_api_gateway/pkg/admin/handler"
)

type AdminRoutes struct {
	client adminpb.AdminServiceClient
}

func (a *AdminRoutes) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client)
}

func (a *AdminRoutes) EditUser(c *gin.Context) {
	handler.EditUserHandler(c, a.client)
}

func (a *AdminRoutes) DeleteUser(c *gin.Context) {
	handler.DeleteUserHandler(c, a.client)
}

func (a *AdminRoutes) FindAllUsers(c *gin.Context) {
	handler.FindAllUsersHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByEmail(c *gin.Context) {
	handler.FindUserByEmailHandler(c, a.client)
}

func (a *AdminRoutes) FindUserByID(c *gin.Context) {
	handler.FindUserByIDHandler(c, a.client)
}


func (a *AdminRoutes) CreateBook(c *gin.Context) {
	handler.CreateBookHandler(c,a.client)
}

func (a *AdminRoutes) FindAllBooks(c *gin.Context) {
	handler.FindAllBooksHandler(c,a.client)
}

func (a *AdminRoutes) FindBook(c *gin.Context) {
	handler.FindBookHandler(c,a.client)
}