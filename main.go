package main

import "GoIntoduction/repositories/filesystem"

func main() {
	filesystem.UserFileRepo{}.GetByEmail("Somedata")
}
