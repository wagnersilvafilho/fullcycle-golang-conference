package users

import "errors"

type profile struct {
	user
	Age uint8
}

func NewUser(name, username, pass string, age uint8) profile {
	u := profile{Age: age}
	u.user = user{
		Name:     name,
		Username: username,
		Secret:   pass,
		Type:     "user",
	}

	return u
}

func (p profile) Validate() error {
	err := p.user.Validate()
	if err != nil {
		return err
	}
	if p.Age < 15 {
		return errors.New("min age is 15")
	}
	return nil
}
