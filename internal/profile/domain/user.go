package domain

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`

	Following []User `json:"following"`
	Followers []User `json:"followers"`
}
