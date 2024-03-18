package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	database "kumondatabase.com/database"
	models "kumondatabase.com/models"
	utils "kumondatabase.com/utils"
)

func initiateParents() models.Parents {
	var Parent models.Parents
	return Parent
}

func allParents(w http.ResponseWriter, r *http.Request) {
	var parents []models.Parents
	result := database.Database.Find(&parents)
	if result.Error == nil {
		utils.JsonResponse(w, parents)
	}
}

func deleteParent(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	var deletedArticle models.Parents
	result := database.Database.Where("username = ?", username).Delete(deletedArticle)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Parent has been deleted",
	})
}

func createParent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	result := database.Database.Create(&newParent)
	fmt.Println(result.Error)

	fmt.Println("Parent has been created")
	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Parent has been created",
	})
}

type RequestBody struct {
	Username string `json:"username"`
}

func checkAcc(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	var parent models.Parents
	result := database.Database.Where("username = ?", newParent.Username).First(&parent)
	if result.Error == nil {
		fmt.Println("Parent found")
		utils.JsonResponse(w, parent)
	} else {
		fmt.Println("Parent not found")
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}

func getParent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newParent models.Parents

	utils.JsonDeserialize(reqBody, &newParent)

	var parent models.Parents
	var result = database.Database.Where("username = ?", newParent.Username).Where("pass = ?", newParent.Pass).First(&parent)
	if result.Error == nil {
		fmt.Println("logging in...")
		utils.JsonResponse(w, parent)
	} else {
		fmt.Println("cannot login")
		utils.JsonResponse(w, models.BaseResult{
			Result:  false,
			Message: "Parent not found",
		})
	}
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)
	myrouter.HandleFunc("/", allParents).Methods("GET")
	myrouter.HandleFunc("/parent", getParent).Methods("POST")
	myrouter.HandleFunc("/parent/{id}", deleteParent).Methods("DELETE")
	myrouter.HandleFunc("/createacc", createParent).Methods("POST")
	myrouter.HandleFunc("/check", checkAcc).Methods("POST")
	handler := cors.Default().Handler(myrouter)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func main() {
	database.Testing()
	handleRequests()
	// Capture connection properties.
	/*var dsn = "host=localhost user=postgres password=4B-R05miyo dbname=kumon_database port=9292 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database_url := "postgres://postgres:4B-R05miyo@localhost:9292/kumon_database"
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	var parent parents
	var user = db.Take(&parent, "username", "maoadrian12")
	/*var username string
	var name string
	var pass string
	err = conn.QueryRow(context.Background(), "select username, name, pass from parents").Scan(&username, &name, &pass)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}*/
	//fmt.Println(username, name, pass)
	/*
		fmt.Printf("username: %s\n", parent.Username)
		fmt.Printf("%d rows effected\n", user.RowsAffected)*/
}
