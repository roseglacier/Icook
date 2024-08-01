package controller

import (
	"encoding/json"
	"icook/src/server"
	"net/http"

	"github.com/gorilla/mux"
)



func Controller(router *mux.Router, serverMachine *server.Server) {
	// 处理主页函数（当前为随即推荐用户期望的数目）
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var args server.GetEveryDayRecipesArgs
		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			http.Error(w, "Unable to parse request body", http.StatusBadRequest)
			return
		}
		result := serverMachine.GetEveryDayRecipes(args)

		response, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	}).Methods(http.MethodPost)

	// // *******    注册处理函数 GetEveryDayRecipes
	// router.HandleFunc("/api/GetEveryDayRecipes", func(w http.ResponseWriter, r *http.Request) {
	// 	var args server.GetEveryDayRecipesArgs
	// 	err := json.NewDecoder(r.Body).Decode(&args)
	// 	if err != nil {
	// 		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
	// 		return
	// 	}
	// 	result := serverMachine.GetEveryDayRecipes(args)

	// 	response, err := json.Marshal(result)
	// 	if err != nil {
	// 		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	fmt.Printf("123")
	// 	w.Write(response)

	// }).Methods(http.MethodPost)

	// *******   注册处理函数GetRecipesByName
	router.HandleFunc("/api/GetRecipesByName", func(w http.ResponseWriter, r *http.Request) {

	}).Methods(http.MethodPost)
}
