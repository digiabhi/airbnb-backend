package services

import (
	db "AuthService/db/repositories"
	"AuthService/models"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
	GetRolePermissions(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
}

type RoleServiceImpl struct {
	roleRepository           db.RoleRepository
	rolePermissionRepository db.RolePermissionRepository
}

func NewRoleService(_roleRepo db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: _roleRepo,
	}
}

func (s *RoleServiceImpl) GetRoleById(id int64) (*models.Role, error) {
	role, err := s.roleRepository.GetRoleByID(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	role, err := s.roleRepository.GetRoleByName(name)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	roles, err := s.roleRepository.GetAllRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RoleServiceImpl) CreateRole(name string, description string) (*models.Role, error) {
	role, err := s.roleRepository.CreateRole(name, description)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleServiceImpl) DeleteRoleById(id int64) error {
	err := s.roleRepository.DeleteRoleByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	role, err := s.roleRepository.UpdateRole(id, name, description)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermission, error) {
	permissions, err := s.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	rolePermission, err := s.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
	if err != nil {
		return nil, err
	}
	return rolePermission, nil
}
