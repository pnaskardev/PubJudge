package user_handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/pnaskardev/pubjudge/gateway/pkg/user"
)

func Login(service user.Service) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var input entities.LoginInput
		if err := c.BodyParser(&input); err != nil {
			return presenter.BadRequest(c, "Invalid Payload")
		}
		fmt.Printf("Username: %s, Password: %s\n", input.Username, input.Password)

		username := input.Username
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
		claims["username"] = username
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
		var requestBody entities.User
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
		result, err := service.InsertUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

// UpdateUser is handler/controller which updates data of Books in the BookShop
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

// RemoveBook is handler/controller which removes Books from the BookShop
func RemoveUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		bookID := requestBody.ID
		err = service.DeleteUsers(bookID)
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

// GetBooks is handler/controller which lists all Books from the BookShop
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
