package reload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

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

func Reload(url string, port int,origin string) {
	fmt.Println("Reload",url,port,origin)
	go server(url,port, origin)
	
}