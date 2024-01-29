package app

import (
	"log/slog"
	"net/http"

	"github.com/Corray333/progg/internal/app/handlers"
	"github.com/Corray333/progg/internal/config"
	"github.com/Corray333/progg/internal/room"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type App struct {
	Rooms map[string]*room.Room
}

func NewApp() *App {
	return &App{
		Rooms: make(map[string]*room.Room),
	}
}

func (a *App) Run() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/rooms", handlers.GetRooms(a.Rooms))
	r.Post("/rooms", handlers.CreateRoom(a.Rooms))
	r.Post("/rooms/{room}", handlers.JoinToRoom(a.Rooms))
	slog.Info("Server is running on " + config.GetAddress())
	if err := http.ListenAndServe(config.GetAddress(), r); err != nil {
		panic(err)
	}
}
