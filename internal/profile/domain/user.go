package domain

type User struct {
	ID       string
	Email    string
	Username string

	Following []User
	Followers []User
}
