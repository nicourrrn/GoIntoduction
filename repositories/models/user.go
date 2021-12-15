package models

import "time"

type User struct {
	Login, Email, Hash string
	TimeStamp          time.Time
}

//
//func NewUser()  {
//
//}