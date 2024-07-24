package main

import (
	"icook/src/appinit"
	"log"
)

func main() {

	//初始化DB
	db, err := appinit.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//启动服务器
	if err := appinit.StartServer(db); err != nil {
		log.Fatal(err)
	}
}
