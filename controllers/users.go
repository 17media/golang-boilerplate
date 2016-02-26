package controllers

import (
        "github.com/17media/golang-boilerplate/models/users"
)

// listUser - User listing
func listUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprint(w, "User List!\n")
}

// createUser - Create a new user
func createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprint(w, "User Created!\n")
}

// getUser - Get User by ID
func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprint(w, "Get User!\n")
}

// updateUser - Update User by ID
func updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprint(w, "Update User!\n")
}

// deleteUser - Delete User by ID
func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprint(w, "Delete User!\n")
}
