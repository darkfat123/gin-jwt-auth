package model

type User struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
