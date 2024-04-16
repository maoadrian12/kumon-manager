package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func GetParent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	var parent models.Parents
	var result = database.Database.Where("username = ?", newParent.Username).Where("pass = ?", newParent.Pass).First(&parent)
	if result.Error == nil {
		utils.JsonResponse(w, parent)
	} else {
		fmt.Println("Error getting parent " + result.Error.Error())
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}
