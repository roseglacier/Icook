package routes

import (
	"database/sql"
	"icook/src/recipes"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	// 创建新的路由器
	r := mux.NewRouter()

	// 《GET》--------获取每日推荐菜谱的处理函数
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			recipes.GetEveryDayRecipes(db, w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)

	// 《POST》--------创建菜谱的处理函数
	r.HandleFunc("/createrecipe", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			recipes.CreateRecipe(db, w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodPost)

	return r
}
