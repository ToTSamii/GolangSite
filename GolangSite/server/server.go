package main

import (
	"html/template"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/index.html")

	if err != nil {
		log.Fatal("Not found index.html!")
	}

	temp.Execute(w, nil)
}

func HandleKb(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/kb.html")

	if err != nil {
		log.Fatal("Not found kb.html!")
	}

	temp.Execute(w, nil)
}

func HandleHaking(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/haking.html")

	if err != nil {
		log.Fatal("Not found haking.html!")
	}

	temp.Execute(w, nil)
}

func HandleAntiv(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/antiv.html")

	if err != nil {
		log.Fatal("Not found antiv.html!")
	}

	temp.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../img/"))))
	mux.HandleFunc("/Хакинг", HandleHaking)
	mux.HandleFunc("/Кибербезопасность", HandleKb)
	mux.HandleFunc("/Антивирусы", HandleAntiv)
	mux.HandleFunc("/", HandleIndex)

	log.Print("Server started!")

	if http.ListenAndServe(":9090", mux) != nil {
		log.Fatal("Server not started!!!")
	}
}
