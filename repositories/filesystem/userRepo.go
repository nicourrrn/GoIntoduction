package filesystem

import (
	"GoIntoduction/repositories/models"
	"fmt"
	"io"
	"log"
	"os"
)

type UserFileRepo struct {
}

func (userRepo UserFileRepo) GetByEmail(email string) *models.User {
	file, err := os.OpenFile("datastore/files/user_1.json", os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	data := make([]byte, 64)
	final_json := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		final_json += string(data[:n])
	}
	fmt.Println(final_json)
	return nil
}
