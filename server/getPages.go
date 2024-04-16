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

func GetPages(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	var parsedData map[string]string
	err := json.Unmarshal([]byte(buf.String()), &parsedData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	// Extract variables from the parsed data
	studentName := parsedData["name"]
	parentName := parsedData["parent_username"]
	wkst := 0
	err = database.Db.QueryRow("SELECT wkst_per_day FROM takes6 WHERE student_name = $1 AND parent_username = $2"+
		"GROUP BY wkst_per_day LIMIT 1", studentName, parentName).Scan(&wkst)
	if err != nil {
		fmt.Println("Error inserting into takes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error updating child",
		})
	} else {
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: fmt.Sprintf("%v", wkst),
		})
	}
}
