package dto

type LoginRequest struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}
