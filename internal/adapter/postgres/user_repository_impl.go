package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"root_academy_go/internal/domain/model/user"
	"root_academy_go/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(u *user.User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, roles, email, password, instructor_title, instructor_bio, avatar)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(query,
		u.ID,
		u.FirstName,
		u.LastName,
		rolesToString(u.Roles),
		u.Email,
		u.Password,
		u.InstructorTitle,
		u.InstructorBio,
		u.Avatar,
	)
	if err != nil {
		return fmt.Errorf("create user %s: %w", u.ID, err)
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*user.User, error) {
	query := `
		SELECT id, first_name, last_name, roles, email, password, instructor_title, instructor_bio, avatar
		FROM users
		WHERE email = $1
	`

	return r.scanUser(r.db.QueryRow(query, email))
}

func (r *UserRepositoryImpl) GetUserByID(id string) (*user.User, error) {
	query := `
		SELECT id, first_name, last_name, roles, email, password, instructor_title, instructor_bio, avatar
		FROM users
		WHERE id = $1
	`

	return r.scanUser(r.db.QueryRow(query, id))
}

func (r *UserRepositoryImpl) UpdateUser(u *user.User) error {
	query := `
		UPDATE users SET
			first_name = $2,
			last_name = $3,
			roles = $4,
			email = $5,
			password = $6,
			instructor_title = $7,
			instructor_bio = $8,
			avatar = $9
		WHERE id = $1
	`

	_, err := r.db.Exec(query,
		u.ID,
		u.FirstName,
		u.LastName,
		rolesToString(u.Roles),
		u.Email,
		u.Password,
		u.InstructorTitle,
		u.InstructorBio,
		u.Avatar,
	)
	if err != nil {
		return fmt.Errorf("update user %s: %w", u.ID, err)
	}

	return nil
}

func (r *UserRepositoryImpl) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user %s: %w", id, err)
	}

	return nil
}

// scanUser lit une ligne users dans un *user.User. Retourne (nil, nil) si aucune ligne ne correspond.
func (r *UserRepositoryImpl) scanUser(row *sql.Row) (*user.User, error) {
	var u user.User
	var roles string

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&roles,
		&u.Email,
		&u.Password,
		&u.InstructorTitle,
		&u.InstructorBio,
		&u.Avatar,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}

	u.Roles = rolesFromString(roles)

	return &u, nil
}

func rolesToString(roles []user.Role) string {
	parts := make([]string, len(roles))
	for i, role := range roles {
		parts[i] = string(role)
	}
	return strings.Join(parts, ",")
}

func rolesFromString(roles string) []user.Role {
	if roles == "" {
		return nil
	}

	parts := strings.Split(roles, ",")
	result := make([]user.Role, len(parts))
	for i, p := range parts {
		result[i] = user.Role(p)
	}
	return result
}
