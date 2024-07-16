package recipes

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Recipe struct {
	ID         string `json:"id"`
	RecipeName string `json:"recipe_name"`
}

// 处理Get请求，从数据库里检索食谱
func GetRecipes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id , recipe_name FROM recipes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recipes []map[string]interface{}

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, map[string]interface{}{
			"id":   id,
			"name": name,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func CreateRecipe(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
