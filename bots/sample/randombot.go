package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"nxj/critter/engine"
	"time"
)

// Action 3: Arrays Example
// BOTS - RANDOM!  Dogs and cats in the streets
func goRandom() engine.Action {
	actions := [4]string{NORTH, SOUTH, EAST, WEST}
	dir := rand.Intn(4)
	action := actions[dir]
	choice := engine.Action{action}
	return choice
}

//==========================================
//
// Standard bot code -- do not touch
//
//=========================================

// consts to match the actions our bot can take
const EAST string = "EAST"
const WEST string = "WEST"
const NORTH string = "NORTH"
const SOUTH string = "SOUTH"
const CRUSH string = "CRUSH"
const STAY string = "STAY"

func init() {
	rand.Seed(time.Now().Unix())
}

func getAction() engine.Action {
	actionFn := goRandom

	choice := actionFn()
	log.Printf("Action %s", choice.Action)
	return choice
}
func actionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(getAction())
	fmt.Fprintf(w, string(json))

	//	time.Sleep(time.Second * 1)
}

func main() {
	http.HandleFunc("/action", actionHandler)
	fmt.Println("listening 8086")
	http.ListenAndServe(":8086", nil)
}

func getBoardFromServer() *engine.Board {
	response, _ := http.Get("http://localhost:8080/board")
	body, _ := ioutil.ReadAll(response.Body)

	var b = new(engine.Board)
	json.Unmarshal(body, &b)
	return b
}
