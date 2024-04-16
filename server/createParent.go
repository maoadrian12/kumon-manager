package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func CreateParent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	result := database.Database.Create(&newParent)
	fmt.Printf("Parent error: %s\n", result.Error)
	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Parent has been created",
	})
}
