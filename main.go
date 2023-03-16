package main

import (

	"fmt"
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// define the data layout in struct
type User struct{
	ID int `json:"id"`
	Name string `json:"name`
	Email string `json:"email`
}


func main(){

	// giving connection to database
	db,err := sql.Open("postgres",os.Getenv("DB_URL"))
	if err!= nil{
		log.Fatal(err)
	}
	defer db.Close()

	// what if table doen't exist , in that we need create the table

	_,err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err! = nil {
		log.Fatal(err)
	}

	// define the route paths

	router := mux.NewRouter()
	router.HandleFunc("/users",getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}",getUser(db)).Methods("GET")
	router.HandleFunc("/users",createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}",updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}",deleteUser(db)).Methods("DELETE")

	// run the server

	log.Fatal(http.ListenAndServe(":8080",jsonMiddleware(router)))

}
     func jsonMiddleware(next http.Handler) http.Handler  {
		return http.HandleFunc(func(w http.ResponseWriter,r *http.Request){
			w.Header().Set("Content-Type","application/json")
			next.ServeHTTP(w,r)

		})
	 }


//  get all users function 	 
func getUsers()  {
	
}

// get the particular user you needed
func getUser()  {
	
}

// create some user 
func createUser(){

}

// To update partiuclat user details
func updateUser(){

}

// To delete a particular user
func deleteUser()  {
	
}

//