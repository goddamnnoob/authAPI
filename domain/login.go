package domain

import "database/sql"

type Login struct {
	UserName string         `db:"username"`
	UserId   sql.NullString `db:"customer_id"`
	Accounts sql.NullString `db:"account_numbers"`
	Role     string         `db:"role"`
}
