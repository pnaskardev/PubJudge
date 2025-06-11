package user

import (
	user_handlers "github.com/pnaskardev/pubjudge/gateway/api/handlers/user"
	"github.com/pnaskardev/pubjudge/gateway/pkg/user"
	"github.com/pnaskardev/pubjudge/gateway/types/router_types"
)

type UserRoutes struct {
	Router *router_types.Router
}

func NewUserRoutes(router *router_types.Router) *UserRoutes {
	return &UserRoutes{Router: router}
}

func (r *UserRoutes) Register() {

	userCollection := r.Router.Deps.Db.Database.Collection("users")
	userRepo := user.NewRepo(userCollection)
	// first create User Service
	userService := user.NewService(userRepo)

	userRouteGroup := r.Router.Api.Group("/user")
	userRouteGroup.Get("/get", user_handlers.GetUsers(userService))
}
