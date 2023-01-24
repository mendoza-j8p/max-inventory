package repository

import (
	"context"
	"max-inventory/internal/entity"
)

const ( 
	qryInsertUser = `
		insert into USERS (email, name, password)
		values (?,?,?);
	`

	qryGetUserByEmail = `
		select 
			id,
			email, name,
			password
		from USERS
		where email =?;
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error)  {
	U := &entity.User{}
	err := r.db.GetContext(ctx, U, qryGetUserByEmail, email)

	return U, err
}
