package main

import (
	"icook/src/appinit"
	"log"
)

func main() {

	//初始化DB
	minhaoDB, err := appinit.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer minhaoDB.Close()

	//启动服务器
	if err := appinit.StartServer(minhaoDB); err != nil {
		log.Fatal(err)
	}
}
