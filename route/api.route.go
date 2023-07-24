package routes

import (
	"time-wise/api"
	"time-wise/middleware"
	"time-wise/resource/user"
)

func Routes() {
	//api.New().CreateRoute("/swagger/*any", "GET", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.CreateRoutes()

	api.New().GetInstance().Use(middleware.Auth)
	user.CreatePrivateRoutes()
}
