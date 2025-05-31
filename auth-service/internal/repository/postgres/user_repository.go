package postgres

import (
	"auth-service/internal/entity"
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// PostgresUserRepository provides access to the PostgreSQL database.
type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository creates a new PostresRepository struct.
func NewPostgresUserRepository(path string) (*PostgresUserRepository, error) {
	log.Printf("Creating Postgres User Repository with connection string: %s", path)

	db, err := newDBConnection(path)
	if err != nil {
		return nil, err
	}

	return &PostgresUserRepository{db: db}, nil
}

// CreateNewUser creates a new user in the database.
func (rep *PostgresUserRepository) CreateNewUser(ctx context.Context, user entity.User) error {
	log.Printf("CreateNewUser: firstName=%s, lastName=%s, login=%s, password_hash=%s, roleId=%d", user.Firstname, user.Lastname, user.Login, user.PasswordHash, user.RoleId)

	query := "Insert INTO userDB.users(firstname, lastname, login, password_hash, role_id) VALUES ($1, $2, $3, $4, $5)"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query,
		user.Firstname, user.Lastname, user.Login, user.PasswordHash, user.RoleId)

	if err != nil {
		log.Printf("ERROR IN CreateNewUser: %v", err)
		return entity.ErrCreateUser
	}

	log.Printf("CreateNewUser: executed successfully")
	return nil
}

// GetUserByLogin returns a user by their login.
func (rep *PostgresUserRepository) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	log.Printf("GetUserByLogin: login=%s", login)

	query := "SELECT id, firstname, lastname, login, password_hash, role_id, registration_time FROM userDB.users WHERE login = $1"
	log.Printf("Executing query: %s", query)

	var user entity.User
	err := rep.db.QueryRowContext(ctx, query, login).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Login, &user.PasswordHash, &user.RoleId, &user.RegistrationTime)
	if err == sql.ErrNoRows {
		return entity.User{}, entity.ErrUserNotExists
	}

	if err != nil {
		log.Printf("ERROR IN GetUserByLogin: %v", err)
		return entity.User{}, entity.ErrGetUserByLogin
	}

	log.Printf("GetUserByLogin: executed successfully")
	return user, nil
}

// IsUserExists checks whether a user with  the specified login exists in the databse.
func (rep *PostgresUserRepository) IsUserExists(ctx context.Context, login string) (bool, error) {
	log.Printf("IsUserExists: login=%s", login)

	query := "SELECT COUNT(*) FROM userDB.users WHERE login = $1"
	log.Printf("Executing query: %s", query)

	var countRows int
	err := rep.db.QueryRowContext(ctx, query, login).Scan(&countRows)

	if err != nil {
		log.Printf("ERROR IN IsUserExists: %v", err)
		return false, entity.ErrCheckUserExistence
	}

	log.Printf("IsUserExists: executed successfully")
	return countRows > 0, nil
}

// ChangeRoleByLogin updates the user's role by their login.
func (rep *PostgresUserRepository) ChangeRoleByLogin(ctx context.Context, login string, newRoleId int64) error {
	log.Printf("ChangeRoleByLogin: login=%s, newRoleId=%d", login, newRoleId)

	query := "UPDATE userDB.users SET role_id = $1 WHERE login = $2"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, newRoleId, login)
	if err != nil {
		log.Printf("ERROR in ChangeRoleByLogin: %v", err)
		return entity.ErrChangeRole
	}

	log.Printf("ChangeRoleByLogin: executed successfully")
	return nil
}

// DeleteUser removes a user from the database by their login
func (rep *PostgresUserRepository) DeleteUser(ctx context.Context, login string) error {
	log.Printf("DeleteUser: login=%s", login)

	query := "DELETE FROM userDB.users WHERE login = $1"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, login)
	if err != nil {
		log.Printf("ERROR in DeleteUser: %v", err)
		return entity.ErrDeleteUser
	}

	log.Printf("DeletUser: executed successfully")
	return nil
}

// GetAllUsers returns a list of all users from the database.
func (rep *PostgresUserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	log.Printf("GetAllUsers: starting")

	query := "SELECT id, firstname, lastname, login, password_hash, role_id, registration_time FROM userDB.users"
	log.Printf("Executing query: %s", query)

	rows, err := rep.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("ERROR in GetAllUsers: %v", err)
		return nil, entity.ErrGetAllUsers
	}

	defer rows.Close()

	users := make([]entity.User, 0)
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Login, &user.PasswordHash, &user.RoleId, &user.RegistrationTime)
		if err != nil {
			log.Printf("ERROR in GetAllUsers: %v", err)
			return nil, entity.ErrScanUserRow
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("ERROR in GetAllUsers: %v", err)
		return nil, entity.ErrScanUserRow
	}

	log.Printf("GetAllUsers: executed successfully")
	return users, nil
}

func (rep *PostgresUserRepository) ChangeUserName(ctx context.Context, login string, firstName string, lastName string) error {
	log.Printf("ChangeUserName: login=%s, firstName=%s, lastName=%s", login, firstName, lastName)

	query := "UPDATE userDB.users SET firstname=$1, lastname=$2 WHERE login=$3"
	log.Printf("Executing query: %s", query)

	_, err := rep.db.ExecContext(ctx, query, firstName, lastName, login)
	if err != nil {
		log.Printf("ERROR in ChangeUserName: %v", err)
		return entity.ErrChangeName
	}

	log.Printf("ChangeUserName: executed successfully")
	return nil
}
