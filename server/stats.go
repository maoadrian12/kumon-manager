package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	var parsedData map[string]interface{}
	err := json.Unmarshal([]byte(buf.String()), &parsedData)
	if err != nil {
		fmt.Println("Error in go.")
		fmt.Println("Error parsing JSON:", err)
		return
	}
	// Extract variables from the parsed datastudentName := parsedData["student_username"].(string)studentName := parsedData["student_username"].(string)
	studentName := parsedData["student_username"].(string)
	parentName := parsedData["parent_username"].(string)
	programName := parsedData["program_name"].(string)
	minWkstNum := int(parsedData["min_wkst_number"].(float64))
	maxWkstNum := int(parsedData["max_wkst_number"].(float64))
	wkstLevel := parsedData["wkst_level"].(string)
	var rows *sql.Rows
	var avg_completion_time float64
	var avg_grade float64
	var stmt *sql.Stmt
	//var err error
	if wkstLevel == "All levels" {
		stmt, err = database.Db.Prepare("SELECT avg(completion_time), avg(grade) FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = $3 AND wkst_num >= $4 AND wkst_num <= $5")

	} else {
		stmt, err = database.Db.Prepare("SELECT avg(completion_time), avg(grade) FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = $3 AND wkst_num >= $4 AND wkst_num <= $5 AND wkst_lvl = $6")
	}
	defer stmt.Close()
	if err != nil {
		fmt.Println("Error querying from completes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error getting wksts",
		})
	} else {
		rows, err = stmt.Query(studentName, parentName, programName, minWkstNum, maxWkstNum, wkstLevel)
		rows.Next()
		rows.Scan(&avg_completion_time, &avg_grade)
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: fmt.Sprintf("%v %v", avg_completion_time, avg_grade),
		})
	}
}
