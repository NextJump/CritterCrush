package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"devpoets/critter/engine"
	"strconv"
)

var game *engine.GameEngine

func init() {
	game = engine.NewGameEngine(8081, 8082)
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(game.Board)
	fmt.Fprintf(w, string(json))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	game.Play()
	fmt.Fprintf(w, `{"status":"OK"}`)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	j, err := json.Marshal(game.Board)
    if err != nil {
      fmt.Println("Error marshalling game board: " + err.Error())
    }

	t, err := template.ParseFiles("src/devpoets/critter/ui/index.html")
    if err != nil {
      fmt.Println("Error parsing template: " + err.Error())
    }
	t.Execute(w, template.JS(string(j)))
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	game.Reset()
	fmt.Fprintf(w, `{"status":"OK"}`)
}

func setportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	port1, port2 := getPortsFromRequest(r)
	game.Ports = []int{port1, port2}

	fmt.Fprintf(w, `{"status":"OK"}`)
}
func getPortsFromRequest(r *http.Request) (int, int) {
	player1, _ := strconv.Atoi(r.FormValue("player1"))
	player2, _ := strconv.Atoi(r.FormValue("player2"))
	return player1, player2
}

func main() {
	fmt.Println("ready player one...")

	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/setport", setportHandler)

    http.Handle("/", http.FileServer(http.Dir("src/devpoets/critter/ui/assets")))

	http.ListenAndServe(":8080", nil)
}
