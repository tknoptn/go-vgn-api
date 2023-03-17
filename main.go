package main

import (
	"encoding/json"

 
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
if err != nil{
	log.Fatal(err)
}
defer db.Close()

// what if table doen't exist , in that we need create the table

_,err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

if err !=nil{
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		next.ServeHTTP(w,r)

	})
	}


//  get all users function 	 
func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(users)
	}
}
// get the particular user you needed
func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){

		vars := mux.Vars(r)
		id := vars["id"]

		var u User
		err := db.QueryRow("SELECT * FROM users WHERE ID =$1",id).Scan(&u.ID,&u.Name,&u.Email)

		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}
}

// create one user 
func createUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var u User
		json.NewDecoder(r.Body).Decode(&u)

		err := db.QueryRow("INSERT INTO users(name, email) VALUES ($1,$2) RETURNING id",u.Name,u.Email).Scan(&u.ID)
		if err != nil{
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}


// To update partiuclat user details
func updateUser(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)

		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}

}

// To delete a particular user
func deleteUser(db *sql.DB)http.HandlerFunc  {
	return func(w http.ResponseWriter,r *http.Request){

		vars := mux.Vars(r)
		id := vars["id"]
		var u User
		err := db.QueryRow("SELECT * FROM users WHERE ID =$1",id).Scan(&u.ID,&u.Name,&u.Email)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}else {
			_,err := db.Exec("DELETE FROM users WHERE id = $1", id)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode("user deleted from list")
		}

		
	}

}


