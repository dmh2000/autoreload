/*
Package autoreload provides a simple way to reload a webpage when a file changes. In conjunction
with the github.com/cosmtrek/air@latest tool, this package will automatically reload the webpage
when air detects a file change. This is useful for web development when you want to see the changes
immediately without having to manually refresh the page.
*/
package autoreload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader is a websocket upgrader that is used to upgrade an http connection to a websocket connection.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// Function server listens for a websocket connection on the specified path and port.
// The origin parameter is used to check the origin of the webserver websocket connection.
// If the origin does not match, the connection is refused.
func server(path string, port int, origin string) {
	handler := func(w http.ResponseWriter, r *http.Request) {

		// only allow connection to specified origin
		upgrader.CheckOrigin = func(r *http.Request) bool { 
			clientOrigin := r.Header.Get("origin")
			return clientOrigin == origin
		}
	
		// upgrade to websocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		
		// listen for loss of connection
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}

	// run the server at the specified url and port
	// must match the url:port in the javascript snippet
	http.HandleFunc(path, handler)
	http.ListenAndServe(fmt.Sprintf(":%d",port), nil)
}

// minified websocket reloader script
func Handler(w http.ResponseWriter, r *http.Request) {
	src := `{
		let active = false;
		// set the URL of the websocket server to the host where the go app is running
		sock = new WebSocket("ws://localhost:8080/reload");
		sock.onopen = function (event) {
		  // console.log("connected");
		  active = true;
		};
	  
		sock.onclose = function (event) {
		  // console.log("disconnected");
		  // the timeout value needs to be long enough for the
		  // go app to reload before refreshing this page.
		  // tune it to what works on your system.
		  if (active) {
			setTimeout(function () {
		      location.reload();
			  active = false;
			}, 2000);
		  }
		};
	  }`
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(src))
}

// Start starts the autoreload server as a goroutine and listens for websocket connections
// on the specified url and port. The origin parameter is used to check the origin of the
// webserver websocket connection. If the origin does not match, the connection is refused.
func Start(url string, port int,origin string) {
	go server(url,port, origin)
}