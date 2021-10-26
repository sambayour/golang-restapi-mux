package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//User Struct /model
type User struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Parent *Parent `json:"parent"`
}

// Spouse Struct
type Parent struct {
	Father string `json:"father"`
	Mother string `json:"mother"`
}

var users []User

//get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(users)
}

//get user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)

	for _, us := range users {

		if us.ID == params["id"] {
			json.NewEncoder(w).Encode(us)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

//create user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(10000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

//update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, us := range users {
		if us.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = strconv.Itoa(rand.Intn(10000000))
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

//delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, us := range users {
		if us.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {

	//init routes
	r := mux.NewRouter()

	//mock data

	users = append(users, User{ID: "1", Name: "Samuel Olubayo", Age: 27,
		Parent: &Parent{Father: "Taiwo Olubayo", Mother: "Elizabeth Olubayo"}})

	users = append(users, User{ID: "2", Name: "Abiola Agbeti", Age: 27,
		Parent: &Parent{Father: "Raimi", Mother: "Adebola"}})

	//routes endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")

	//serve application
	log.Fatal(http.ListenAndServe(":8080", r))
}
