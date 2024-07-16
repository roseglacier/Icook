package main

import (
	"database/sql"
	"fmt"
	"icook/src/recipes"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//链接SQLite数据库
	db, err := sql.Open("sqlite3", "./recipes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 注册根路径处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			recipes.GetRecipes(db, w, r)
		case http.MethodPost:
			recipes.CreateRecipe(db, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// 启动服务器，监听 8080 端口
	fmt.Println("Starting server at :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
