package main

import (
	"encoding/json"
	"fmt"
	"icook/src/database"
	"icook/src/server"
	"net/http"

	"github.com/gorilla/mux"
)

func Controller(router *mux.Router, serverMachine *server.Server) {
	// 注册处理函数0
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)

		errorMessage := map[string]string{"error": "invaild url", "url": r.URL.Path}
		response, err := json.Marshal(errorMessage)
		if err != nil {
			http.Error(w, "Error encoding error response", http.StatusInternalServerError)
			return
		}

		w.Write(response)

	}).Methods(http.MethodPost)

	// 注册处理函数1
	router.HandleFunc("/api/GetEveryDayRecipes", func(w http.ResponseWriter, r *http.Request) {
		var args server.GetEveryDayRecipesArgs
		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			http.Error(w, "Unable to parse request body", http.StatusBadRequest)
			return
		}
		// 获取菜谱数据
		result := serverMachine.GetEveryDayRecipes(args)

		// 将结果编码为 JSON
		response, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

		// 设置响应头并写入响应体
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	}).Methods(http.MethodPost)

	// 注册处理函数2
	router.HandleFunc("/api/GetRecipesByName", func(w http.ResponseWriter, r *http.Request) {

	}).Methods(http.MethodPost)
}

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

	//DB
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("start DB failed ", err)
		return
	}

	//server
	serverMachine := server.NewServer(db)

	// http handeler
	router := mux.NewRouter()
	Controller(router, serverMachine)
	handler := CorsMiddleware(router)

	// 启动服务器
	fmt.Printf("start at http://localhost:8080/")
	err = http.ListenAndServe(":8080", handler)

	if err != nil {
		fmt.Printf("start server failed: %v\n", err)
	}

}
