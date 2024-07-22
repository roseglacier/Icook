package recipes

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Recipe struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RespBody struct {
	Recipes []Recipe `json:"recipes"`
}

// CORS 中间件
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 对于预检请求，直接返回
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 《GET》，从数据库里随机推荐每天的推荐食谱
func GetEveryDayRecipes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id , name FROM recipes ORDER BY RAND() LIMIT 3")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		var recipe Recipe //每次迭代中都会创建一个新的 Recipe 变量
		if err := rows.Scan(&recipe.ID, &recipe.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, recipe) //将该变量追加到 recipes 切片中
	}

	respBody := RespBody{Recipes: recipes}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

// 《GET》，从数据库里随机推荐每天的推荐食谱
func GetRecipes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id , name FROM recipes ORDER BY RAND() LIMIT 2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		var recipe Recipe //每次迭代中都会创建一个新的 Recipe 变量
		if err := rows.Scan(&recipe.ID, &recipe.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, recipe) //将该变量追加到 recipes 切片中
	}

	respBody := RespBody{Recipes: recipes}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

// 《POST》 ToDo 从前端添加菜谱到数据库
func CreateRecipe(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
