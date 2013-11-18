package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"cw/critter/engine"
	"time"
)

// traversal direction
var direction string = EAST

// Action 1: Variable example // Just move east...
func goEast() engine.Action {
	action := engine.Action{EAST}
	return action
}

// Action 2: Branching Example
// BOTS - Move either east or west with each turn
func getEastOrWest() engine.Action {
	if rand.Intn(2) == 0 {
		return engine.Action{EAST}
	}
	return engine.Action{WEST}
}

// Action 3: Arrays Example
// BOTS - RANDOM!  Dogs and cats in the streets
func goRandom() engine.Action {
	actions := [4]string{NORTH, SOUTH, EAST, WEST}

	choice := engine.Action{actions[rand.Intn(4)]}
	return choice
}

// Action 4: Arrays Example (continued)
// BOTS - Random! Please avoid walls :)
func goRandomNoWallHits() engine.Action {
	board := getBoardFromServer()
	player := board.GetPlayer(0)

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

// Action 5: Simpler version (use functions)
// BOTS - Random! Please avoid walls :)
func goRandomNoWallHitsSimple() engine.Action {
	board := getBoardFromServer()
	player := board.GetPlayer(0)

	var choice engine.Action
	for {
		choice = getChoice(rand.Intn(4))

		x, y := getCoords(player, choice)
		if isWall(board, x, y) == false {
			break
		}
	}
	return choice
}

// Action 6: Continuing down the line
// BOTS - traverse horizontally until wall hit
func goTraverseSimple() engine.Action {
	board := getBoardFromServer()
	player := board.GetPlayer(0)

	// use a map to converting strings
	//   to numerics required by getChoice
	dirMap := map[string]int{
		NORTH: 0,
		SOUTH: 1,
		EAST:  2,
		WEST:  3,
	}

	choice := getChoice(dirMap[direction])
	x, y := getCoords(player, choice)
	if isWall(board, x, y) == false {
		return choice
	}

	// we know our new direction for next turn must change
	if direction == EAST {
		direction = WEST
	} else {
		direction = EAST
	}

	// we hit a wall.. so let's try south
	choice = getChoice(dirMap[SOUTH])
	x, y = getCoords(player, choice)
	if isWall(board, x, y) == false {
		return choice
	}

	choice = getChoice(dirMap[direction])
	return choice
}

func goRandomDestroy() engine.Action {
	// if we are on top of a mob, HULK SMASH!!!
	board := getBoardFromServer()
	if board == nil {
		fmt.Println("No board?!")
	}
	actions := [4]string{NORTH, SOUTH, EAST, WEST}

	choice := engine.Action{actions[rand.Intn(4)]}
	return choice
}

func getChoice(index int) engine.Action {
	actions := [4]string{NORTH, SOUTH, EAST, WEST}
	return engine.Action{actions[index]}
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
	actionFn := goTraverseSimple

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
	fmt.Println("listening 8081")
	http.ListenAndServe(":8081", nil)
}

func getBoardFromServer() *engine.Board {
	response, _ := http.Get("http://localhost:8080/board")
	body, _ := ioutil.ReadAll(response.Body)

	var b = new(engine.Board)
	json.Unmarshal(body, &b)
	return b
}
