package main

import (
	"fmt"
	"net/http"
)

type Websites struct {
	Websites []string `json:"websites"`
}

var websiteStatus = make(map[string]string)

func main() {
	fmt.Println("Starting Server...")
	go UpdateWebsiteStatus()
	http.HandleFunc("/", HandlerDefault)
	http.HandleFunc("/website", HandlerWebsite)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
