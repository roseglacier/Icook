package appinit

import (
	"fmt"
	"icook/src/minhaodb"
	"icook/src/routes"
	"icook/src/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库
func InitDB() (*minhaodb.MinhaoDB, error) {
	dsn := "root:123@tcp(127.0.0.1:3306)/myrecipes"
	minhaodb, err := minhaodb.NewMinhaoDB(dsn)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	return minhaodb, nil
}

// 启动server
func StartServer(minhaodb *minhaodb.MinhaoDB) error {
	r := routes.NewRouter(minhaodb)    // 创建路由器
	handlerWithCors := service.Cors(r) // 应用 CORS 中间件

	fmt.Println("Starting server at :http://localhost:8080/")
	// fmt.Println("Start server successful !")
	if err := http.ListenAndServe(":8080", handlerWithCors); err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}
	return nil
}
