package appinit

import (
	"database/sql"
	"fmt"
	"icook/src/recipes"
	"icook/src/routes"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := "root:123@tcp(127.0.0.1:3306)/myrecipes"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func StartServer(db *sql.DB) error {
	r := routes.NewRouter(db)          // 创建路由器
	handlerWithCors := recipes.Cors(r) // 应用 CORS 中间件

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", handlerWithCors); err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}
	return nil
}
