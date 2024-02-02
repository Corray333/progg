package player

import "github.com/gorilla/websocket"

type Player struct {
	Username  string          `json:"username"`
	Money     int             `json:"money"`
	Avatar    int             `json:"avatar"`
	Ready     bool            `json:"ready"`
	Companies []string        `json:"companies"` // companyName + num of programs
	Position  int             `json:"position"`
	Hand      []int           `json:"hand"`
	Conn      *websocket.Conn `json:"-"`
}

func NewPlayer(username string) *Player {
	return &Player{
		Username:  username,
		Money:     1500,
		Avatar:    0,
		Companies: []string{},
		Position:  1,
		Hand:      []int{},
		Ready:     false,
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
