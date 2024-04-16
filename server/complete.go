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

func Complete(w http.ResponseWriter, r *http.Request) {
	tx, err := database.Db.Begin()
	defer r.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	parsedData := make(map[string]interface{})
	err = json.Unmarshal([]byte(buf.String()), &parsedData)
	//var parsedData map[string]string
	//err := json.Unmarshal([]byte(buf.String()), &parsedData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	// Extract variables from the parsed data
	studentUsername := parsedData["student_username"]
	parentUsername := parsedData["parent_username"]
	wkstNumber := parsedData["wkst_number"]
	wkstLevel := parsedData["wkst_level"]
	programName := parsedData["program_name"]
	completionTime := parsedData["completion_time"]
	grade := parsedData["grade"]
	wkst := 0
	tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE")
	_, err = tx.Exec("INSERT INTO completes (student_name, parent_username, wkst_num, wkst_lvl, program_name, completion_time, grade) VALUES ($1, $2, $3, $4, $5, $6, $7)", studentUsername, parentUsername, wkstNumber, wkstLevel, programName, completionTime, grade)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error inserting into completes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error updating child",
		})
	} else {
		tx.Commit()
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: fmt.Sprintf("%v", wkst),
		})
	}
}
