package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "jopaks/pkg/mocks"
)

func GetSong(w http.ResponseWriter, r *http.Request) {
    // Read dynamic id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Iterate over all the mock books
    for _, Song := range mocks.Songs {
        if Song.Id == id {
            // If ids are equal send book as a response
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode(Song)
            break
        }
    }
}