package main

import (
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID string `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range users {
		if v.ID == params["id"]{
			json.NewEncoder(w).Encode(v)
			return
		} 
	}
	json.NewEncoder(w).Encode(&User{})
}

func postUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(&user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, v := range users {
		if v.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)

			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}

	}
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, v := range users {
		if v.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {

	router := mux.NewRouter()

	users = append(users, User{ID: "1", FirstName: "Bukola", LastName: "Odunayo", UserName: "bukkyo", Email: "bukkyo@gmail.com", Password: "bukkyzzz", Age: 27})
	users = append(users, User{ID: "2", FirstName: "Lolu", LastName: "Demola", UserName: "lolu99", Email: "lolu@gmail.com", Password: "loluoo", Age: 28})
	users = append(users, User{ID: "3", FirstName: "Seun", LastName: "Tobiloba", UserName: "Sewenzy", Email: "seun@gmail.com", Password: "chessmadam", Age: 30})
	users = append(users, User{ID: "4", FirstName: "Kunle", LastName: "Emmanuel", UserName: "kk", Email: "kunlekk@gmail.com", Password: "myname", Age: 25})

	router.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", postUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", updateUser).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":3050", router))
};