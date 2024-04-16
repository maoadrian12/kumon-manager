package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func CheckAcc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	var parent models.Parents
	result := database.Database.Where("username = ?", newParent.Username).First(&parent)
	if result.Error == nil {
		utils.JsonResponse(w, parent)
	} else {
		fmt.Println("Error checking account " + result.Error.Error())
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}
