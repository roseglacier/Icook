package main

import (
	"database/sql"
	"fmt"
	"icook/src/recipes"
	"icook/src/routes"

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

	// 创建路由器
	r := routes.NewRouter(db)

	// 应用 CORS 中间件
	handlerWithCors := recipes.Cors(r)

	// 启动服务器，监听 8080 端口
	fmt.Println("Starting server at :8080")

	if err := http.ListenAndServe(":8080", handlerWithCors); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	// 链接SQLite数据库
// 	db, err := sql.Open("sqlite3", "./recipes.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// 创建表
// 	createTableSQL := `CREATE TABLE IF NOT EXISTS recipes (
//         id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
//         name TEXT
//     );`
// 	_, err = db.Exec(createTableSQL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 批量插入数据
// 	recipes := []struct {
// 		name string
// 	}{
// 		{"披萨"},
// 		{"蛋炒饭"},
// 		{"沙拉"},
// 		{"肉沫茄子"},
// 		{"紫菜蛋汤"},
// 		{"青椒肉丝"},
// 		{"孜然土豆片"},
// 		{"土豆肉丝"},
// 	}

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	stmt, err := tx.Prepare("INSERT INTO recipes (name) VALUES (?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	for _, recipe := range recipes {
// 		_, err = stmt.Exec(recipe.name)
// 		if err != nil {
// 			tx.Rollback()
// 			log.Fatal(err)
// 		}
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 查询数据
// 	rows, err := db.Query("SELECT id, name FROM recipes")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// 打印查询结果
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("Recipe: %d, %s\n", id, name)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
