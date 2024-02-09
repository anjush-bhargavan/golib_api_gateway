package user

import (
	"github.com/anjush-bhargavan/golib_api_gateway/pkg/services"
	"github.com/anjush-bhargavan/golib_api_gateway/pkg/user/handler"
	pb "github.com/anjush-bhargavan/golib_api_gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type User struct {
	client pb.UserServiceClient
	Redis  *services.RedisClient
}

func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, u.client, "user")
}

func (u *User) Signup(c *gin.Context) {
	handler.UserSignupHandler(c, u.client)
}

func (u *User) ViewProfile(c *gin.Context) {
	handler.ViewProfileHandler(c, u.client)
}

func (u *User) UpdateProfile(c *gin.Context) {
	handler.UpdateProfileHandler(c, u.client)
}

func (u *User) FindAllBooks(c *gin.Context) {
	handler.FindAllBooksHandler(c,u.client)
}

func (u *User) FindBook(c *gin.Context) {
	handler.FindBookHandler(c,u.client)
}

func (u *User) ShowPlans(c *gin.Context) {
	handler.ShowPlansHandler(c)
}

func (u *User) GetPlan(c *gin.Context) {
	handler.GetPlanHandler(c,u.Redis)
}

func (u *User) Membership(c *gin.Context) {
	handler.MembershipHandler(c,u.client,u.Redis)
}


func (u *User) RazorpaySuccess(c *gin.Context) {
	handler.RazorpaySuccessHandler(c,u.client,u.Redis)
}

func (u *User) SuccessPage(c *gin.Context) {
	handler.SuccessPageHandler(c)
}