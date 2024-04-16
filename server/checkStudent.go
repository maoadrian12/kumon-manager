package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func CheckStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newStudent models.Student

	utils.JsonDeserialize(reqBody, &newStudent)

	var student models.Student
	result := database.Database.Where("name = ?", newStudent.Name).Where("parent_username = ?", newStudent.Parent_username).First(&student)
	if result.Error == nil {
		utils.JsonResponse(w, student)
	} else {
		fmt.Println("Error checking student " + result.Error.Error())
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}
