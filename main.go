package main

import (
	"fmt"
	auth "login/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/home", auth.Home)
	http.HandleFunc("/logout", auth.Logout)
	fmt.Println("Server started at port : 8080")
	http.ListenAndServe(":8080", nil)
}
