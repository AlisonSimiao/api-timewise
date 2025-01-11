package user

import (
	"time-wise/api"
	"time-wise/middleware"

	"github.com/gin-gonic/gin"
)

var uv = NewUserValidate()

func CreateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("", "GET", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	api.CreateRoute("api/login", "POST", middleware.Validator(uv.Login), uc.Login)
	api.CreateRoute("api/signup", "POST", middleware.Validator(uv.Create), uc.Create)
}

func CreatePrivateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("api/user", "GET", uc.FindOne)
	api.CreateRoute("api/user", "PATCH", middleware.Validator(uv.Update), uc.Update)
}
