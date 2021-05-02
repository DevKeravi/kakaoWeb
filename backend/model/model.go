package model

import (
	"encoding/json"
)

type userData struct {
	Name      string `json:"nickName"`
	Img       string `json:"profileImageURL"`
	Thumbnail string `json:"thumbnailURL"`
	Country   string `json:"countryISO"`
}

var userList map[string]*userData

var size int

func Create(data []byte) {
	temp := &userData{}

	json.Unmarshal(data, temp)

	userList[temp.Name] = temp
}

func NewData(data string) userData {
	temp := &userData{}
	json.Unmarshal([]byte(data), temp)

	return *temp
}

func Get(Name string) userData {
	temp := userData{}
	if _, ok := userList[Name]; ok {
		temp = *userList[Name]
	}
	return temp
}

func Init() {
	userList = make(map[string]*userData)
	size = 0
}
