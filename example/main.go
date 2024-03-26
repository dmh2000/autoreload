// This Go program is a simple web server that serves two page sand uses github.com/dmh2000/autoreload
// to enable live reloading. Live reloading is a feature that automatically refreshes
// the web page whenever there are changes in the source code.
// The main program must be run with the air tool, which is a file watcher that automatically restarts
// the web server when it detects changes in the source code.
package main

import (
	"fmt"
	"html/template"
	"net/http"

	autoreload "github.com/dmh2000/reload"
)

type Data struct {
	PageTitle string
}

var index = template.Must(template.ParseFiles("views/index.html"))
var about = template.Must(template.ParseFiles("views/about.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	data := Data{PageTitle: "Live Reload"}
	index.Execute(w, data)
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About")
	data := Data{PageTitle: "About Live Reload"}
	about.Execute(w, data)
}

func main() {
	// --------------------------------------------------------------------------------------------
	// ADD THE FOLLOWING CODE TO ENABLE LIVE RELOADING
	fmt.Println("Starting autoreload server")
	const reloadPort = 8080                 // port that the reload websocket server listens on
	const reloadUrl = "/reload"             // url that the reload websocket server listens on
	const origin =  "http://localhost:8001" // used for origin check in websocket upgrade
	// start the autoreload server which will spawn a goroutine that listens for websocket connections
	autoreload.Start(reloadUrl,reloadPort, origin)
	// add autoreload handler to the mux. url should match the snippet in the html file
	http.HandleFunc("/reload/reloader.js", autoreload.Handler)
	// --------------------------------------------------------------------------------------------

	// add the main application handlers
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	const serverPort = ":8001"				// local server port of the main application
	fmt.Printf("Server Listening on port %s\n", serverPort)
	err := http.ListenAndServe(serverPort,nil)
	if err != nil {
		panic(err)
	}
}
