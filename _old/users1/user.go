package users1

type user struct {
	Name string
	Age  uint8
}

func New(name string, age uint8) user {
	return user{
		Name: name,
		Age:  age,
	}
}

func AddYear(u *user) {
	u.Age++
}
