package users

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewApp(name string) user {
	u := user{Name: name, Type: "app"}
	u.Username = generateHash(5)
	u.Secret = generateHash(18)

	return u
}

func generateHash(n int) string {
	hash := make([]rune, n)
	for i := range hash {
		hash[i] = letters[rand.Intn(len(letters))]
	}
	return string(hash)
}
