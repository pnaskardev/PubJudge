package submit

import (
	submission_handlers "github.com/pnaskardev/pubjudge/gateway/api/handlers/submit"
	"github.com/pnaskardev/pubjudge/gateway/pkg/submit"
	"github.com/pnaskardev/pubjudge/gateway/types/router_types"
)

type SubmitRoutes struct {
	Router *router_types.Router
}

func NewSubmitRoutes(router *router_types.Router) *SubmitRoutes {
	return &SubmitRoutes{Router: router}
}

func (r *SubmitRoutes) Register() {

	submitCollection := r.Router.Deps.Db.Database.Collection("submissions")

	submitRepo := submit.NewRepo(submitCollection)

	submitService := submit.NewService(submitRepo)

	submitRouteGroup := r.Router.Api.Group("/submit")

	submitRouteGroup.Post("/", submission_handlers.HandleSubmit(submitService))

}
