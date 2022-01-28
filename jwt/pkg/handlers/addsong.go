package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"

    "jopaks/pkg/mocks"
    "jopaks/pkg/models"
)

func AddSong(w http.ResponseWriter, r *http.Request) {
    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var Song models.Song
    json.Unmarshal(body, &Song)
    Song.Id = rand.Intn(100)
    mocks.Songs = append(mocks.Songs, Song)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}