package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID string `json:"id" ch:"id"`
	TenantID string `json:'tenant_id' ch:'tenant_id'`
	Email string `json:'email' ch:'email'`
	PasswordHash string `json:'-' ch:'password_hash'`
	FirstName string `json:'first_name' ch:'first_name'`
	LastName string `json:'last_name' ch:'last_name'`
	Role string `json:'role' ch:'role'`
	CreateAt time.Time `json:'create_at' ch:'create_at'`
	UpdatedAt time.Time `json:'updated_at' ch:'updated_at'`
	IsActive bool `json:'is_active' ch:'is_active'`
}

type UserModel struct {
	db *sql.DB
}

func (m *UserModel) Create(user *User) error {
	query := `
			INSERT INTO users (
					id, tenant_id, email, password_hash, first_name, last_name,
					role, created_at, updated_at, is_active
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := m.db.Exec(query,
		user.ID, user.TenantID, user.Email, user.PasswordHash, user.FirstName, user.LastName,
		user.Role, user.CreateAt, user.UpdatedAt, user.IsActive,
	)
	return err
}

func (m *UserModel) FindByEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = ? AND is_active = true`
	row := m.db.QueryRow(query, email)

	var user User

	err := row.Scan(query,
		&user.ID, &user.TenantID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
		&user.Role, &user.CreateAt, &user.UpdatedAt, &user.IsActive,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) FindByTenant(tenantID string) ([]User, error) {
	query := `SELECT * FROM users WHERE tenant_id = ? AND is_active = true`
	rows, err := m.db.Query(query, tenantID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.TenantID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
			&user.Role, &user.CreateAt, &user.UpdatedAt, &user.IsActive,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}