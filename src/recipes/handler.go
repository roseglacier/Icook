package recipes

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Recipe struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CoverImage string `json:"cover_image"`
	VideoLink  string `json:"video_link"`
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

// 《GET》，从数据库里随机推荐每天的食谱
func GetEveryDayRecipes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	query := "SELECT id , name , cover_image , video_link FROM recipes ORDER BY RAND() LIMIT 3"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		var recipe Recipe //每次迭代中都会创建一个新的 Recipe 变量
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.CoverImage, &recipe.VideoLink); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, recipe)
	}

	respBody := RespBody{Recipes: recipes}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

// 《GET》，根据名字来搜索菜谱
func GetRecipesByName(db *sql.DB, name string) ([]Recipe, error) {
	var recipes []Recipe
	query := "SELECT id, name, cover_image, video_link FROM recipes WHERE name LIKE ?"
	rows, err := db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.CoverImage, &recipe.VideoLink); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

// 《POST》 ToDo 从前端添加数据到数据库
func CreateRecipe(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
