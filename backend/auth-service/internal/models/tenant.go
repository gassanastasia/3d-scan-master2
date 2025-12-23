package models

import (
	"database/sql"
	"time"
)

type Tenant struct {
	ID string "json:'id' ch:'id'"
	Name string "json:'name' ch:'name'"
	Description string "json:'description' ch:'description'"
	Plan string "json:'plan' ch:'plan'"
	MaxUsers int "json:'max_users' ch:'max_users'"
	MaxPrinters int "json:'max_printers' ch:'max_printers'"
	IsActive bool "json:'is_active' ch:'is_active'"
	CreateAt time.Time "json:'create_at' ch:'create_at'"
	UpdatedAt time.Time "json:'updated_at' ch:'updated_at'"
}

type TenantModel struct {
	db *sql.DB
}

func NewTenantModel(db *sql.DB) *TenantModel {
	return &TenantModel{db: db}
}
func (m *TenantModel) Create(tenant *Tenant) error {
	query := `
			INSERT INTO tenants (
					id, name, description, plan, max_users, max_printers,
					is_active, created_at, updated_at
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := m.db.Exec(query,
		tenant.ID, tenant.Name, tenant.Description, tenant.Plan, tenant.MaxUsers, tenant.MaxPrinters,
		tenant.IsActive, tenant.CreateAt, tenant.UpdatedAt,
	)
	return err
}

func (m *TenantModel) FindByID(id string) (*Tenant, error) {
	query := `SELECT * FROM tenants WHERE id = ? AND is_active = true`
	row := m.db.QueryRow(query, id)

	var tenant Tenant

	err := row.Scan(query,
		&tenant.ID, &tenant.Name, &tenant.Description, &tenant.Plan, &tenant.MaxUsers, &tenant.MaxPrinters,
		&tenant.IsActive, &tenant.CreateAt, &tenant.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &tenant, nil
}