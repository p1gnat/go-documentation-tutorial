package greetings

import "fmt"

func Hello (name string) string {
	string := fmt.Sprintf("Hi %v, welcome!", name)
	return string
}