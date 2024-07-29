# Icook

## 代码篇 （Go）
1.GET

    - GetEveryDayRecipes(db *sql.DB, w http.ResponseWriter, r *http.Request){} //随机推荐每天的食谱
    - GetRecipesByName(db *sql.DB, name string) ([]Recipe, error) {} //根据名字来搜索菜谱
2.POST

    - func CreateRecipe(db *sql.DB, w http.ResponseWriter, r *http.Request) //ToDo 从前端添加数据到数据库
3.PUT
4.DELETE

## 数据库篇（MariaDB）
1.创建表 3个
    
    - 创建菜谱表
    CREATE TABLE recipes (
        id INT PRIMARY KEY,
        name VARCHAR(255),
        cover_image VARCHAR(255),
        video_link VARCHAR(255)
    );

    - 创建标签表
    CREATE TABLE tags (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50) NOT NULL UNIQUE
    );

    - 创建菜谱标签关联表
    CREATE TABLE recipe_tags (
        recipe_id INT,
        tag_id INT,
        FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
        FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
        PRIMARY KEY (recipe_id, tag_id)
    );

2.在excel中编辑好需要的数据，保存为csv文件，再通过命令导入到数据库里去。

    - 类似下方==>
        -- 先禁用外键检查
        SET foreign_key_checks = 0;

        -- 删除表中的数据
        DROP TABLE IF EXISTS recipes;

        --创建表
        -- 导入数据（确保外键表数据已存在）
        LOAD DATA INFILE 'recipes.csv'
        INTO TABLE recipes
        FIELDS TERMINATED BY ',' 
        ENCLOSED BY '"'
        LINES TERMINATED BY '\r\n'
        IGNORE 1 ROWS
        (id, name, cover_image, video_link);
        
        -- 启用外键检查
        SET foreign_key_checks = 1;
////////////

        LOAD DATA INFILE 'tags.csv'
        INTO TABLE tags
        FIELDS TERMINATED BY ',' 
        ENCLOSED BY '"'
        LINES TERMINATED BY '\r\n'
        IGNORE 1 ROWS
        (id, name);

        LOAD DATA INFILE 'recipe_tags.csv'
        INTO TABLE recipe_tags
        FIELDS TERMINATED BY ',' 
        ENCLOSED BY '"'
        LINES TERMINATED BY '\r\n'
        IGNORE 1 ROWS
        (recipe_id, tag_id);
