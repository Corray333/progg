package player

import "github.com/gorilla/websocket"

type Player struct {
	Username  string
	Money     int
	Companies []string
	Conn      *websocket.Conn
}

func NewPlayer(username string) Player {
	return Player{
		Username:  username,
		Money:     1500,
		Companies: []string{},
	}
}

func (p *Player) Pay(money int) {
	p.Money -= money
}

func (p *Player) AddMoney(money int) {
	p.Money += money
}

func (p *Player) AddCompany(company string) {
	p.Companies = append(p.Companies, company)
}
