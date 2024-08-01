// icook/src/main.go

package main

import (
	"fmt"
	"icook/src/database"
	"icook/src/server"
	"icook/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 对于 OPTIONS 请求，直接返回状态码 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func main() {
	// DB
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("start DB failed ", err)
		return
	}

	// server
	serverMachine := server.NewServer(db)

	// http handler
	router := mux.NewRouter()
	controller.Controller(router, serverMachine) // 调用Controller函数
	handler := CorsMiddleware(router)

	// 启动服务器
	fmt.Printf("start at http://localhost:8080/")
	err = http.ListenAndServe(":8080", handler)

	if err != nil {
		fmt.Printf("start server failed: %v\n", err)
	}
}
