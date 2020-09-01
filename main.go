package main

import (
	"blog/model"
	"blog/routers"
)

func main() {
	model.InitDB()
	routers.InitRouters()


}
