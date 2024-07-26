package routes

import (
	"encoding/json"
	"icook/src/minhaodb"
	"icook/src/service"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(minhaodb *minhaodb.MinhaoDB) *mux.Router {
	// 创建新的路由器
	r := mux.NewRouter()

	// 《GET》--------获取每日推荐菜谱的处理函数
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			service.GetEveryDayRecipes(minhaodb, w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)

	// 《GET》--------搜索菜谱的处理函数
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			if name == "" {
				http.Error(w, "Missing name parameter", http.StatusBadRequest)
				return
			}

			recipes, err := service.GetRecipesByName(minhaodb, name)
			if err != nil {
				http.Error(w, "Failed to search recipes", http.StatusInternalServerError)
				return
			}

			response, err := json.Marshal(recipes)
			if err != nil {
				http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)

	return r
}
