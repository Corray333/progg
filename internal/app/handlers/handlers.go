package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/Corray333/progg/internal/room"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

func GetRooms(rooms map[string]*room.Room) http.HandlerFunc {
	type Resp struct {
		Rooms  []string `json:"rooms"`
		Amount int      `json:"amount"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		roomName := r.URL.Query().Get("name")
		roomList := make([]string, 0, 50)
		numQuery := r.URL.Query().Get("number")
		number := 0
		if numQuery != "" {
			number, err := strconv.Atoi(numQuery)
			if err != nil || number < 0 || number > len(rooms) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		i := 0
		if roomName != "" {
			re := regexp.MustCompile(".*" + roomName + ".*")
			for room := range rooms {
				if re.Find([]byte(room)) == nil {
					continue
				}
				if i < number {
					continue
				}
				if i == number+50 {
					break
				}

				roomList = append(roomList, room)
				i++
			}
		} else {
			for room := range rooms {
				if i < number {
					i++
					continue
				}
				if i == number+50 {
					break
				}
				roomList = append(roomList, room)
				i++
			}
		}
		if err := json.NewEncoder(w).Encode(Resp{roomList, number + len(roomList)}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

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
		slog.Info(fmt.Sprintf("Room created: %v", room))
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
			slog.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func Play(rooms map[string]*room.Room) http.HandlerFunc {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		defer c.Close()
		slog.Info("Test")
		_, msg, err := c.ReadMessage()
		if err != nil {
			slog.Error("read: " + err.Error())
		}
		slog.Info("msg: " + string(msg))
		for {
			err = c.WriteMessage(websocket.TextMessage, []byte("Hello, world!"))
			if err != nil {
				slog.Error("write:" + err.Error())
				break
			}
			time.Sleep(time.Second * 3)
		}
	}
}
