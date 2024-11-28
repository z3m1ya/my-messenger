package entities

type User struct {
	ID       string
	Username string
	Password string
}

func NewUser(id, username, password string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
	}
}
