package model

import (
	"encoding/json"
	"log"
)

type userData struct {
	Id        int    `json:"id"`
	Name      string `json:"nickName"`
	Img       string `json:"profileImageURL"`
	Thumbnail string `json:"thumbnailURL"`
	Country   string `json:"countryISO"`
}

var userList map[int]*userData

var size int

func Create(data []byte) {
	temp := &userData{}

	json.Unmarshal(data, temp)
	temp.Id = size
	log.Printf("\n Id : %d \n Name : %s \n Img : %s \n ", temp.Id, temp.Name, temp.Img)

	size++
	userList[temp.Id] = temp
}

func Init() {
	userList = make(map[int]*userData)
	size = 0
}
