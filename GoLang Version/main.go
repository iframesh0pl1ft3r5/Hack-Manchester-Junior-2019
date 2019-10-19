package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//MySQL - used for storing MySQL Username and password
type MySQL struct {
	Username string
	password string
}

// MainSiteURL - defines the site url
var MainSiteURL string
var SitePort string
var NonHttpsPort string
var MySQLUsername string
var MySQLPassword string

func main() {
	MainSiteURL = "localhost"
	SitePort = "443"
	NonHttpsPort = "8080"
	MySQLUsername = "HACK"
	MySQLPassword = "HackManchester2019"
	initPages()
	fmt.Println("The server ip is: " + GetServerIp(0))
	//openbrowser("https://" + MainSiteURL + ":" + SitePort + "/index.html")
	// http.ListenAndServe(":"+NonHttpsPort, nil)
	//  Start HTTP
	go func() {
		err_http := http.ListenAndServe(":"+NonHttpsPort, nil)
		if err_http != nil {
			log.Fatal("Web server (HTTP): ", err_http)
		}
	}()
	err_https := http.ListenAndServeTLS(":"+SitePort, "./server.crt", "./server.key", nil)
	if err_https != nil {
		log.Fatal("Web server (HTTPS): \n", err_https)
	}
}

func GetServerIp(ipNum int) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	var ips [255]string
	var i int
	i = 0
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips[i] = ipnet.IP.String()
				i = i + 1
			}
		}
	}
	return ips[ipNum]
}

// allows me to ouput preforamtted date and time
func dateTimeFormatted() string {
	// get Date time.
	t := time.Now()
	currTime := time.Now()
	date := strings.Split(currTime.String(), " ")[0]
	splits := strings.Split(date, "-")
	year := splits[0]
	month := splits[1]
	day := splits[2]
	h, m, s := t.Clock()
	hour := strconv.Itoa(h)
	min := strconv.Itoa(m)
	sec := strconv.Itoa(s)
	return hour + ":" + min + ":" + sec + " " + day + "/" + month + "/" + year
}
