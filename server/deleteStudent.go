package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var deleteStudent models.Student

	utils.JsonDeserialize(reqBody, &deleteStudent)
	tx := database.Database.Begin()
	//var parent models.Parents
	result := tx.Delete(&deleteStudent)
	if result.Error == nil {
		tx.Commit()
		fmt.Printf("successfully deleted\n")
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
