package user

import (
	"fmt"
	"time"
	"time-wise/repository"
	rest_error "time-wise/restError"
	"time-wise/token"

	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur *repository.Repository
}

func NewUserService() *UserService {
	return &UserService{
		ur: NewUserRepository(),
	}
}

func hashPassword(password string) string {
	bcrypt, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bcrypt)
}

func (u *UserService) update(id int, body map[string]interface{}) (any, *rest_error.Err) {
	var user UserResponse

	fmt.Println(id, body)
	u.ur.FindOne("id = @id", map[string]interface{}{"id": id}, &user)
	if user.Id == 0 {
		return UserResponse{}, rest_error.NewNotFoundError("usuario não encontrado")
	}

	if body["email"] != nil {
		var user UserResponse
		u.ur.FindOne("email = @email and id != @id", map[string]interface{}{"email": body["email"].(string), "id": id}, &user)
		if user.Id != 0 {
			return UserResponse{}, rest_error.NewConflictError("Já existe um usuario com esse email")
		}
	}

	if body["password"] != nil {
		body["Password"] = hashPassword(body["password"].(string))
	}

	u.ur.Update("id = @id", map[string]interface{}{"id": id}, body)
	return body, nil
}

func (u *UserService) create(body User) (UserResponse, *rest_error.Err) {
	var user UserResponse

	u.ur.FindOne("email = @email", map[string]interface{}{"email": body.Email}, &user)
	if user.Id != 0 {
		return UserResponse{}, rest_error.NewConflictError("Já existe um usuario com esse email")
	}

	body.Password = hashPassword(body.Password)

	u.ur.Create(&body)
	if body.Id == 0 {
		return UserResponse{}, rest_error.NewInternalError()
	}
	return UserResponse{
		Id:    body.Id,
		Name:  body.Name,
		Email: body.Email,
	}, nil
}

func (u *UserService) login(body UserLogin) (LoginResponse, *rest_error.Err) {
	values := make(map[string]interface{})
	var user User

	values["email"] = body.Email
	u.ur.FindOne("email = @email", values, &user)
	if user.Email == "" {
		return LoginResponse{}, rest_error.NewUnauthorizedError("email ou senha incorretos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return LoginResponse{}, rest_error.NewUnauthorizedError("email ou senha incorretos")
	}

	t, errToken := token.CreateToken(user.Id, time.Hour*24)
	if errToken != nil {
		return LoginResponse{}, rest_error.NewInternalError()
	}

	return LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: t,
	}, nil
}

func (u *UserService) findOne(id int) (UserResponse, *rest_error.Err) {
	var user UserResponse

	u.ur.FindOne("id = @id", map[string]interface{}{"id": id}, &user)

	if user.Id == 0 {
		return UserResponse{}, rest_error.NewNotFoundError("usuario não encontrado")
	}

	return user, nil
}
