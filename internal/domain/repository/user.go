package repository

import (
	"context"
	"database/sql"

	"github.com/kobayashiyabako16g/tiny-go/internal/domain/model"
	"github.com/kobayashiyabako16g/tiny-go/pkg/db"
	"github.com/kobayashiyabako16g/tiny-go/pkg/logger"
)

type Users interface {
	FindById(ctx context.Context, id int) (*model.User, error)
}

type usersRepositroy struct {
	db *db.Client
}

func NewUsersRepository(db *db.Client) Users {
	return usersRepositroy{db}
}

func (r usersRepositroy) FindById(ctx context.Context, id int) (*model.User, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT id, name, email FROM users WHERE id = ?")
	if err != nil {
		logger.Error(ctx, "Database Error", err)
		return nil, err
	}
	defer stmt.Close()

	var user model.User
	if err = stmt.QueryRowContext(ctx, id).Scan(&user.Id, &user.Name, &user.Email); err != nil {
		// Not Found
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// Error
		logger.Error(ctx, "Database Error", err)
		return nil, err
	}

	return &user, nil
}
