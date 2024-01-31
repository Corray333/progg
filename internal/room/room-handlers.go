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
