package room

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	"slices"
	"strconv"

	"github.com/Corray333/progg/internal/player"
)

func (r *Room) RemovePlayer(playerName string) {
	for i, p := range r.Players {
		if p.Username == playerName {
			r.Players = append(r.Players[:i], r.Players[i+1:]...)
			return
		}
	}
}

func (r *Room) SendAllInfo(p *player.Player) {
	res, err := json.Marshal(r.Players)
	if err != nil {
		p.Conn.WriteMessage(1, []byte("error"))
		return
	}
	res = append([]byte("05"), res...)
	p.Conn.WriteMessage(1, res)

}

func (r *Room) NewBuy(username, request string) {
	type Req struct {
		Company string `json:"company"`
		Answer  string `json:"answer"`
		QuizID  int    `json:"quiz_id"`
		IsGame  bool   `json:"is_game"`
	}
	var req Req
	err := json.Unmarshal([]byte(request), &req)
	if err != nil {
		return
	}

	for _, p := range r.Players {
		if p.Username != username {
			continue
		}
		if req.IsGame && req.Answer != GameQuizes[req.QuizID].Answer || req.Answer != Quizes[req.QuizID].Answer {
			for _, p1 := range r.Players {
				p1.Conn.WriteMessage(1, []byte("04no"))
			}
			return
		}
		if p.Money < Companies[req.Company].Price {
			p.Conn.WriteMessage(1, []byte("04money"))
			return
		}
		p.Money -= Companies[req.Company].Price
		for _, p1 := range r.Players {
			for _, comp := range p1.Companies {
				if comp[:len(comp)-1] == req.Company {
					return
				}
			}
		}
		p.Companies = append(p.Companies, req.Company+"1")
		msg, err := json.Marshal(p)
		if err != nil {
			return
		}
		r.StopTheQuiz()
		for _, p1 := range r.Players {
			r.Mu.Lock()
			p1.Conn.WriteMessage(1, append([]byte("04"), request...))
			r.Mu.Unlock()
		}
		for _, p1 := range r.Players {
			r.Mu.Lock()
			p1.Conn.WriteMessage(1, append([]byte("00"), msg...))
			r.Mu.Unlock()
		}
		break
	}
}

func (r *Room) Walk(username string) {
	for _, p := range r.Players {
		if p.Username == username && p.Username == r.ActivePlayer {
			steps := rand.Int()%6 + 1
			p.Position = (p.Position+steps)%36 + 1
			if p.AlreadyWalked {
				return
			}
			p.AlreadyWalked = true
			if slices.Contains(Chances, p.Position) {
				card := rand.Int() % len(ChanceCards)
				ChanceCards[card].Func(p)
				r.Mu.Lock()
				p.Conn.WriteMessage(1, []byte("08"+strconv.Itoa(card)))
				r.Mu.Unlock()
				msg, err := json.Marshal(p)
				if err != nil {
					return
				}
				for _, p1 := range r.Players {
					r.Mu.Lock()
					p1.Conn.WriteMessage(1, append([]byte("00"), msg...))
					r.Mu.Unlock()
				}
			}
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
				r.Mu.Lock()
				p1.Conn.WriteMessage(1, append([]byte("03"), resp...))
				r.Mu.Unlock()
			}
			return
		}
	}
}

func (r *Room) StartQuiz(query string) {
	var quiz Quiz
	var qid int
	if query[2:] == "true" {
		qid = rand.Int() % len(GameQuizes)
		quiz = GameQuizes[qid]
	} else {
		qid = rand.Int() % len(Quizes)
		quiz = Quizes[qid]
	}
	type Resp struct {
		Quiz
		QuizID int `json:"quiz_id"`
	}
	resp, err := json.Marshal(Resp{quiz, qid})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	for _, p := range r.Players {
		r.Mu.Lock()
		p.Conn.WriteMessage(1, append([]byte("06"), resp...))
		r.Mu.Unlock()
	}

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		r.StopTheQuiz = cancel
		for {
			select {
			case <-ctx.Done():
				for _, p := range r.Players {
					r.Mu.Lock()
					p.Conn.WriteMessage(1, []byte("04time"))
					r.Mu.Unlock()
				}
				return
			default:
				continue
			}
		}
	}()
}
