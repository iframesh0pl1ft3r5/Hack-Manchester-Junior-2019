package main

import (
	"net/http"
)

//OtherItemsInit -
func assets() {
	http.HandleFunc(GetServerIp(0)+"/", ConnectedByIP)
	// Making the assets folder work.
	// Location of local file
	fs := http.FileServer(http.Dir("./html/assets/"))
	//location on server when hosted
	http.Handle(MainSiteURL+"/assets/", http.StripPrefix("/assets/", fs))
}

func ConnectedByIP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+MainSiteURL+":"+SitePort+"/", 307)
}
