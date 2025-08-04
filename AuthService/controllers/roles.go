package controllers

import (
	"AuthService/dto"
	"AuthService/services"
	"AuthService/utils"
	"fmt"
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

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(dto.CreateRoleRequestDTO)

	role, err := rc.RoleService.CreateRole(payload.Name, payload.Description)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to create role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusCreated, "Role created successfully", role)
}

func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
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

	payload := r.Context().Value("payload").(dto.UpdateRoleRequestDTO)

	role, err := rc.RoleService.UpdateRole(id, payload.Name, payload.Description)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to update role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role updated successfully", role)
}

func (rc *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
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

	err = rc.RoleService.DeleteRoleById(id)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role deleted successfully", nil)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
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

	permissions, err := rc.RoleService.GetRolePermissions(id)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role permissions", err)
		return
	}

	if len(permissions) == 0 {
		utils.WriteJSONSuccessResponse(w, http.StatusOK, "No permissions found for this role", nil)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", permissions)
}

func (rc *RoleController) AssignPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleId := chi.URLParam(r, "id")
	if roleId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("Role ID is missing"))
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid Role ID format", err)
		return
	}

	payload := r.Context().Value("payload").(dto.AssignPermissionRequestDTO)

	rolePermission, err := rc.RoleService.AddPermissionToRole(id, payload.PermissionId)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to assign permission to role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission assigned to role successfully", rolePermission)
}

func (rc *RoleController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
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

	payload := r.Context().Value("payload").(dto.RemovePermissionRequestDTO)

	err = rc.RoleService.RemovePermissionFromRole(id, payload.PermissionId)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to remove permission from role", err)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission removed from role successfully", nil)
}

func (rc *RoleController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := rc.RoleService.GetAllRolePermissions()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch all role permissions", err)
		return
	}

	if len(permissions) == 0 {
		utils.WriteJSONSuccessResponse(w, http.StatusOK, "No role permissions found", nil)
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "All role permissions fetched successfully", permissions)
}
