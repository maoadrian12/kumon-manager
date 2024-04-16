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

	//var parent models.Parents
	result := database.Database.Delete(&deleteStudent)
	if result.Error == nil {
		fmt.Printf("successfully deleted\n")
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: "Parent has been deleted",
		})
	} else {
		fmt.Printf("failure in deletion\n")
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Deletion failed",
		})
	}
}
