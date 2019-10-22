// websockets.go
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
	// Create a simple file server
    fs := http.FileServer(http.Dir("../public"))
    http.Handle("/", fs)
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            fmt.Printf(string (msg))
            if err != nil {
                return
            }

            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
            msg = []byte("Here is a string....")
            // Write message back to browser
            for i := 0; i < 50; i++ {
                time.Sleep(time.Second * 1)
                t := time.Now()
                //fmt.Println(t.Format("20060102150405"))
                msg = []byte(t.Format("20060102150405"))

                if err = conn.WriteMessage(msgType, msg); err != nil {
                    return
                }            }
            msg = []byte("Last string....")
            conn.WriteMessage(msgType, msg)
            //if err = conn.WriteMessage(msgType, msg); err != nil {
            //    return
            //}
        }
    })
    http.ListenAndServe(":8080", nil)
}
