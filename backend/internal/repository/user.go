package repository

import (
	"context"

	"github.com/SYSU-ECNC/ecnc-oa/backend/internal/domain"
)

func (repo *Repository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := domain.User{
		Username: username,
	}

	query := `
		SELECT id, password_hash, full_name, role, is_active, created_at, updated_at
		FROM users 
		WHERE username = $1
	`
	if err := repo.db.QueryRowContext(ctx, query, user.Username).Scan(
		&user.ID,
		&user.PasswordHash,
		&user.FullName,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
