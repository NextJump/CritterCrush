package engine

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type GameEngine struct {
    Turns     int
    Board     *Board
    isRunning bool
    control   chan string
    Ports     []int
}

type environment struct {
    Id int
    X  int
    Y  int
}

type Action struct {
    Action string
}

func NewGameEngine(port1 int, port2 int) *GameEngine {
    e := new(GameEngine)
    e.control = make(chan string)
    e.Board = new(Board)
    e.Board.setup()
    e.Ports = []int{port1, port2}
    e.isRunning = false
    return e
}

func (e *GameEngine) Play() {
    if e.isRunning == false {
        e.isRunning = true
        e.Turns = 0
        e.Turn()
        return
    }
    e.Stop()
}

func (e *GameEngine) Turn() {
    if e.checkRun() == false {
        return
    }
    e.Turns++
    e.movePlayers()

    if e.Turns < 110 {
        time.AfterFunc(250*time.Millisecond, e.Turn)
    } else {
        e.isRunning = false
    }
}

func (e *GameEngine) Reset() {
    e.Stop()
    e.Board.setup()
    e.Play()
}

func (e *GameEngine) Stop() {
    if e.isRunning == false {
        return
    }
    e.control <- "Stop"
    e.isRunning = false
}

func (e *GameEngine) checkRun() bool {
    if e.isRunning == false {
        return false
    }
    select {
    case <-e.control:
        return false
    default:
    }
    return e.isRunning
}

func (e *GameEngine) movePlayers() {
    for i, p := range e.Board.Players {
        url := fmt.Sprintf(p.Url, e.Ports[i], i)
        response, _ := http.Get(url)
        fmt.Println(url)
        body, _ := ioutil.ReadAll(response.Body)

        var a Action
        json.Unmarshal(body, &a)

        e.takeAction(e.Board.Players[i], a.Action)
    }
}

func (e *GameEngine) takeAction(p *Player, action string) {
    switch action {
    case "NORTH":
        if isWall(e.Board, p.X, p.Y-1) {
            return
        }
        p.Y--
    case "SOUTH":
        if isWall(e.Board, p.X, p.Y+1) {
            return
        }
        p.Y++
    case "EAST":
        if isWall(e.Board, p.X+1, p.Y) {
            return
        }
        p.X++
    case "WEST":
        if isWall(e.Board, p.X-1, p.Y) {
            return
        }
        p.X--
    case "CRUSH":
        if c := getEnemy(e.Board, p.X, p.Y); c != nil {
            c.IsCrushed = true
            p.Score += c.Score
        }
    }
}
