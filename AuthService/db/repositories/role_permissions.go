package db

import (
	"AuthService/models"
	"database/sql"
)

type RolePermissionRepository interface {
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
}

type RolePermissionRepositoryImpl struct {
	db *sql.DB
}

func NewRolePermissionRepository(_db *sql.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{db: _db}
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermission, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE id = ?"
	row := r.db.QueryRow(query, id)

	rolePermission := &models.RolePermission{}
	err := row.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // No role permission found
		}
		return nil, err // Other error
	}
	return rolePermission, nil
}

func (r *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.RolePermission, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions WHERE role_id = ?"
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err // Error executing query
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
		if err != nil {
			return nil, err // Error scanning row
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	if err = rows.Err(); err != nil {
		return nil, err // Error after iterating through rows
	}

	return rolePermissions, nil
}

func (r *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)"
	result, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return nil, err // Error executing query
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err // Error getting last insert ID
	}

	rolePermission := &models.RolePermission{
		Id:           id,
		RoleId:       roleId,
		PermissionId: permissionId,
	}

	return rolePermission, nil
}

func (r *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
	_, err := r.db.Exec(query, roleId, permissionId)
	if err != nil {
		return err // Error executing query
	}
	return nil
}

func (r *RolePermissionRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	query := "SELECT id, role_id, permission_id, created_at, updated_at FROM role_permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err // Error executing query
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
		if err != nil {
			return nil, err // Error scanning row
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	if err = rows.Err(); err != nil {
		return nil, err // Error after iterating through rows
	}

	return rolePermissions, nil
}
