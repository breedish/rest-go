package database

import (
	"errors"
	"time"
)

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()

	if err != nil {
		return User{}, err
	}

	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}

	if _, ok := db.Users[user.Email]; ok {
		return user, errors.New("duplicate user")
	}

	db.Users[user.Email] = user
	err = c.updateDB(db)
	return user, err
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}

	user.Password = password
	user.Name = name
	user.Age = age

	return user, err
}

func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")
	}
	return user, nil
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	if _, ok := db.Users[email]; ok {
		delete(db.Users, email)
	}

	return c.updateDB(db)
}
