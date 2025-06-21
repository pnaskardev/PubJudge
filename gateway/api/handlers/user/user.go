package user_handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/pnaskardev/pubjudge/gateway/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

func Login(service user.Service) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var input entities.LoginInput
		if err := c.BodyParser(&input); err != nil {
			return presenter.BadRequest(c, "Invalid Payload")
		}
		fmt.Printf("Username: %s, Password: %s\n", input.Username, input.Password)

		pass := input.Password

		fmt.Println(pass)

		// if username != "ender" || pass != "ender" {
		// 	return presenter.BadRequest(c, "Invalid Payload")
		// }

		// need to check in the db if the user is there or not
		user, err := service.FetchUser(&input)
		if user == nil || err != nil {
			return presenter.LoginError(c, err)

		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Duration(time.Hour * 72)).Unix()
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return presenter.LoginSuccess(c, user, t)
	}

}

// AddUser is handler/controller which creates Books in the BookShop
func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.RegisterInput
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		// if requestBody.Author == "" || requestBody.Title == "" {
		// 	c.Status(http.StatusInternalServerError)
		// 	return c.JSON(presenter.UserErrorResponse(errors.New(
		// 		"Please specify title and author")))
		// }

		validate := validator.New()
		if err := validate.Struct(&requestBody); err != nil {
			return presenter.BadRequest(c, "Please provide all fields")
		}

		// creating the user model
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
		user := entities.User{
			Firstname: requestBody.Firstname,
			Lastname:  requestBody.Lastname,
			Username:  requestBody.Username,
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		result, err := service.InsertUser(&user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func UpdateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		result, err := service.UpdateUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		userID := requestBody.ID
		err = service.DeleteUsers(userID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchUsers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersSuccessResponse(fetched))
	}
}
