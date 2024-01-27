package app

import (
	"net/http"

	"github.com/Corray333/progg/internal/app/handlers"
	"github.com/Corray333/progg/internal/config"
	"github.com/Corray333/progg/internal/room"
	"github.com/go-chi/chi/v5"
)

type App struct {
	Rooms map[string]*room.Room
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	r := chi.NewRouter()
	r.Post("/rooms", handlers.CreateRoom(a.Rooms))
	r.Post("/rooms/{room}", handlers.JoinToRoom(a.Rooms))
	if err := http.ListenAndServe(config.GetAddress(), r); err != nil {
		panic(err)
	}
}
