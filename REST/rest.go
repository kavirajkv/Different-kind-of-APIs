package rest

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

var users = map[string]string{}

func createuser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users[user.Username] = user.Email

	msg := Response{Message: "User created"}
	json.NewEncoder(w).Encode(&msg)

}

func getuser(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("name")

	if user == "" {
		http.Error(w, "give username", http.StatusBadRequest)
		return
	}

	var founduser User
	if email, exist := users[user]; exist {
		founduser.Username = user
		founduser.Email = email

	} else {
		msg := Response{Message: "user not found"}
		json.NewEncoder(w).Encode(&msg)
		return
	}

	json.NewEncoder(w).Encode(&founduser)
}

func Controller() {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/createuser", createuser)
	r.HandleFunc("GET /api/getuser/{name}", getuser)

	println("server started at port 8000")
	http.ListenAndServe(":8000", r)
}
