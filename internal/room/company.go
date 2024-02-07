package room

import "github.com/Corray333/progg/internal/player"

type Company struct {
	// Name        string
	// Description string
	Price int
	Color string
	// Programs    []string
	// ProgsCount  int
	ProgPrice int
	// Owner       string
	// PlayersHere []string
	Key        string
	Number     int
	RentLevels []int
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

var CompanyKeys = []string{
	"skyeng",
	"yota",
	"hh",
	"europlan",
	"tensor",
	"finam",
	"megafon",
	"kaspersky",
	"sber",
	"astra",
	"geekbrains",
	"rostech",
	"1c",
	"beeline",
	"tinkoff",
	"rostelekom",
	"wot",
	"netbynet",
	"magnit",
	"yandex",
	"alpha",
	"ozon",
	"intellij",
	"vk",
}

var Companies = map[string]Company{
	"skyeng":     {60, "light-blue", 50, "skyeng", 2, []int{2, 10, 30, 90, 160}},
	"yota":       {60, "light-blue", 50, "yota", 4, []int{4, 20, 60, 180, 320}},
	"hh":         {200, "no", 0, "yota", 5, []int{}},
	"europlan":   {100, "orange", 50, "europlan", 6, []int{6, 30, 90, 270, 400}},
	"tensor":     {100, "orange", 50, "tensor", 8, []int{6, 30, 90, 270, 400}},
	"finam":      {120, "orange", 50, "finam", 9, []int{8, 40, 100, 300, 450}},
	"megafon":    {140, "green", 100, "megafon", 11, []int{10, 50, 150, 450, 625}},
	"kaspersky":  {140, "green", 100, "kaspersky", 13, []int{10, 50, 150, 450, 625}},
	"sber":       {160, "green", 100, "sber", 14, []int{12, 60, 180, 500, 700}},
	"astra":      {180, "white", 100, "astra", 16, []int{14, 70, 200, 550}},
	"geekbrains": {200, "no", 0, "geekbrains", 17, []int{}},
	"rostech":    {200, "white", 100, "rostech", 18, []int{16, 80, 220, 600}},
	"1c":         {220, "yellow", 150, "1c", 20, []int{18, 90, 250, 700}},
	"beeline":    {220, "yellow", 150, "beeline", 22, []int{20, 100, 300, 750}},
	"tinkoff":    {240, "yellow", 150, "tinkoff", 23, []int{20, 100, 300, 750}},
	"rostelekom": {280, "purple", 150, "rostelekom", 25, []int{24, 120, 360, 850}},
	"wot":        {200, "no", 0, "wot", 26, []int{}},
	"netbynet":   {260, "purple", 150, "netbynet", 27, []int{22, 110, 330, 800}},
	"magnit":     {300, "red", 200, "magnit", 29, []int{26, 130, 390, 900}},
	"yandex":     {300, "red", 200, "yandex", 30, []int{28, 250, 450, 1000}},
	"alpha":      {320, "red", 200, "alpha", 32, []int{26, 130, 390, 900}},
	"ozon":       {350, "blue", 200, "ozon", 34, []int{35, 175, 500, 1100}},
	"intellij":   {200, "no", 0, "intellij", 35, []int{}},
	"vk":         {400, "blue", 200, "vk", 36, []int{50, 200, 600, 1400}},
}
