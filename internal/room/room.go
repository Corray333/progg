package room

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"slices"
	"time"

	"github.com/Corray333/progg/internal/player"
)

const (
	TurnTime = 30
)

type Room struct {
	Name             string           `json:"name`
	Password         [32]byte         `json:"-"`
	ActivePlayer     string           `json:"active_player"`
	Started          bool             `json:"started"`
	EndOfCurrentTurn time.Time        `json:"end_of_current_turn"`
	Players          []*player.Player `json:"players"`
}

func NewRoom(roomName, playerName, password string) *Room {
	return &Room{
		Name:         roomName,
		Password:     sha256.Sum256([]byte(password)),
		ActivePlayer: playerName,
		Started:      false,
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
	if r.Started {
		return errors.New("game already started")
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
		fmt.Printf("%v\n", p)
		r.SendAllInfo(p)
		p.Conn.WriteMessage(1, []byte("01"))
	}
	slog.Info("Game started in room " + r.Name)
	for i := 0; i < len(r.Players); i++ {
		r.ActivePlayer = r.Players[i].Username
		// TODO: turn it to function
		type Turn struct {
			Username  string    `json:"username"`
			EndOfTurn time.Time `json:"end_of_turn"`
		}
		r.EndOfCurrentTurn = time.Now().Add(TurnTime * time.Second)
		req := Turn{
			Username:  r.Players[i].Username,
			EndOfTurn: r.EndOfCurrentTurn,
		}
		res, err := json.Marshal(req)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		go func() {
			res = append([]byte("02"), res...)
			for _, p := range r.Players {
				err := p.Conn.WriteMessage(1, res)
				if err != nil {
					slog.Error("writing message: " + err.Error())
				}
			}
		}()
		time.Sleep(TurnTime * time.Second)
		r.Players[i].AlreadyWalked = false
		if i == len(r.Players)-1 {
			i = -1
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
	slog.Info("Handling query: " + query)
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
			go r.Start()
		}
	case "02":
		if !r.Started {
			for _, p := range r.Players {
				if p.Username != username {
					continue
				}
				err := p.Conn.WriteMessage(1, []byte("02no"))
				if err != nil {
					slog.Error("writing message: " + err.Error())
				}
				return
			}
		}
		type Turn struct {
			Username  string    `json:"username"`
			EndOfTurn time.Time `json:"end_of_turn"`
		}
		req := Turn{
			Username:  r.ActivePlayer,
			EndOfTurn: r.EndOfCurrentTurn,
		}
		res, err := json.Marshal(req)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		res = append([]byte("02"), res...)
		for _, p := range r.Players {
			if p.Username != username {
				continue
			}
			err := p.Conn.WriteMessage(1, res)
			if err != nil {
				slog.Error("writing message: " + err.Error())
			}
			return
		}
	case "03":
		for _, p := range r.Players {
			if p.Username == username && p.Username == r.ActivePlayer {
				steps := rand.Int()%6 + 1
				p.Position = (p.Position+steps)%36 + 1
				if p.AlreadyWalked {
					return
				}
				p.AlreadyWalked = true
				for _, p1 := range r.Players {
					type Resp struct {
						Username string `json:"username"`
						Position int    `json:"position"`
						Steps    int    `json:"steps"`
					}
					resp, err := json.Marshal(Resp{r.ActivePlayer, p.Position, steps})
					if err != nil {
						slog.Error(err.Error())
						return
					}
					p1.Conn.WriteMessage(1, append([]byte("03"), resp...))
				}
				return
			}
		}
	case "04":
		r.NewBuy(username, query[2:])
	case "05":
		// TODO
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
