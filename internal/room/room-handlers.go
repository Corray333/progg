package room

import (
	"encoding/json"

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
		for _, p1 := range r.Players {
			p1.Conn.WriteMessage(1, append([]byte("04"), request...))
		}
		for _, p1 := range r.Players {
			p1.Conn.WriteMessage(1, append([]byte("00"), msg...))
		}
		break
	}
}
