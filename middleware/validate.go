package middleware

import (
	message "time-wise/utils"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

var Messages = govalidator.MapData{
	"email": []string{
		"required:" + message.Required("email"),
		"email:" + message.Email(),
	},
	"password": []string{
		"required:" + message.Required("password"),
		"min" + message.MinLength("password", "6"),
		"max:" + message.MaxLength("password", "20"),
		"alpha_num:" + message.AlphaNum("password"),
	},
	"name": []string{
		"required:" + message.Required("name"),
		"min" + message.MinLength("name", "3"),
		"max:" + message.MaxLength("name", "20"),
		"alpha:" + message.Alpha("name"),
	},
}

func Validator(rules govalidator.MapData) func(c *gin.Context) {

	return func(c *gin.Context) {
		body := make(map[string]interface{}, 0)

		opcs := govalidator.Options{
			Request:  c.Request,
			Data:     &body,
			Rules:    rules,
			Messages: Messages,
		}

		validate := govalidator.New(opcs)

		e := validate.ValidateJSON()
		if len(e) > 0 {
			var errors []string

			for _, err := range e {
				errors = append(errors, err...)
			}

			c.AbortWithStatusJSON(422, errors)
			return
		}

		c.Set("body", body)
		c.Next()
	}
}
