package main

import "kakaoWeb/backend/app"

func main() {
	app.Init()
	app.Run(":10102")
}
