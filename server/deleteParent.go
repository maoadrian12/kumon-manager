package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func DeleteParent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var deleteParent models.Parents

	utils.JsonDeserialize(reqBody, &deleteParent)
	tx := database.Database.Begin()

	//var parent models.Parents
	result := tx.Delete(&deleteParent)
	tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE")
	if result.Error == nil {
		tx.Commit()
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: "Parent has been deleted",
		})
	} else {
		tx.Rollback()
		fmt.Printf("failure in deletion\n")
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Deletion failed",
		})
	}
}
