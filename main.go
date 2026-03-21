package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/",http.FileServer(http.Dir(".")))
	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	err := http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		fmt.Println(err)
	}

}