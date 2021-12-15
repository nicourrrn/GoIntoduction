package main

import (
	"GoIntoduction/repositories/filesystem"
	"fmt"
)

func main() {
	fmt.Println(filesystem.UserFileRepo{}.GetByEmail("ex@g.com"))
}
