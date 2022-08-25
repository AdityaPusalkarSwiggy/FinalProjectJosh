package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Server!\n")
	fmt.Fprint(w, "Visit '127.0.0.1:3000/website' to send POST requests for adding websites to the status watchlist\n")
	fmt.Fprint(w, "Visit '127.0.0.1:3000/website' to send GET requests for viewing the status of all websites in the watchlist\n")
	fmt.Fprint(w, "Visit '127.0.0.1:3000/website?name={website}' to send GET requests for viewing the status of a particular website in the watchlist\n")
}

func HandlerWebsite(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s := Websites{}
		fmt.Println("\nIncoming POST Request!")
		json.NewDecoder(r.Body).Decode(&s)
		for _, website := range s.Websites {
			_, exists := websiteStatus[website]
			if !exists {
				websiteStatus[website] = "UNKNOWN"
			}
		}
	case "GET":
		fmt.Println("\nIncoming GET Request")
		query := r.URL.Query()
		if qval, exists := query["name"]; exists {
			fmt.Println("Website:", qval[0])
			if status, exists := websiteStatus[qval[0]]; exists {
				json.NewEncoder(w).Encode(map[string]string{qval[0]: status})
			} else {
				json.NewEncoder(w).Encode(map[string]string{qval[0]: "UNKNOWN"})
			}
		} else {
			json.NewEncoder(w).Encode(websiteStatus)
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported\n")
	}
}
