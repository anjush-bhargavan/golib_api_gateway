package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/anjush-bhargavan/golib_api_gateway/pkg/services"
	userpb "github.com/anjush-bhargavan/golib_api_gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/razorpay/razorpay-go"
)

// Subscription holds the detail of plan
type Subscription struct {
	Plan string `json:"plan" validate:"plan"`
}

type pageVariables struct {
	OrderID string
}

// ValidatePlan function helps to validate plan details entered by user
func ValidatePlan(fl validator.FieldLevel) bool {
	plan := fl.Field().String()
	return plan == "3Y" || plan == "1Y" || plan == "5M"
}

// ShowPlansHandler function shows the membership plans
func ShowPlansHandler(c *gin.Context) {
	plans := []gin.H{
		{"plan": "5M", "description": "5 months you can use the platform to borrow books for 500"},
		{"plan": "1Y", "description": "1 year you can use the platform to borrow books for 1000"},
		{"plan": "3Y", "description": "3 years you can use the platform to borrow books for 2000"},
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "User Plans",
		"data":    plans,
	})
}

// GetPlanHandler function gets the plan from users
func GetPlanHandler(c *gin.Context, client *services.RedisClient) {
	var user Subscription
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Error binding",
			"data":    err.Error(),
		})
		return
	}

	validate.RegisterValidation("plan", ValidatePlan)
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please select correct plan",
			"data":    err.Error(),
		})
		return
	}

	userIDContext, _ := c.Get("user_id")
	userID := userIDContext.(uint64)
	userIDString := fmt.Sprint(userID)
	// fmt.Println(userIDString)
	if err := client.SetValue("plan"+userIDString, user.Plan, 30*time.Minute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Error seting data in redis",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Plan selected",
		"data":    user,
	})
}

// Membership handles the membership of users
func MembershipHandler(c *gin.Context, userClient userpb.UserServiceClient, redisClient *services.RedisClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()
	userID := c.Query("id")

	// userIDContext,_:=c.Get("user_id")

	// userID:=userIDContext.(uint64)
	// userIDString:=fmt.Sprint(userID)
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "id not found",
			"data":    err.Error(),
		})
		return
	}

	response, err := userClient.UserProfile(ctx, &userpb.UserID{Id: uint32(id)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "User not found",
			"data":    err.Error(),
		})
		return
	}

	// var existingMember models.Membership
	// if err:=config.DB.Where("user_id = ? AND is_active = ?",user.UserID,true).First(&existingMember).Error; err == nil {
	// 		c.JSON(http.StatusConflict,gin.H{
	// 										"status":"Failed",
	// 										"message":"Already have a membership",
	// 										"data":nil,
	// 									})
	// 		return
	// }else if err != gorm.ErrRecordNotFound {
	// 	c.JSON(http.StatusBadRequest,gin.H{
	// 										"status":"Failed",
	// 										"message":"Databae error",
	// 										"data":err.Error(),
	// 									})
	// 	return
	// }

	userPlan, err := redisClient.GetValue("plan" + userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "Plan not found",
			"data":    err.Error(),
		})
		return
	}
	amount := 0
	if userPlan == "5M" {
		amount = 500
	} else if userPlan == "1Y" {
		amount = 1000
	} else if userPlan == "3Y" {
		amount = 2000
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid plan",
			"data":    err.Error(),
		})
	}
	amountInPaisa := amount * 100

	client := razorpay.NewClient(os.Getenv("RAZORPAY_KEY_ID"), os.Getenv("RAZORPAY_SECRET"))

	data := map[string]interface{}{
		"amount":   amountInPaisa,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Printf("Problem getting repository information: %v\n", err)
		os.Exit(1)
	}

	value := body["id"]
	str := value.(string)

	homepageVariables := pageVariables{
		OrderID: str,
	}

	c.HTML(http.StatusOK, "app.html", gin.H{
		"userID":      userID,
		"totalPrice":  amountInPaisa / 100,
		"total":       amountInPaisa,
		"orderID":     homepageVariables.OrderID,
		"email":       response.Email,
		"phoneNumber": response.Phone,
	})
}

// RazorpaySuccess handles successful RazorPay payments.
func RazorpaySuccessHandler(c *gin.Context, userClient userpb.UserServiceClient, redisClient *services.RedisClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	userIDString := c.Query("user_id")
	UserID, _ := strconv.Atoi(userIDString)
	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	signature := c.Query("signature")
	fmt.Println(signature, orderID)
	userPlan, err := redisClient.GetValue("plan" + userIDString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Failed",
			"message": "OTP Not found",
			"data":    err.Error(),
		})
		return
	}

	_, err = userClient.CreateMembership(ctx, &userpb.Membership{
		UserId:         uint32(UserID),
		SubscriptionId: paymentID,
		Plan:           userPlan,
		IsActive:       true,
		StartedOn:      "",
		ExpiresAt:      "",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "error creating membership",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true})
}

// SuccessPage renders the success page.
func SuccessPageHandler(c *gin.Context) {
	pID := c.Query("id")
	userID := c.Query("user_id")
	fmt.Println(pID)
	fmt.Println("Fully successful")

	c.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": pID,
		"userID":    userID,
	})
}
