package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"nxj/critter/engine"
	"strconv"
	"time"
)

var playerNum int = 0

// Action 4: Arrays Example (continued)
// BOTS - Random! Please avoid walls :)
func goRandomCrush() engine.Action {
	board := getBoardFromServer()
	player := board.GetPlayer(playerNum)

	if isEnemy(board, player.X, player.Y) {
		return engine.Action{CRUSH}
	}

	var choice engine.Action
	actions := [4]string{NORTH, SOUTH, EAST, WEST}
	for {
		roll := rand.Intn(4)
		choice = engine.Action{actions[roll]}
		x, y := getCoords(player, choice)
		if isWall(board, x, y) == false {
			break
		}
		fmt.Println("Avoided a wall hit!")
	}
	return choice
}

func getCoords(player *engine.Player, dir engine.Action) (x int, y int) {
	x = player.X
	y = player.Y

	switch dir.Action {
	case EAST:
		x++
	case WEST:
		x--
	case NORTH:
		y--
	case SOUTH:
		y++
	}
	return
}

func isWall(board *engine.Board, x int, y int) bool {
	if x < 0 || y < 0 {
		return true
	}
	if x > (board.Width - 1) {
		return true
	}
	if y > (board.Height - 1) {
		return true
	}

	for _, env := range board.Environments {
		if env.Id != 0 {
			continue
		}
		if env.X == x && env.Y == y {
			fmt.Println("inner wall hit avoided")
			return true
		}
	}
	return false
}

func isEnemy(board *engine.Board, x int, y int) bool {
	for _, e := range board.Enemies {
		if e.IsCrushed {
			continue
		}
		if e.X == x && e.Y == y {
			return true
		}
	}
	return false
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
	actionFn := goRandomCrush

	choice := actionFn()
	log.Printf("Action %s", choice.Action)
	return choice
}
func actionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	playerNum, _ = strconv.Atoi(r.FormValue("playernum"))
	json, _ := json.Marshal(getAction())
	fmt.Fprintf(w, string(json))

	//	time.Sleep(time.Second * 1)
}

func main() {
	http.HandleFunc("/action", actionHandler)
	fmt.Println("listening 8087")
	http.ListenAndServe(":8087", nil)
}

func getBoardFromServer() *engine.Board {
	response, _ := http.Get("http://localhost:8080/board")
	body, _ := ioutil.ReadAll(response.Body)

	var b = new(engine.Board)
	json.Unmarshal(body, &b)
	return b
}
