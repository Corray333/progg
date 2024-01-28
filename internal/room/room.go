package room

import (
	"crypto/sha256"
	"errors"

	"github.com/Corray333/progg/internal/player"
)

type Room struct {
	Name     string
	Password [32]byte
	Master   string
	Players  []player.Player
}

func NewRoom(roomName, playerName, password string) *Room {
	return &Room{
		Name:     roomName,
		Password: sha256.Sum256([]byte(password)),
		Master:   playerName,
		Players: []player.Player{
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
			return errors.New("player already exists")
		}
	}
	r.Players = append(r.Players, player.NewPlayer(playerName))
	return nil
}
