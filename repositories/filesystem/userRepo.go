package filesystem

import (
	"GoIntoduction/repositories/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type UserFileRepo struct {
}

func (userRepo UserFileRepo) GetByEmail(email string) *models.User {
	file, err := os.OpenFile(fmt.Sprintf("datastore/files/%s.json", email), os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	user := new(models.User)
	err = json.NewDecoder(file).Decode(user)
	if err != nil {
		log.Fatalln(err)
	}
	return user
}
