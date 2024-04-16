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

func GetWkst(w http.ResponseWriter, r *http.Request) {
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
	var reading_wkst models.Completes
	var math_wkst models.Completes
	err = database.Db.QueryRow("SELECT * FROM completes WHERE student_name = $1 AND parent_username = $2"+
		"AND program_name = 'READING' ORDER BY wkst_lvl, wkst_num DESC LIMIT 1", studentName, parentName).Scan(&reading_wkst.Student_name, &reading_wkst.Parent_username, &reading_wkst.Wkst_num, &reading_wkst.Wkst_lvl, &reading_wkst.Program_name, &reading_wkst.Completion_time, &reading_wkst.Grade)

	if err != nil {
		fmt.Println("Error querying from completes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error getting wksts",
		})
	}
	err = database.Db.QueryRow("SELECT * FROM completes WHERE student_name = $1 AND parent_username = $2"+
		"AND program_name = 'MATH' ORDER BY wkst_lvl, wkst_num DESC LIMIT 1", studentName, parentName).Scan(&math_wkst.Student_name, &math_wkst.Parent_username, &math_wkst.Wkst_num, &math_wkst.Wkst_lvl, &math_wkst.Program_name, &math_wkst.Completion_time, &math_wkst.Grade)
	if err != nil {
		fmt.Println("Error querying from completes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error getting wksts",
		})
	} else {
		//fmt.Println(reading_wkst)
		//fmt.Println(math_wkst)
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: fmt.Sprintf("%v %v", reading_wkst, math_wkst),
		})
	}
}
