package user_handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/pnaskardev/pubjudge/gateway/pkg/user"
)

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
