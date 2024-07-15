package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//链接SQLite数据库
	db, err := sql.Open("sqlite", "./recipes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//创建数据库的表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		recipe_name TEXT,
	)`)

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
