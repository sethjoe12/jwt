package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "jopaks/pkg/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/songs", handlers.AddSong).Methods(http.MethodPost)
    router.HandleFunc("/songs", handlers.GetAllSong).Methods(http.MethodGet)
    router.HandleFunc("/songs/{id}", handlers.GetSong).Methods(http.MethodGet)
    router.HandleFunc("/songs/{id}", handlers.UpdateSong).Methods(http.MethodPut)
    router.HandleFunc("/songs/{id}", handlers.DeleteSong).Methods(http.MethodDelete)

    log.Println("API is running!")
    http.ListenAndServe(":8000", router)
}
