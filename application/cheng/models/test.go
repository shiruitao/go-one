package models

import (
	"fmt"
)

func Helloworld(name string) string{
	return fmt.Sprintf("Hello %s!", name)
}