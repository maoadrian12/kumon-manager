package server

import (
	"fmt"
	"io"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func CreateChild(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := io.ReadAll(r.Body)
	var newChild models.Student
	tx := database.Database.Begin()

	utils.JsonDeserialize(reqBody, &newChild)
	result := tx.Create(&newChild)
	fmt.Printf("child error: %s\n", result.Error)
	if result.Error == nil {
		tx.Commit()
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: "Child has been created",
		})
	} else {
		tx.Rollback()
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "No.",
		})
	}
}
