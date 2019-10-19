package main

import (
	"html/template"
	"net/http"
)

func initPages() {
	http.HandleFunc(MainSiteURL+"/", indexPage)
	assets()
}

//indexPage
func indexPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./html/index.html")
	t.Execute(w, r)
}
