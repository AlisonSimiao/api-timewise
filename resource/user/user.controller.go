package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

var userService = NewUserService()

func (uc *UserController) FindOne(c *gin.Context) {
	userIDString := c.Param("id")

	// Converta o parâmetro para int usando strconv.Atoi()
	idUser, err := strconv.Atoi(userIDString)
	if err != nil {
		// Se a conversão falhar, retorne um erro 400 Bad Request
		c.AbortWithError(400, err)
		return
	}

	user, rest_error := userService.findOne(idUser)
	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Update(c *gin.Context) {	
	body, exist := c.Get("body")
	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	userIDString := c.Param("id")

	idUser, err := strconv.Atoi(userIDString)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	user, rest_error := userService.update(idUser, body.(map[string]interface{}))
	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Create(c *gin.Context) {
	body, exist := c.Get("body")
	if !exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "Erro no servidor"})
		return
	}

	user, rest_error := userService.create(User{
		Name:     body.(map[string]interface{})["name"].(string),
		Email:    body.(map[string]interface{})["email"].(string),
		Password: body.(map[string]interface{})["password"].(string),
	})

	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Login(c *gin.Context) {

	body, exist := c.Get("body")

	if !exist {
		c.JSON(500, "Erro no servidor")
		return
	}

	user, rest_error := userService.login(UserLogin{
		Email:    body.(map[string]interface{})["email"].(string),
		Password: body.(map[string]interface{})["password"].(string),
	})
	if rest_error != nil {
		c.AbortWithStatusJSON(rest_error.GetStatus(), rest_error.JsonError())
		return
	}

	c.JSON(200, user)
}
