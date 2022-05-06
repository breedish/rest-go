package main

import (
	"errors"
	"fmt"
)

func userIsEligible(email, password string, age int) error {
	if len(email) == 0 {
		return errors.New("email can't be empty")
	}

	if len(password) == 0 {
		return errors.New("password can't be empty")
	}

	if age < 18 {
		return fmt.Errorf("age must be at least %v years old", 18)
	}

	return nil
}
