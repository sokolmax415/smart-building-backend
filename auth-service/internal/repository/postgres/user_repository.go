package postgres

import (
	"auth-service/internal/entity"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// PostgresUserRepository provides access to the PostgreSQL database.
type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository creates a new PostresRepository struct.
func NewPostgresRepository(path string) (*PostgresUserRepository, error) {
	db, err := newDBConnection(path)

	if err != nil {
		return nil, err
	}

	return &PostgresUserRepository{db: db}, nil
}

// CreateNewUser creates a new user in the database.
func (rep *PostgresUserRepository) CreateNewUser(ctx context.Context, user entity.User) error {
	_, err := rep.db.ExecContext(ctx,
		`Insert INTO userDB.users(firstname, lastname, login, password_hash, role_id) 
		VALUES ($1, $2, $3, $4, $5)`,
		user.Firstname, user.Lastname, user.Login, user.PasswordHash, user.RoleId)

	if err != nil {
		return entity.ErrCreateUser
	}

	return nil
}

// GetUserByLogin returns a user by their login.
func (rep *PostgresUserRepository) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, firstname, lastname, login, password_hash, role_id, registration_time FROM userDB.users WHERE login = $1"
	err := rep.db.QueryRowContext(ctx, query, login).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Login, &user.PasswordHash, &user.RoleId, &user.RegistrationTime)

	if err == sql.ErrNoRows {
		return entity.User{}, entity.ErrUserNotExists
	}

	if err != nil {
		return entity.User{}, entity.ErrGetUserByLogin
	}

	return user, nil

}

// IsUserExists checks whether a user with  the specified login exists in the databse.
func (rep *PostgresUserRepository) IsUserExists(ctx context.Context, login string) (bool, error) {
	var countRows int
	query := "SELECT COUNT(*) FROM userDB.users WHERE login = $1"
	err := rep.db.QueryRowContext(ctx, query, login).Scan(&countRows)

	if err != nil {
		return false, entity.ErrCheckUserExistence
	}

	return countRows > 0, nil
}

// ChangeRoleByLogin updates the user's role by their login.
func (rep *PostgresUserRepository) ChangeRoleByLogin(ctx context.Context, login string, newRoleId int64) error {
	query := "UPDATE userDB.users SET role_id = $1 WHERE login = $2"
	_, err := rep.db.ExecContext(ctx, query, newRoleId, login)

	if err != nil {
		return entity.ErrChangeRole
	}

	return nil
}

// DeleteUser removes a user from the database by their login
func (rep *PostgresUserRepository) DeleteUser(ctx context.Context, login string) error {
	query := "DELETE FROM userDB.users WHERE login = $1"
	_, err := rep.db.ExecContext(ctx, query, login)

	if err != nil {
		return entity.ErrDeleteUser
	}

	return nil
}

// GetAllUsers returns a list of all users from the database.
func (rep *PostgresUserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, firstname, lastname, login, password_hash, role_id, registration_time FROM userDB.users"

	rows, err := rep.db.QueryContext(ctx, query)

	if err != nil {
		return nil, entity.ErrGetAllUsers
	}

	defer rows.Close()

	users := make([]entity.User, 0)
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Login, &user.PasswordHash, &user.RoleId, &user.RegistrationTime)
		if err != nil {
			return nil, entity.ErrScanUserRow
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, entity.ErrScanUserRow
	}

	return users, nil

}

func (rep *PostgresUserRepository) ChangeUserName(ctx context.Context, login string, firstname string, lastname string) error {
	query := "UPDATE userDB.users SET firstname=$1, lastname=$2 WHERE login=$3"

	_, err := rep.db.ExecContext(ctx, query, firstname, lastname, login)
	if err != nil {
		return entity.ErrChangeName
	}
	return nil
}
