package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"

	"github.com/Corray333/progg/internal/room"
	"github.com/go-chi/chi/v5"
)

func CreateRoom(rooms map[string]*room.Room) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RoomName   string `json:"room_name"`
			PlayerName string `json:"player_name"`
			Password   string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		room := room.NewRoom(req.RoomName, req.PlayerName, req.Password)
		rooms[req.RoomName] = room
		w.WriteHeader(http.StatusCreated)
	}
}

func JoinToRoom(rooms map[string]*room.Room) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomName := chi.URLParam(r, "room")
		var req struct {
			PlayerName string `json:"player_name"`
			Password   string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		room, ok := rooms[roomName]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if room.Password != sha256.Sum256([]byte(req.Password)) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err := room.NewPlayer(req.PlayerName); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
