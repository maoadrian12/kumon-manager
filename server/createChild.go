package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func CreateChild(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newChild models.Student

	utils.JsonDeserialize(reqBody, &newChild)
	result := database.Database.Create(&newChild)
	fmt.Printf("child error: %s\n", result.Error)
	if result.Error == nil {
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: "Child has been created",
		})
	} else {
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "No.",
		})
	}
}
