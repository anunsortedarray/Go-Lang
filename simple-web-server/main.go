package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

// func serveStaticIndexFile(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "static")
// }

func echoPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	// http.HandleFunc("/", serveStaticIndexFile)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/echo", echoPath)
	http.HandleFunc("/increment", incrementCounter)

	log.Fatal(http.ListenAndServe(":1000", nil))
}
