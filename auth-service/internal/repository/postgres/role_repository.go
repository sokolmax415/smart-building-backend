package postgres

import (
	"auth-service/internal/entity"
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// PostgresRoleRepository provides access to the PostgreSQL database.
type PostgresRoleRepository struct {
	db *sql.DB
}

// NewPostgresRoleRepository creates a new PostresRepository struct.
func NewRoleRepository(path string) (*PostgresRoleRepository, error) {
	log.Printf("Creating Postgres Role Repository with connection string: %s", path)

	db, err := newDBConnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresRoleRepository{db: db}, nil
}

// GetIdByRole returns the Id of the specified role.
func (rep *PostgresRoleRepository) GetIdByRole(ctx context.Context, role string) (int64, error) {
	log.Printf("GetIdByRole: role=%s", role)

	query := "SELECT role_id FROM userDB.roles WHERE role = $1"
	log.Printf("Executing query: %s", query)

	var id int64
	err := rep.db.QueryRowContext(ctx, query, role).Scan(&id)
	if err == sql.ErrNoRows {
		log.Printf("ERROR IN GetIdByRole: %v", err)
		return -1, entity.ErrRoleNotExists
	}

	if err != nil {
		log.Printf("ERROR IN GetIdByRole: %v", err)
		return -1, entity.ErrGetRoleId
	}

	log.Printf("GetIdByRole: executed successfully")
	return id, nil
}

// GetAllRoles returns a list of all roles from the database.
func (rep *PostgresRoleRepository) GetAllRoles(ctx context.Context) ([]entity.Role, error) {
	log.Printf("GetAllRoles: starting")

	query := "SELECT role_id, role, permissions FROM userDB.roles"
	log.Printf("Executing query: %s", query)

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
			log.Printf("ERROR IN GetAllRoles: %v", err)
			return nil, entity.ErrScanRoleRow
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR IN GetAllRoles: %v", err)
		return nil, entity.ErrScanRoleRow
	}

	log.Printf("GetAllRoles: executed successfully")
	return roles, nil
}

// CreateNewRole creates a new role with the specified permissions.
func (rep *PostgresRoleRepository) CreateNewRole(ctx context.Context, newRole string, permission string) error {
	log.Printf("CreateNewRole: newRole=%s, permission=%s", newRole, permission)

	query := "INSERT INTO userDB.roles (role,permissions) VALUES ($1, $2)"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, newRole, permission)
	if err != nil {
		log.Printf("ERROR IN CreateNewRole: %v", err)
		return entity.ErrCreateRole
	}

	log.Printf("CreateNewRole: executed successfully")
	return nil
}

// GetRoleById returns the Role of the specified Id.
func (rep *PostgresRoleRepository) GetRoleById(ctx context.Context, id int64) (string, error) {
	log.Printf("GetRoleById: id=%d", id)

	query := "SELECT role FROM userDB.roles WHERE role_id = $1"
	log.Printf("Executing query: %s", query)

	var role string
	err := rep.db.QueryRowContext(ctx, query, id).Scan(&role)

	if err == sql.ErrNoRows {
		log.Printf("ERROR IN GetRoleById: %v", err)
		return "", entity.ErrRoleNotExists
	}

	if err != nil {
		log.Printf("ERROR IN GetRoleById: %v", err)
		return "", entity.ErrGetRole
	}

	log.Printf("GetRoleById: executed successfully")
	return role, nil
}
