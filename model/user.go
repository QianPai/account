package model

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
