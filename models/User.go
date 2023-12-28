package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Username string `gorm:"not null"`
	Password string
	Confirm  bool   `gorm:"not null"`
	Rol      string `gorm:"not null"`
	Sub      string `gorm:"->:false;<-:create;"`
}

func (u *User) ConvertMapStruct(jsonStruct interface{}) {
	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		fmt.Println("Error convert")
		return
	}

	json.Unmarshal(jsonData, u)
}
