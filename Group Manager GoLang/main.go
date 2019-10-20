package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	common "./common"
	helpers "./helpers"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", common.LoginPageHandler) // GET
	router.HandleFunc("/about", abouthtml)
	router.HandleFunc("/index", common.IndexPageHandler) // GET
	router.HandleFunc("/index.html", abouthtml)          // GET
	router.HandleFunc("/login", common.LoginHandler).Methods("POST")
	router.HandleFunc("/login.html", loginredirect) // GET
	router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", common.RegisterHandler).Methods("POST")

	router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")

	// Making the assets folder work.
	// Location of local file
	fs := http.FileServer(http.Dir("./assets/"))
	//location on server when hosted
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}

func abouthtml(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("templates/about.html")
	fmt.Fprintf(response, body)
}

func loginredirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}
