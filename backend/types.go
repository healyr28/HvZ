package main

import (
    "fmt"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
)

type State int

const (
    CoreZombie State = iota
    Zombie
    StunnedCoreZombie
    StunnedZombie
    InfectedHuman
    Human
    Dead
)

type Player struct {
    Id             string    `json:"id"`
    Name           string    `json:"name"`
    password       string
    State          State     `json:"state"`
    Last_tagged    time.Time `json:"last_tagged"`
    Kills          int       `json:"kills"`
    Last_kill      time.Time `json:"last_kill"`
    Cures          int       `json:"cures"`
    Revives        int       `json:"revives"`
    Extensions     int       `json:"extensions"`
}

type Login struct {
    Name string `json:"name"`
    Pass string `json:"password"`
}

func gen_pass(password string) (string) {
    password_bytes := []byte(password)

    hashed_bytes, _ := bcrypt.GenerateFromPassword(password_bytes, bcrypt.MinCost)

    return string(hashed_bytes)
}

func new_player(name string, password string, state State) (Player) {
    new_player := Player{
        Id:          uuid.New().String()[0:7],
        Name:        name,
        password:    gen_pass(password),
        State:       state,
        Last_tagged: time.Now(),
        Kills:       0,
        Last_kill:   time.Now(),
        Cures:       3,
        Revives:     3,
        Extensions:  3,
    }
    return new_player
}

func (p *Player) login (password string) (bool) {
    succ := bcrypt.CompareHashAndPassword([]byte(p.password), []byte(password))
    if (succ == nil) {
        return true
    } else {
        fmt.Println(succ)
        return false
    }
}
