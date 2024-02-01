package room

import (
	"crypto/sha256"
	"errors"
	"log/slog"
	"math/rand"
	"os"
	"slices"

	"github.com/Corray333/progg/internal/player"
)

const (
	TurnTime = 30
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
	avatars := []int{}
	for _, p := range r.Players {
		avatars = append(avatars, p.Avatar)
		if p.Username == playerName {
			return nil
		}
	}
	player := player.NewPlayer(playerName)
	dir, _ := os.ReadDir("../frontend/src/assets/avatars")
	l := len(dir)
	for {
		num := rand.Int() % l
		if !slices.Contains(avatars, num) {
			player.Avatar = num
			break
		}
	}
	r.Players = append(r.Players, player)
	return nil
}

func (r *Room) Start() {
	// TODO: exception is not enough players
	r.Started = true

	rand.Shuffle(len(r.Players), func(i, j int) {
		r.Players[i], r.Players[j] = r.Players[j], r.Players[i]
	})

	for _, p := range r.Players {
		p.Conn.WriteMessage(1, []byte("05"))
		p.Conn.WriteMessage(1, []byte("01"))
	}
	slog.Info("Game started in room " + r.Name)
	for i := 0; i < len(r.Players); i++ {
		// code of turn
		if i == len(r.Players)-1 {
			i = 0
		}
	}
}

// 00 - exit
// 01 - start the game
// 02 - turn started
// 03 - player moved
// 04 - player bought
// 05 - get all info

func (r *Room) Handle(query string, username string) {
	if len(query) < 2 {
		return
	}
	switch query[:2] {
	case "00":
		r.RemovePlayer(query[2:])
	case "01":
		f := true
		for _, p := range r.Players {
			if p.Username == username {
				p.Ready = true
			}
			f = f && p.Ready
		}
		if f {
			r.Start()
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
