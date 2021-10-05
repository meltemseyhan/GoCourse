package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new User must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("no user found with id '%v'", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("no user found with id '%v'", u.ID)
}

func RemoveUserByID(id int) error {
	for i, candidate := range users {
		if candidate.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no user found with id '%v'", id)
}
