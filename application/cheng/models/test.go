package models

import (
	"fmt"
)

func Helloworld(name string) string{
	return fmt.Sprintf("Hello %s!", name)
}
func Add(a, b int) int{
	c := a + b

	return c
}