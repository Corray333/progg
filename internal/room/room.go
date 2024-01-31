package room

import (
	"crypto/sha256"
	"errors"
	"log/slog"

	"github.com/Corray333/progg/internal/player"
)

type Room struct {
	Name     string           `json:"name`
	Password [32]byte         `json:"-"`
	Master   string           `json:"master"`
	Started  bool             `json:"started"`
	Players  []*player.Player `json:"players"`
}

func NewRoom(roomName, playerName, password string) *Room {
	return &Room{
		Name:     roomName,
		Password: sha256.Sum256([]byte(password)),
		Master:   playerName,
		Started:  false,
		Players: []*player.Player{
			player.NewPlayer(playerName),
		},
	}
}

func (r *Room) NewPlayer(playerName string) error {
	if len(r.Players) >= 6 {
		return errors.New("room is full")
	}
	for _, p := range r.Players {
		if p.Username == playerName {
			return nil
		}
	}
	r.Players = append(r.Players, player.NewPlayer(playerName))
	return nil
}

func (r *Room) Start() {
	// TODO: exception is not enough players
	r.Started = true
}

// 00 - exit
// 01 - start the game
// 02 - turn started
// 03 - player moved
// 04 - player bought
// 05 - get all info

func (r *Room) Handle(query string) {
	if len(query) < 2 {
		return
	}
	switch query[:2] {
	case "00":
		r.RemovePlayer(query[2:])
	case "01":
		r.Start()
		for _, p := range r.Players {
			r.SendAllInfo(p)
		}
	case "02":
	case "03":
	case "04":
	case "05":
		player := string(query[2:])
		for _, p := range r.Players {
			if p.Username == player {
				r.SendAllInfo(p)
			}
		}
	default:
		slog.Info("Unknown query: " + query)
	}
}
