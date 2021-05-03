package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Host: ", r.Host)

	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	var host string = "127.0.0.1"
	var port string = "8080"

	fmt.Println("Start Listening at http://" + host + ":" + port)
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
