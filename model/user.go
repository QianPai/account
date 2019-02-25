package model

type User struct {
	ID interface{} `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
