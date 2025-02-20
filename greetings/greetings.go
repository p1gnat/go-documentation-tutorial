package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello (name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	string := fmt.Sprintf(randomFormat(), name)
	return string, nil
}

func randomFormat() string {
	formats := []string{
		"Hi %v, welcome!",
		"Great to see you %v",
		"Hail, %v, Well met!",
	}
	return formats[rand.Intn(len(formats))]
}