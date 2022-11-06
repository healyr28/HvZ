package main

import (
    "database/sql"
    "encoding/json"
    //"errors"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "github.com/swaggo/http-swagger"

    _ "github.com/mattn/go-sqlite3"
    _ "games.dcu.ie/hvz/v2/docs"
)

var store = sessions.NewCookieStore([]byte("wuhcwe789cui"))
var db *sql.DB

func get_all_players() []Player {
    rows, _ := db.Query("SELECT id FROM players")
    players := []Player{}
    var t_player Player
    var id string
    for rows.Next() {
        rows.Scan(&id)
        t_player, _ = get_player(id)
        players = append(players, t_player)
    }
    return players
}

func update() {
    for {
        players := get_all_players()
        for _, player := range players {
            if (player.State == InfectedHuman &&
            time.Since(player.Last_tagged).Seconds() >= 30) {
                player.State = Zombie
                player.Last_kill = time.Now()
            }
            if (player.State == StunnedZombie &&
            time.Since(player.Last_tagged).Seconds() >= 15) {
                player.State = Zombie
                player.Last_kill = time.Now()
            }
            if (player.State == StunnedCoreZombie &&
            time.Since(player.Last_tagged).Seconds() >= 10) {
                player.State = CoreZombie
                player.Last_kill = time.Now()
            }
            if (player.State == Zombie &&
            time.Since(player.Last_kill).Seconds() >= 24) {
                if (player.Extensions == 0) {
                    player.State = Dead
                } else {
                    player.Extensions -= 1
                }
            }
            save_player(player)
            time.Sleep(5 * time.Second)
        }
    }
}

func get_player(id string) (Player, error) {
    tmp_player := Player{}
    var last_tagged string
    var last_kill string
    time_format := time.RFC3339
    row := db.QueryRow("SELECT * FROM players WHERE id = ?", id)
    err := row.Scan(&tmp_player.Id, &tmp_player.Name, &tmp_player.password, &tmp_player.State, &last_tagged, &tmp_player.Kills, &last_kill, &tmp_player.Cures, &tmp_player.Revives, &tmp_player.Extensions)
    tmp_player.Last_tagged, _ = time.Parse(time_format, last_tagged)
    tmp_player.Last_kill, _ = time.Parse(time_format, last_kill)
    if err != nil {
        return Player{}, err
    }
    return tmp_player, nil
}

func save_player(player Player) {
    tmp_id := ""
    row := db.QueryRow("SELECT id FROM players WHERE id = ?", player.Id)
    err := row.Scan(&tmp_id)
    var statement *sql.Stmt
    if err != nil {
        statement, _ = db.Prepare("INSERT INTO players (id, name, password, state, last_tagged, kills, last_kill, cures, revives, extensions) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
        last_tagged, _ := player.Last_tagged.MarshalText()
        last_kill, _ := player.Last_kill.MarshalText()
        statement.Exec(player.Id, player.Name, player.password, player.State, last_tagged, player.Kills, last_kill, player.Cures, player.Revives, player.Extensions) 
    } else {
        statement, err = db.Prepare("UPDATE players SET id = ?, name = ?, password = ?, state = ?, last_tagged = ?, kills = ?, last_kill = ?, cures = ?, revives = ?, extensions = ? WHERE id = ?")
        last_tagged, _ := player.Last_tagged.MarshalText()
        last_kill, _ := player.Last_kill.MarshalText()
        statement.Exec(player.Id, player.Name, player.password, player.State, last_tagged, player.Kills, last_kill, player.Cures, player.Revives, player.Extensions, player.Id) 
    }
}

// @title DCU Games Society Humans Vs Zombies
// @version 1.0
// @description This is the Swagger documentation for DCU Games Society's 2022 Humans Vs Zombies event.

// @host
// @BasePath /
func handleRequests() {
    router := mux.NewRouter().StrictSlash(false)
    router.Path("/me").Methods(http.MethodGet).HandlerFunc(me)
    router.Path("/login").Methods(http.MethodPost).HandlerFunc(login)
    router.Path("/tag/{target}").Methods(http.MethodGet).HandlerFunc(tag)
    router.Path("/kill/{target}").Methods(http.MethodGet).HandlerFunc(kill)
    router.Path("/cure/{target}").Methods(http.MethodGet).HandlerFunc(cure)
    router.Path("/revive/{target}").Methods(http.MethodGet).HandlerFunc(revive)
    router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

    log.Fatal(http.ListenAndServe(":8000", router))
}

// @Summary Login
// @Description Login
// @Produce json
// @Accept json
// @Param info body Login true "Player login"
// @Success 200 {object} Player
// @Failure 400
// @Failure 403
// @Router /login [post]
func login(w http.ResponseWriter, r *http.Request) {
    l := Login{}

    if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
        fmt.Println(err)
        http.Error(w, "Error decoding response object", http.StatusBadRequest)
        return
    }

    players := get_all_players()
    for _, player := range players {
        if (player.Name == l.Name && player.login(l.Pass)) {
            session, _ := store.Get(r, "session")
            session.Values["userID"] = player.Id
            session.Save(r, w)
            http.Redirect(w, r, "/me", http.StatusFound)
            return
        }
    }
    http.Error(w, "Username or password incorrect", http.StatusForbidden)
}

// @Summary Tag
// @Description Tag a human as a zombie. Gives target the infected state if user is a zombie or core zombie.
// @Produce plain
// @Param target path string true "Target ID"
// @Success 200
// @Failure 401 {string} string "User is not logged in"
// @Failure 403 {string} string "Tagger is not a zombie"
// @Failure 404 {string} string "Target does not exist"
// @Failure 409 {string} string "Target is not human"
// @Router /tag [get]
func tag(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    target_id := vars["target"]

    sess, _ := store.Get(r, "session")
    id, ok := sess.Values["userID"]
    if !ok {
        http.Error(w, "User not logged in", http.StatusUnauthorized)
        return
    }

    player, _ := get_player(id.(string))
    if (player.State != CoreZombie && player.State != Zombie) {
        http.Error(w, "Tagger is not a zombie", http.StatusForbidden)
        return
    }

    target, err := get_player(target_id)
    if (err != nil) {
        http.Error(w, "Target not found", http.StatusNotFound)
        return
    }
    if (target.State != Human) {
        http.Error(w, "Target not human", http.StatusConflict)
        return
    }
    target.State = InfectedHuman
    target.Last_tagged = time.Now()
    player.Last_kill = time.Now()
    player.Kills += 1
    save_player(target)
    save_player(player)
    fmt.Fprintf(w, "Successfully tagged target")
}

// @Summary Kill
// @Description Kills a zombie if user is authenticated as a human. Gives zombie target the stunned state
// @Produce plain
// @Param target path string true "Target ID"
// @Success 200
// @Failure 401 {string} string "User is not logged in"
// @Failure 403 {string} string "Killer is not a human"
// @Failure 404 {string} string "Target does not exist"
// @Failure 409 {string} string "Target is not a zombie"
// @Router /kill [get]
func kill(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    target_id := vars["target"]

    sess, _ := store.Get(r, "session")
    id, ok := sess.Values["userID"]
    if !ok {
        http.Error(w, "User not logged in", http.StatusUnauthorized)
        return
    }

    tagger, _ := get_player(id.(string))
    if (tagger.State != Human) {
        http.Error(w, "Killer is not a human", http.StatusForbidden)
        return
    }

    target, err := get_player(target_id)
    if err != nil {
        http.Error(w, "Target not found", http.StatusNotFound)
        return
    }
    if (target.State != Zombie && target.State != CoreZombie) {
        http.Error(w, "Target not a zombie", http.StatusConflict)
    }

    target.State = StunnedZombie
    target.Last_tagged = time.Now()
    tagger.Last_kill = time.Now()
    tagger.Kills += 1
    save_player(target)
    save_player(tagger)
    fmt.Fprintln(w, "Successfully killed target")
}

// @Summary Me
// @Description Shows user information
// @Produce json
// @Success 200 {object} Player
// @Failure 401 {string} string "User is not logged in"
// @Router /me [get]
func me (w http.ResponseWriter, r *http.Request) {
    sess, _ := store.Get(r, "session")
    id, ok := sess.Values["userID"]
    if !ok {
        http.Error(w, "User not logged in", http.StatusUnauthorized)
        return
    }

    player, _ := get_player(id.(string))
    json.NewEncoder(w).Encode(player)
}

// @Summary Cure
// @Description Cures a zombie or infected human
// @Produce plain
// @Success 200
// @Failure 401 {string} string "User is not logged in"
// @Failure 403 {string} string "No cures available"
// @Failure 404 {string} string "Target does not exist"
// @Failure 409 {string} string "Target is not a zombie"
// @Param target path string true "Target ID"
// @Router /cure [get]
func cure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    target_id := vars["target"]

    sess, _ := store.Get(r, "session")
    id, ok := sess.Values["userID"]
    if !ok {
        http.Error(w, "User not logged in", http.StatusUnauthorized)
        return
    }

    target, err := get_player(target_id)
    if err != nil {
        http.Error(w, "Target does not exist", http.StatusNotFound)
        return
    }

    player, _ := get_player(id.(string))
    if player.Cures == 0 {
        http.Error(w, "Not enough cures", http.StatusForbidden)
        return
    }

    if (target.State != Zombie && target.State != StunnedZombie && target.State != InfectedHuman) {
        http.Error(w, "Target is not a zombie", http.StatusConflict)
        return
    }

    target.State = Human
    player.Cures -= 1
    save_player(target)
    save_player(player)
}

// @Summary Revive
// @Description Revives a stunned zombie
// @Produce plain
// @Success 200
// @Failure 401 {string} string "User is not logged in"
// @Failure 403 {string} string "No revives available"
// @Failure 404 {string} string "Target does not exist"
// @Failure 409 {string} string "Target is not stunned"
// @Param target path string true "Target ID"
// @Router /revive [get]
func revive(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    target_id := vars["target"]

    sess, _ := store.Get(r, "session")
    id, ok := sess.Values["userID"]
    if !ok {
        http.Error(w, "User not logged in", http.StatusUnauthorized)
        return
    }

    target, err := get_player(target_id)
    if err != nil {
        http.Error(w, "Target does not exist", http.StatusNotFound)
        return
    }

    player, _ := get_player(id.(string))
    if player.Revives == 0 {
        http.Error(w, "Not enough revives", http.StatusForbidden)
        return
    }

    if (target.State != StunnedZombie) {
        http.Error(w, "Target is not stunned", http.StatusConflict)
        return
    }

    target.State = Zombie
    player.Revives -= 1
    save_player(target)
    save_player(player)
}

func main() {
    db, _ = sql.Open("sqlite3", "players.db")
    statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS players (id VARCHAR(7) PRIMARY KEY, name VARCHAR(64), password varchar(128), state INTEGER, last_tagged VARCHAR(25), kills INTEGER, last_kill VARCHAR(25), cures INTEGER, revives INTEGER, extensions INTEGER)")

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("table created successfully")
    }
    statement.Exec()

    go update()
    handleRequests()
}
