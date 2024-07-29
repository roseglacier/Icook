package server

import (
	"icook/src/database"
)

type IServer interface {
	GetEveryDayRecipes()
	GetRecipesByName()
}

type Server struct {
	db *database.Database
}

type Recipe struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CoverImage string `json:"cover_image"`
	VideoLink  string `json:"video_link"`
}

// GetEveryDayRecipes
type GetEveryDayRecipesArgs struct {
	Count int
}

type GetEveryDayRecipesRet struct {
	Recipes []Recipe
}

func (th *Server) GetEveryDayRecipes(args GetEveryDayRecipesArgs) GetEveryDayRecipesRet {
	dbRecipes := th.db.GetRandomItems(args.Count)

	// 将 database.Recipe 转换为 Recipe
	var recipes []Recipe
	for _, dbRecipe := range dbRecipes {
		recipes = append(recipes, Recipe{
			ID:         dbRecipe.ID,
			Name:       dbRecipe.Name,
			CoverImage: dbRecipe.CoverImage,
			VideoLink:  dbRecipe.VideoLink,
		})
	}
	// fmt.Printf("Converted Recipes: %+v\n", recipes)

	return GetEveryDayRecipesRet{Recipes: recipes}

}

// GetRecipesByName
type GetRecipesByNameArgs struct {
	name string
}

type GetRecipesByNameRet struct {
	Recipes []Recipe
}

func (th *Server) GetRecipesByName(args GetRecipesByNameArgs) GetRecipesByNameRet {
	th.db.GetItemsByName(args.name)
	return GetRecipesByNameRet{}
}

func NewServer(db *database.Database) *Server {
	return &Server{db: db}
}
