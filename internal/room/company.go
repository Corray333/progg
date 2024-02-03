package room

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

var Companies = map[string]Company{
	"skyeng":     Company{60, "light-blue", 50, "skyeng", 2},
	"yota":       Company{60, "light-blue", 50, "yota", 4},
	"hh":         Company{200, "no", 0, "yota", 5},
	"europlan":   Company{100, "orange", 50, "europlan", 6},
	"tensor":     Company{100, "orange", 50, "tensor", 8},
	"finam":      Company{120, "orange", 50, "finam", 9},
	"megafon":    Company{140, "green", 100, "megafon", 11},
	"kaspersky":  Company{140, "green", 100, "kaspersky", 13},
	"sber":       Company{160, "green", 100, "sber", 14},
	"astra":      Company{180, "white", 100, "astra", 16},
	"geekbrains": Company{200, "no", 0, "geekbrains", 17},
	"rostech":    Company{200, "white", 100, "rostech", 18},
	"1c":         Company{220, "yellow", 150, "1c", 20},
	"beeline":    Company{220, "yellow", 150, "beeline", 22},
	"tinkoff":    Company{240, "yellow", 150, "tinkoff", 23},
	"rostelekom": Company{280, "purple", 150, "rostelekom", 25},
	"wot":        Company{200, "no", 0, "wot", 26},
	"netbynet":   Company{260, "purple", 150, "netbynet", 27},
	"magnit":     Company{300, "red", 200, "magnit", 29},
	"yandex":     Company{300, "red", 200, "yandex", 30},
	"alpha":      Company{320, "red", 200, "alpha", 32},
	"ozon":       Company{350, "blue", 200, "ozon", 34},
	"intellij":   Company{200, "no", 0, "intellij", 35},
	"vk":         Company{400, "blue", 200, "vk", 36},
}
