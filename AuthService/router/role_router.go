package router

import (
	"AuthService/controllers"
	"AuthService/middlewares"
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
	r.With(middlewares.CreateRoleRequestValidator).Post("/roles", rr.RoleController.CreateRole)
	r.With(middlewares.UpdateRoleRequestValidator).Put("/roles/{id}", rr.RoleController.UpdateRole)
	r.Delete("/roles/{id}", rr.RoleController.DeleteRole)
	r.Get("/roles/{id}/permissions", rr.RoleController.GetRolePermissions)
	r.With(middlewares.AssignPermissionRequestValidator).Post("/roles/{id}/permissions", rr.RoleController.AssignPermissionToRole)
	r.With(middlewares.RemovePermissionRequestValidator).Delete("/roles/{id}/permissions/{permissionId}", rr.RoleController.RemovePermissionFromRole)
	r.Get("/roles/permissions", rr.RoleController.GetAllRolePermissions)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/roles/{userId}/assign/{roleId}", rr.RoleController.AssignRoleToUser)
}
