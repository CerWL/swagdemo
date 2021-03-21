package db

//User model
type User struct {
	UserID int  `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
