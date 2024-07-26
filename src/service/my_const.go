package service

const (
	getEveryDayRecipes_x3 = "SELECT id , name , cover_image , video_link FROM recipes ORDER BY RAND() LIMIT 3"
	getRecipesByName      = "SELECT id, name, cover_image, video_link FROM recipes WHERE name LIKE ?"
)
