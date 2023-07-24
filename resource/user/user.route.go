package user

import (
	"time-wise/api"
	"time-wise/middleware"
)

var uv = NewUserValidate()

func CreateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("/login", "POST", middleware.Validator(uv.Login), uc.Login)
	api.CreateRoute("/signup", "POST", middleware.Validator(uv.Create), uc.Create)
}

func CreatePrivateRoutes() {
	api := api.New()
	uc := NewUserController()

	api.CreateRoute("/user/:id", "GET", uc.FindOne)
	api.CreateRoute("/user/:id", "PATCH", middleware.Validator(uv.Update), uc.Update)
}
