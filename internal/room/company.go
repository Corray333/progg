package room

import "github.com/Corray333/progg/internal/player"

type Company struct {
	// Name        string
	// Description string
	Price int
	Color string
	// Programs    []string
	// ProgsCount  int
	// RentLevels  []int
	ProgPrice int
	// Owner       string
	// PlayersHere []string
	Key    string
	Number int
}

type ChanceCard struct {
	Text          string
	IsImmediately bool
	Func          func(*player.Player)
}

var Chances = []int{3, 7, 12, 15, 21, 24, 31, 33}
var ChanceCards = []ChanceCard{
	// "Отправляйтесь в тюрьму.",
	// "Получите 200к.",
	// "Выплатите 200к.",
	// "Отправляйтесь на старт.",
	// "Вернитесь на три поля назад.",
	// "Отправляйтесь на три поля вперед.",
	{
		Text:          "Отправляйтесь в тюрьму.",
		IsImmediately: true,
		Func: func(p *player.Player) {
			p.Position = 10
			p.InPrisonFor = 2
		},
	},
	{
		Text: "Получите 200к.",
		Func: func(p *player.Player) {
			p.Money += 200
			return
		},
	},
	{
		Text: "Выплатите 200к.",
		Func: func(p *player.Player) {
			p.Money -= 200
			return
		},
	},
	{
		Text: "Отправляйтесь на старт.",
		Func: func(p *player.Player) {
			p.Position = 1
			return
		},
	},
	{
		Text: "Вернитесь на три поля назад.",
		Func: func(p *player.Player) {
			p.Position = (p.Position-3)%36 + 1
			return
		},
	},
	{
		Text: "Отправьтесь на три поля вперед.",
		Func: func(p *player.Player) {
			p.Position = (p.Position+3)%36 + 1
			return
		},
	},
	{
		Text: "Освободитесь из тюрьмы.",
		Func: func(p *player.Player) {
			p.InPrisonFor = 0
			return
		},
	},
	// TODO
	// {
	// 	Text: "Освободитесь от вопроса.",
	// 	Func: func(p *player.Player) {
	// 		p.InPrisonFor = 0
	// 		return
	// 	},
	// },
}

var Companies = map[string]Company{
	"skyeng":     {60, "light-blue", 50, "skyeng", 2},
	"yota":       {60, "light-blue", 50, "yota", 4},
	"hh":         {200, "no", 0, "yota", 5},
	"europlan":   {100, "orange", 50, "europlan", 6},
	"tensor":     {100, "orange", 50, "tensor", 8},
	"finam":      {120, "orange", 50, "finam", 9},
	"megafon":    {140, "green", 100, "megafon", 11},
	"kaspersky":  {140, "green", 100, "kaspersky", 13},
	"sber":       {160, "green", 100, "sber", 14},
	"astra":      {180, "white", 100, "astra", 16},
	"geekbrains": {200, "no", 0, "geekbrains", 17},
	"rostech":    {200, "white", 100, "rostech", 18},
	"1c":         {220, "yellow", 150, "1c", 20},
	"beeline":    {220, "yellow", 150, "beeline", 22},
	"tinkoff":    {240, "yellow", 150, "tinkoff", 23},
	"rostelekom": {280, "purple", 150, "rostelekom", 25},
	"wot":        {200, "no", 0, "wot", 26},
	"netbynet":   {260, "purple", 150, "netbynet", 27},
	"magnit":     {300, "red", 200, "magnit", 29},
	"yandex":     {300, "red", 200, "yandex", 30},
	"alpha":      {320, "red", 200, "alpha", 32},
	"ozon":       {350, "blue", 200, "ozon", 34},
	"intellij":   {200, "no", 0, "intellij", 35},
	"vk":         {400, "blue", 200, "vk", 36},
}
