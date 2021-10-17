package domain

import (
	"github.com/goddamnnoob/notReddit/errs"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppError)
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) FindBy(username string, password string) (*Login, *errs.AppError) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u
                  LEFT JOIN accounts a ON a.customer_id = u.customer_id
                WHERE username = ? and password = ?
                GROUP BY a.customer_id`
	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		return nil, errs.NewNotFoundError("Failed Authentication")
	}
	return &login, nil
}

func NewAuthRepository(client sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client: &client}
}
