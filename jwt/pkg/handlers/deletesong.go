package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "jopaks/pkg/mocks"
)

func DeleteSong(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    for index, Song := range mocks.Songs {
        if Song.Id == id {
            
            mocks.Songs = append(mocks.Songs[:index], mocks.Songs[index+1:]...)

            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Deleted")
            break
        }
    }
}