package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/dmh2000/reload"
)

type Data struct {
	PageTitle string
}

var tmpl = template.Must(template.ParseFiles("views/hello.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	data := Data{PageTitle: "Hello Live Reload"}
	tmpl.Execute(w, data)
}


func main() {
	const reloadPort = 8080                 // port that the reload server listens on
	const reloadUrl = "/reload"             // url that the reload server listens on
	const origin =  "http://localhost:8001" // used for origin check in websocket upgrade
	const serverPort = ":8001"					// local server port

	reload.Reload(reloadUrl,reloadPort, origin)

	http.HandleFunc("/", Home)

	fmt.Printf("Server Listening on port %s\n", serverPort)
	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		panic(err)
	}
}
