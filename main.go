package main

import (
	"CodeColony/internal"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// we got a map that is build as a coordinate system.

// we got a main building that can produce creeps

// we also got creeps that can move, collect and carry
// there should also be natural structures that the creep has to move around
// also there should definitly be more than one creep

func main() {
	go runGame()
	go runHttpListener()
	select {}
}
func runGame() {
	g := internal.NewGame(40, 30)
	g.Tick()
}

func runHttpListener() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	http.Handle("/", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>htmx Example</title>
		<script src="https://cdn.jsdelivr.net/npm/htmx.org@1.7.1/dist/htmx.min.js"></script>
	</head>
	<body>
		<h1>htmx Example</h1>
		<button hx-get="/data" hx-trigger="click" hx-target="#result">Load Data</button>
		<div id="result"></div>
	</body>
	</html>
	`
	t, _ := template.New("index").Parse(tmpl)
	t.Execute(w, nil)
}
