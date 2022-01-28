package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "jopaks/pkg/mocks"
    "jopaks/pkg/models"
)

func UpdateSong(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedSong models.Song
    json.Unmarshal(body, &updatedSong)

    for index, Song := range mocks.Songs {
        if Song.Id == id {
            
            Song.Title = updatedSong.Title
            Song.Artist = updatedSong.Artist
            

            mocks.Songs[index] = Song
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode("Updated")
            break
        }
    }
}