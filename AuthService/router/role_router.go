package router

import (
	"AuthService/controllers"
	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	RoleController *controllers.RoleController
}

func NewRoleRouter(_userController *controllers.RoleController) Router {
	return &RoleRouter{
		RoleController: _userController,
	}
}

func (rr *RoleRouter) Register(r chi.Router) {
	r.Get("/roles/{id}", rr.RoleController.GetRoleById)
	r.Get("/roles", rr.RoleController.GetAllRoles)
}
