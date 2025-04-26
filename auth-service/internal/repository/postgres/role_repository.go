package postgres

import (
	"auth-service/internal/entity"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgresRoleRepository provides access to the PostgreSQL database.
type PostgresRoleRepository struct {
	db *sql.DB
}

// NewPostgresRoleRepository creates a new PostresRepository struct.
func NewRoleRepository(path string) (*PostgresRoleRepository, error) {
	db, err := newDBConnection(path)

	if err != nil {
		return nil, err
	}

	return &PostgresRoleRepository{db: db}, nil
}

// GetIdByRole returns the Id of the specified role.
func (rep *PostgresRoleRepository) GetIdByRole(ctx context.Context, role string) (int64, error) {
	var id int64
	err := rep.db.QueryRowContext(ctx, "SELECT role_id FROM userDB.roles WHERE role = $1", role).Scan(&id)

	if err == sql.ErrNoRows {
		return -1, entity.ErrRoleNotExists
	}

	if err != nil {
		return -1, entity.ErrGetRoleId
	}

	return id, nil
}

// GetAllRoles returns a list of all roles from the database.
func (rep *PostgresRoleRepository) GetAllRoles(ctx context.Context) ([]entity.Role, error) {
	query := "SELECT role_id, role, permissions FROM userDB.roles"
	rows, err := rep.db.QueryContext(ctx, query)

	if err != nil {
		return nil, entity.ErrGetAllRoles
	}
	defer rows.Close()

	roles := make([]entity.Role, 0)
	for rows.Next() {
		var role entity.Role

		err := rows.Scan(&role.RoleId, &role.Role, &role.Permissions)

		if err != nil {
			return nil, entity.ErrScanRoleRow
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, entity.ErrScanRoleRow
	}

	return roles, nil
}

// CreateNewRole creates a new role with the specified permissions.
func (rep *PostgresRoleRepository) CreateNewRole(ctx context.Context, newRole string, permission string) error {
	query := "INSERT INTO userDB.roles (role,permissions) VALUES ($1, $2)"
	_, err := rep.db.ExecContext(ctx, query, newRole, permission)
	if err != nil {
		return entity.ErrCreateRole
	}
	return nil
}

// GetRoleById returns the Role of the specified Id.
func (rep *PostgresRoleRepository) GetRoleById(ctx context.Context, id int64) (string, error) {
	query := "SELECT role FROM userDB.roles WHERE role_id = $1"
	var role string
	err := rep.db.QueryRowContext(ctx, query, id).Scan(&role)

	if err == sql.ErrNoRows {
		return "", entity.ErrRoleNotExists
	}

	if err != nil {
		return "", entity.ErrGetRole
	}
	return role, nil
}
