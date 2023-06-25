package controller

import (
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo/v4"
	"github.com/dgrijalva/jwt-go"
	
	"echo-clean-architecture/lib"
	"echo-clean-architecture/domain"
	"echo-clean-architecture/usecase"
	"echo-clean-architecture/api/database"
	"echo-clean-architecture/api/response"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Register
// @Summary     ユーザー登録
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       username body string true "ユーザー名"
// @Param       password body string true "パスワード"
// @Param       email    body string true "メールアドレス"
// @Success     200  {object} response.UserRegister
// @Failure     500  {array}  response.Error
// @Router      /auth/user_register [post]
func (controller *UserController) Register(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)

	user, err := controller.Interactor.UserByEmail(u.Email)
	if err == nil {
		return c.JSON(500, response.NewError(err))
	}

	userKey, err := lib.GenerateKey()
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	u.UserKey = userKey
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	user, err = controller.Interactor.Add(u)
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	return c.JSON(200, response.ToUserRegister(user))
}

// Login
// @Summary     ユーザーログイン
// @tags        Auth
// @Accept      json
// @Produce     json
// @Param       password body string true "パスワード"
// @Param       email    body string true "メールアドレス"
// @Success     200  {object} response.UserLogin
// @Failure     500  {array}  response.Error
// @Router      /auth/user_login [post]
func (controller *UserController) Login(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	
	email := u.Email
	user, err := controller.Interactor.UserByEmail(email)
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}
	
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	baseToken := jwt.New(jwt.SigningMethodHS256)
	claims := baseToken.Claims.(jwt.MapClaims)
	claims["user_key"] = user.UserKey
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	token, err := baseToken.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	return c.JSON(200, response.ToUserLogin(user, token))
}

// Check
// @Summary     ユーザー取得
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Success     200  {object} response.UserCheck
// @Failure     500  {array}  response.Error
// @Router      /user/user_check [get]
func (controller *UserController) Check(c echo.Context) (err error) {
	baseToken := c.Request().Header.Get("Authorization")
	token, err := jwt.Parse(baseToken[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		
		return []byte("secret"), nil
	})
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(500, response.NewError(fmt.Errorf("Invalid token")))
	}

	userKey := claims["user_key"].(string)
	username := claims["username"].(string)
	email := claims["email"].(string)
  
	return c.JSON(200,  response.ToUserCheck(userKey, username, email))
}

// Check
// @Summary     ユーザー削除
// @tags        User
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @param       Authorization header string true "Authorization"
// @Param       user_key path int true "ユーザーキー"
// @Success     200  {object} response.UserDelete
// @Failure     500  {array}  response.Error
// @Router      /user/user_delete/{userKey} [delete]
func (controller *UserController) Delete(c echo.Context) (err error) {
	userKey := c.Param("userKey")
	user := domain.User{ UserKey: userKey }

	err = controller.Interactor.DeleteByUserKey(user)
	if err != nil {
		return c.JSON(500, response.NewError(err))
	}

	return c.JSON(200, response.ToUserDelete())
}
