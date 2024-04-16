package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func GetLevels(w http.ResponseWriter, r *http.Request) {
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
	var reading_wksts []string
	var math_wksts []string
	var rows *sql.Rows
	stmt, err := database.Db.Prepare("SELECT DISTINCT wkst_lvl FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = 'READING'")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()
	rows, err = stmt.Query(studentName, parentName)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&reading_wkst.Wkst_lvl); err != nil {
			log.Fatal(err.Error())
		} else {
			reading_wksts = append(reading_wksts, reading_wkst.Wkst_lvl)
		}
	}
	stmt, err = database.Db.Prepare("SELECT DISTINCT wkst_lvl FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = 'MATH'")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(studentName, parentName)
	if err != nil {
		fmt.Println("Error querying from completes:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&math_wkst.Wkst_lvl); err != nil {
			log.Fatal(err)
		} else {
			math_wksts = append(math_wksts, math_wkst.Wkst_lvl)
		}
	}
	if err != nil {
		fmt.Println("Error querying from completes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error getting wksts",
		})
	} else {
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: fmt.Sprintf("%v %v", reading_wksts, math_wksts),
		})
	}
}
