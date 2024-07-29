package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

const (
	QGetRandomRecipes = "SELECT id , name , cover_image , video_link FROM recipes ORDER BY RAND() LIMIT ?"
	QGetRecipesByName = "SELECT id, name, cover_image, video_link FROM recipes WHERE name LIKE ?"
)

type IDatabase interface {
	GetRandomItems(count int)
	GetItemsByName(name string)
}

type Database struct {
	db *sql.DB
}

// 《GET》，从数据库里随机推荐每天的食谱
func (d *Database) GetRandomItems(count int) []Recipe {
	rows, err := d.db.Query(QGetRandomRecipes, count)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var recipes []Recipe

	for rows.Next() {
		var recipe Recipe //每次迭代中都会创建一个新的 Recipe 变量
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.CoverImage, &recipe.VideoLink); err != nil {

			return nil
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil
	}

	return recipes
}

// 《GET》，根据名字来搜索菜谱
func (d *Database) GetItemsByName(name string) []Recipe {
	rows, err := d.db.Query(QGetRecipesByName, "%"+name+"%")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var recipes []Recipe

	for rows.Next() {
		var recipe Recipe //每次迭代中都会创建一个新的 Recipe 变量
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.CoverImage, &recipe.VideoLink); err != nil {
			return nil
		}
		recipes = append(recipes, recipe)
	}
	if err := rows.Err(); err != nil {
		return nil
	}

	return recipes
}

func NewDatabase() (*Database, error) {
	dsn := "root:123@tcp(127.0.0.1:3306)/myrecipes"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping不通数据库: %w", err)
	}
	return &Database{db: db}, nil
}
