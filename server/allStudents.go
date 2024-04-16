package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func AllStudents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var students []models.Student
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	var parsedData map[string]string
	err := json.Unmarshal([]byte(buf.String()), &parsedData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Error parsing JSON",
		})
		return
	}
	parentName := parsedData["Parent_username"]
	result := database.Database.Where("parent_username = ?", parentName).Find(&students)
	if result.Error == nil {
		utils.JsonResponse(w, students)
	} else {
		fmt.Printf("error is %s\n", result.Error)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}
