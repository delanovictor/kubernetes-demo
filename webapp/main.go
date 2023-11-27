package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	database "webapp/database"
)

type RootData struct {
	PodIP   string
	PodName string
}

type MessageData struct {
	Messages []*database.Message
}

func main() {

	http.HandleFunc("/messages", handleMessages)

	http.HandleFunc("/", handleRoot)

	log.Println("Listening on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/index.html"))

	data := RootData{
		PodIP:   os.Getenv("MY_POD_IP"),
		PodName: os.Getenv("MY_POD_NAME"),
	}

	template.Execute(w, data)
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/messages.html"))

	switch r.Method {
	case "POST":

		r.ParseForm()

		message := r.Form.Get("message")

		fmt.Println("Add message: ", message)

		err := database.AddMessage(message)

		if err != nil {
			panic(err)
		}

		messages, err := database.ListMessages()

		if err != nil {
			panic(err)
		}

		template.Execute(w, MessageData{messages})

	case "GET":
		messages, err := database.ListMessages()

		if err != nil {
			panic(err)
		}

		template.Execute(w, MessageData{messages})
	}

}
