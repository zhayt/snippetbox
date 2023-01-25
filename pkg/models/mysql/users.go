package mysql

import (
	"database/sql"
	"github.com/zhayt/snippetbox-full-version/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Insert use to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate use to verify whether a user exists with the provided email address and password.
// This will return the relevant user iD if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get use to fetch details for a specific user based on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
