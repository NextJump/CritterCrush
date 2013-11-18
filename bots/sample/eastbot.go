package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"nxj/critter/engine"
	"time"
)

func hangOut() engine.Action {
	return engine.Action{"EAST"}
}

// ==============
// Standard bot listening stuff.. do not chance
// ===============
func getAction() engine.Action {
	actionFn := hangOut

	choice := actionFn()
	log.Printf("Action %s", choice.Action)
	return choice
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(getAction())
	fmt.Fprintf(w, string(json))
}

func main() {
	http.HandleFunc("/action", actionHandler)
	fmt.Println("listening 8084")
	http.ListenAndServe(":8084", nil)
}

func init() {
	rand.Seed(time.Now().Unix())
}
