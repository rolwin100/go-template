package repository

import "database/sql"

func (repository *RepositoryConfig) CreateUser(name string, email string, password string) sql.Result {
	// Create a new user.
	// defer repository.db.Close()
	result := repository.db.MustExec("INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)", name, email, password)
	return result
}

func (repository *RepositoryConfig) GetUserByEmail(email string) (*sql.Rows, error) {
	// Get a user by email.
	// defer repository.db.Close()
	return repository.db.Query("SELECT id FROM users WHERE email = $1", email)

}
