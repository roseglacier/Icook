package main

import (
	"database/sql"
	"fmt"
	"icook/src/recipes"
	"icook/src/routes"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // MariaDB 驱动
)

func main() {

	//链接数据库
	dsn := "root:123@tcp(127.0.0.1:3306)/myrecipes"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := routes.NewRouter(db) // 创建路由器

	handlerWithCors := recipes.Cors(r) // 应用 CORS 中间件

	fmt.Println("Starting server at :8080")

	if err := http.ListenAndServe(":8080", handlerWithCors); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
