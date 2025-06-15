package auth

import (
	user_handlers "github.com/pnaskardev/pubjudge/gateway/api/handlers/user"
	"github.com/pnaskardev/pubjudge/gateway/pkg/user"
	"github.com/pnaskardev/pubjudge/gateway/types/router_types"
)

type AuthRoutes struct {
	Router *router_types.Router
}

func NewAuthRoutes(router *router_types.Router) *AuthRoutes {
	return &AuthRoutes{Router: router}
}

func (r *AuthRoutes) Register() {

	userCollection := r.Router.Deps.Db.Database.Collection("users")
	userRepo := user.NewRepo(userCollection)
	// first create User Service
	authService := user.NewService(userRepo)

	userRouteGroup := r.Router.Api.Group("/auth")
	userRouteGroup.Post("/login", user_handlers.Login(authService))
}
