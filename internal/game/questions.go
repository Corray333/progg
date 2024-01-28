package game

var Questions map[string][]Question // map of questions by theme

type Question struct {
	Question   string   `json:"question"`
	Answer     string   `json:"answer"`
	Type       uint8    `json:"type"` // 0 - type answer, 1 - choose answer, 2 - choose an order
	Variants   []string `json:"variants"`
	Attachment string   `json:"attachment"`
}
