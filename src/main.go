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
	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			recipes.GetEveryDayRecipes(db, w, r)
		case http.MethodPost:
			recipes.CreateRecipe(db, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	// 创建新的路由器
	mux := http.NewServeMux()
	mux.Handle("/", rootHandler)

	// 应用 CORS 中间件
	handlerWithCors := recipes.Cors(mux)

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
//         name TEXT,
//         ingredients TEXT
//     );`
// 	_, err = db.Exec(createTableSQL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 批量插入数据
// 	recipes := []struct {
// 		name        string
// 		ingredients string
// 	}{
// 		{"Pancake", "Flour, Eggs, Milk"},
// 		{"Scrambled Eggs", "Eggs, Butter, Salt"},
// 		{"Salad", "Lettuce, Tomato, Cucumber"},
// 	}

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	stmt, err := tx.Prepare("INSERT INTO recipes (name, ingredients) VALUES (?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	for _, recipe := range recipes {
// 		_, err = stmt.Exec(recipe.name, recipe.ingredients)
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
// 	rows, err := db.Query("SELECT id, name, ingredients FROM recipes")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// 打印查询结果
// 	for rows.Next() {
// 		var id int
// 		var name, ingredients string
// 		err = rows.Scan(&id, &name, &ingredients)
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
