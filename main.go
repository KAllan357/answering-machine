package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var messages []Message

type Message struct {
	From string `json:"from"`
	Body string `json:"body"`
}

// Credit to this! https://groups.google.com/forum/#!topic/golang-nuts/s7Xk1q0LSU0
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/messages", messagesHandler)
	http.ListenAndServe(":8080", Log(http.DefaultServeMux))
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	if method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		var message Message
		if err == nil {
			json.Unmarshal(body, &message)
		}

		// 201
		messages = append(messages, message)
		fmt.Fprint(w, message)
	} else if method == "GET" {
		// view some messages
		name := os.Getenv("ANSWERING_MACHINE_OWNER")
		if name == "" {
			name = "Docker"
		}

		t, _ := template.ParseFiles("messages.html")
		t.Execute(w, map[string]interface{}{"Messages": messages, "Name": name})
	} else {
		// youre a bad boy
	}
}
