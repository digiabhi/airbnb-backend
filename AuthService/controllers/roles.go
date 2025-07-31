package controllers

import (
	"AuthService/services"
	"AuthService/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", nil)
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID format", err)
		return
	}
	role, err := rc.RoleService.GetRoleById(id)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role", err)
		return
	}

	if role == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.RoleService.GetAllRoles()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch roles", err)
		return
	}
	if len(roles) == 0 {
		utils.WriteJSONSuccessResponse(w, http.StatusOK, "No roles found", nil)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
}
