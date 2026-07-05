package user

type User struct {
	ID   int
	Name string
}

func New(id int, name string) User {
	return User{
		ID:   id,
		Name: name,
	}
}
