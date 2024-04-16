package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	//"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	server "kumondatabase.com/server"
	utils "kumondatabase.com/utils"
)

var dbConnection *pgx.Conn

func updateChild(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	var parsedData map[string]string
	err := json.Unmarshal([]byte(buf.String()), &parsedData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	studentName := parsedData["Student_name"]
	parentName := parsedData["Parent_name"]
	levelName := parsedData["math_level"]
	programName := parsedData["reading_level"]
	programName = strings.ToUpper(programName)
	pagesPerDay := parsedData["pages_per_day"]
	_, err = dbConnection.Exec(context.Background(),
		"INSERT INTO takes6 (student_name, parent_username, level_name, program_name, wkst_per_day) VALUES ($1, $2, $3, $4, $5)",
		studentName, parentName, levelName, "MATH", pagesPerDay)
	_, err = database.Db.Exec("INSERT INTO completes (student_name, parent_username, wkst_num, wkst_lvl, program_name, completion_time, grade) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		studentName, parentName, 0, levelName, "MATH", -1, -1)
	_, err = dbConnection.Exec(context.Background(),
		"INSERT INTO takes6 (student_name, parent_username, level_name, program_name, wkst_per_day) VALUES ($1, $2, $3, $4, $5)",
		studentName, parentName, programName, "READING", pagesPerDay)
	_, err = database.Db.Exec("INSERT INTO completes (student_name, parent_username, wkst_num, wkst_lvl, program_name, completion_time, grade) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		studentName, parentName, 0, programName, "READING", -1, -1)
	if err != nil {
		fmt.Println("Error inserting into takes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error updating child",
		})
	} else {

		if err != nil {
			fmt.Println("Error inserting into completes:", err)
			utils.JsonResponse(w, models.BaseResult{
				Result:  false,
				Message: "error updating child",
			})
		} else {
			utils.JsonResponse(w, models.BaseResult{
				Result:  true,
				Message: "succesful",
			})
		}
	}
}

/*
func allParents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var parents []models.Parents
	result := database.Database.Find(&parents)
	if result.Error == nil {
		utils.JsonResponse(w, parents)
	}
}

func deleteParent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var deleteParent models.Parents

	utils.JsonDeserialize(reqBody, &deleteParent)

	//var parent models.Parents
	result := database.Database.Delete(&deleteParent)
	if result.Error == nil {
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: "Parent has been deleted",
		})
	} else {
		fmt.Printf("failure in deletion\n")
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Deletion failed",
		})
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var deleteStudent models.Student

	utils.JsonDeserialize(reqBody, &deleteStudent)

	//var parent models.Parents
	result := database.Database.Delete(&deleteStudent)
	if result.Error == nil {
		fmt.Printf("successfully deleted\n")
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: "Parent has been deleted",
		})
	} else {
		fmt.Printf("failure in deletion\n")
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Deletion failed",
		})
	}
}

type RequestBody struct {
	Username string `json:"username"`
}

func checkAcc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	var parent models.Parents
	result := database.Database.Where("username = ?", newParent.Username).First(&parent)
	if result.Error == nil {
		utils.JsonResponse(w, parent)
	} else {
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}

func checkStudent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newStudent models.Student

	utils.JsonDeserialize(reqBody, &newStudent)

	var student models.Student
	result := database.Database.Where("name = ?", newStudent.Name).Where("parent_username = ?", newStudent.Parent_username).First(&student)
	if result.Error == nil {
		utils.JsonResponse(w, student)
	} else {
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}

func getParent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	var parent models.Parents
	var result = database.Database.Where("username = ?", newParent.Username).Where("pass = ?", newParent.Pass).First(&parent)
	if result.Error == nil {
		utils.JsonResponse(w, parent)
	} else {
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}

func createParent(w http.ResponseWriter, r *http.Request) {
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

func createChild(w http.ResponseWriter, r *http.Request) {
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



func getPages(w http.ResponseWriter, r *http.Request) {
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

func allStudents(w http.ResponseWriter, r *http.Request) {
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

func getWkst(w http.ResponseWriter, r *http.Request) {
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

func complete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf := new(strings.Builder)
	io.Copy(buf, r.Body)
	parsedData := make(map[string]interface{})
	err := json.Unmarshal([]byte(buf.String()), &parsedData)
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
	_, err = database.Db.Exec("INSERT INTO completes (student_name, parent_username, wkst_num, wkst_lvl, program_name, completion_time, grade) VALUES ($1, $2, $3, $4, $5, $6, $7)", studentUsername, parentUsername, wkstNumber, wkstLevel, programName, completionTime, grade)
	if err != nil {
		fmt.Println("Error inserting into completes:", err)
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

func stats(w http.ResponseWriter, r *http.Request) {
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
	if wkstLevel == "All levels" {
		rows, err = database.Db.Query("SELECT avg(completion_time), avg(grade) FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = $3 AND wkst_num >= $4 AND wkst_num <= $5", studentName, parentName, programName, minWkstNum, maxWkstNum)

	} else {
		rows, err = database.Db.Query("SELECT avg(completion_time), avg(grade) FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = $3 AND wkst_num >= $4 AND wkst_num <= $5 AND wkst_lvl = $6", studentName, parentName, programName, minWkstNum, maxWkstNum, wkstLevel)
	}
	defer rows.Close()
	if err != nil {
		fmt.Println("Error querying from completes:", err)
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "error getting wksts",
		})
	} else {
		rows.Next()
		rows.Scan(&avg_completion_time, &avg_grade)
		utils.JsonResponse(w, models.BaseResult{
			Result:  true,
			Message: fmt.Sprintf("%v %v", avg_completion_time, avg_grade),
		})
	}
}

func getLevels(w http.ResponseWriter, r *http.Request) {
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
	rows, err = database.Db.Query("SELECT DISTINCT wkst_lvl FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = 'READING'",
		studentName, parentName)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&reading_wkst.Wkst_lvl); err != nil {
			log.Fatal(err)
		} else {
			reading_wksts = append(reading_wksts, reading_wkst.Wkst_lvl)
		}
	}
	rows, err = database.Db.Query("SELECT DISTINCT wkst_lvl FROM completes WHERE student_name = $1 AND parent_username = $2 AND program_name = 'MATH'",
		studentName, parentName)
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
*/

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)
	myrouter.HandleFunc("/", server.AllParents).Methods("GET")
	myrouter.HandleFunc("/students", server.AllStudents).Methods("POST")
	myrouter.HandleFunc("/parent", server.GetParent).Methods("POST")
	myrouter.HandleFunc("/delete", server.DeleteParent).Methods("POST")
	myrouter.HandleFunc("/deletstudent", server.DeleteStudent).Methods("POST")
	myrouter.HandleFunc("/createacc", server.CreateParent).Methods("POST")
	myrouter.HandleFunc("/createchild", server.CreateChild).Methods("POST")
	myrouter.HandleFunc("/updatechild", updateChild).Methods("POST")
	myrouter.HandleFunc("/check", server.CheckAcc).Methods("POST")
	myrouter.HandleFunc("/checkstudent", server.CheckStudent).Methods("POST")
	myrouter.HandleFunc("/getpages", server.GetPages).Methods("POST")
	myrouter.HandleFunc("/getinfo", server.GetWkst).Methods("POST")
	myrouter.HandleFunc("/getlevels", server.GetLevels).Methods("POST")
	myrouter.HandleFunc("/complete", server.Complete).Methods("POST")
	myrouter.HandleFunc("/stats", server.Stats).Methods("POST")
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Your React app's origin
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(myrouter)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func main() {
	database_url := fmt.Sprintf("postgres://postgres:%s@localhost:9292/kumon_data", database.Password)

	var err error
	dbConnection, err = pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer database.Connection.Close(context.Background()) // Defer the Close call here
	database.Testing()
	handleRequests()
}
