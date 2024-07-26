package service

import (
	"encoding/json"
	"icook/src/minhaodb"
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
func GetEveryDayRecipes(minhaodb *minhaodb.MinhaoDB, w http.ResponseWriter, r *http.Request) {
	rows, err := minhaodb.QueryWrapper(QGetEveryDayRecipes_x3)
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
func GetRecipesByName(minhaodb *minhaodb.MinhaoDB, name string) ([]Recipe, error) {
	var recipes []Recipe
	rows, err := minhaodb.QueryWrapper(QGetRecipesByName, "%"+name+"%")
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
func CreateRecipe(minhaodb *minhaodb.MinhaoDB, w http.ResponseWriter, r *http.Request) {

}
