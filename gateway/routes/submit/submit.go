package submit

import (
	"github.com/pnaskardev/pubjudge/gateway/types/router_types"
)

type SubmissionRoutes struct {
	Router *router_types.Router
}

func NewSubmissionRoutes(router *router_types.Router) *SubmissionRoutes {
	return &SubmissionRoutes{Router: router}
}

func (r *SubmissionRoutes) Register() {
	// group := r.app.Group("/api/submission")
}
