package server

import (
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func AllParents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var parents []models.Parents
	result := database.Database.Find(&parents)
	if result.Error == nil {
		utils.JsonResponse(w, parents)
	}
}
