package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

var messages []Message

type Message struct {
	From string `json:"from"`
	Body string `json:"body"`
}

func main() {
	http.HandleFunc("/messages", messagesHandler)
	http.ListenAndServe(":8080", nil)
}

//todo add the ENV variable for the user's name. Default to Docker

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
		t, _ := template.ParseFiles("messages.html")
		t.Execute(w, map[string]interface{}{"Messages": messages})
	} else {
		// youre a bad boy
	}
}
