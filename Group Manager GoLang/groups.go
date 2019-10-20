package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	helpers "./helpers"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	groupss template.HTML
}

func initGroups() {
	router.HandleFunc("/groups", groupIndex)
}

var inject string

var groupname string
var groupinfo string
var users int

func groupIndex(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "hack:HackManchester2019@tcp(10.0.2.58:3306)/group_mgr")
	if err != nil {
	}
	defer db.Close()
	var MaxID int
	var maxidstr string
	results, err := db.Query("SELECT MAX(GroupID) from qgroups")
	for results.Next() {
		var name Name
		// for each row, scan the result into our tag composite object
		err = results.Scan(&name.Name)
		if err != nil {
		} else {
			maxidstr = name.Name
			MaxID, _ = strconv.Atoi(maxidstr)
		}
	}
	for i := 1; i <= MaxID; i++ {
		formatted := fmt.Sprintf("Select GroupName from qgroups where GroupID = %d;", i)
		results, err := db.Query(formatted)
		for results.Next() {
			var name Name
			// for each row, scan the result into our tag composite object
			err = results.Scan(&name.Name)
			if err != nil {
			} else {
				groupname = name.Name
			}
		}
		formatted = fmt.Sprintf("Select GroupInfo from qgroups where GroupID = %d;", i)
		results, err = db.Query(formatted)
		for results.Next() {
			var name Name
			// for each row, scan the result into our tag composite object
			err = results.Scan(&name.Name)
			if err != nil {
			} else {
				groupinfo = name.Name
			}
		}
		formatted = fmt.Sprintf("Select Users from qgroups where GroupID = %d;", i)
		results, err = db.Query(formatted)
		for results.Next() {
			var name Nametwo
			// for each row, scan the result into our tag composite object
			err = results.Scan(&name.Nametwo)
			if err != nil {
			} else {
				users = name.Nametwo
			}
		}
		var groupsss string
		groupsss = "<div class=\"col-md-4 col-sm-6 item\"><div data-aos=\"fade\" class=\"box\" style=\"filter: brightness(100%) grayscale(0%) hue-rotate(88deg) invert(0%);background-image: url(&quot;assets/img/eGNkeoo.png&quot;);opacity: 1;\"><i class=\"fas fa-server icon\" style=\"color: rgb(255,255,255);\"></i><h3 class=\"name\" style=\"color: rgb(255,255,255);\">" + groupname + "</h3><p class=\"description\" style=\"color: rgb(186,186,186);\">" + groupinfo + ": Join today.</p><a class=\"learn-more\" role=\"button\" data-toggle=\"modal\" href=\"" + "#" + "\" style=\"color: rgb(251,251,251);\">Goto&gt;&gt;</a></div></div>"
		inject = inject + groupsss
	}
	//t, _ := template.ParseFiles("./templates/groups.html")
	//fmt.Print(template.HTML(inject))
	//t.Execute(w, template.HTML(inject))

	var body, _ = helpers.LoadFile("templates/groupspt1.html")
	var bodyt, _ = helpers.LoadFile("templates/groupspt2.html")
	body = body + inject + bodyt
	fmt.Fprintf(w, body)
}

type Nametwo struct {
	Nametwo int `json:"name"`
}
type Name struct {
	Name string `json:"name"`
}
