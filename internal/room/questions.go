package room

type Quiz struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Variants []string `json:"variants"`
}

var Quizes = []Quiz{
	{
		Question: "Какая git команда отправляет изменения в удаленный репозиторий?",
		Answer:   "git push",
		Variants: []string{"git push", "git pull", "git commit", "git add"},
	},
	{
		Question: "Какая git команда создает новую ветку?",
		Answer:   "git branch",
		Variants: []string{"git branch", "git checkout", "git merge", "git add"},
	},
	{
		Question: "Какая git команда переключает ветки?",
		Answer:   "git checkout",
		Variants: []string{"git checkout", "git branch", "git merge", "git add"},
	},
	{
		Question: "... - это принцип, позволяющий объектам разных типов обрабатываться как объекты одного типа.",
		Answer:   "Полиморфизм",
		Variants: []string{"Инкапсуляция", "Полиморфизм", "Наследование", "Абстракция"},
	},
	{
		Question: "... - это простейшая форма искусственной нейронной сети, модель бинарного классификатора.",
		Answer:   "Персептрон",
		Variants: []string{"Сеть Хопфилда", "Сеть Кохонена", "Персептрон", "Сеть Болцмана"},
	},
}

var GameQuizes = []Quiz{
	{
		Question: "Первой игрой в жанре first person shooter была ...",
		Answer:   "Wolfenstein 3D",
		Variants: []string{"Doom", "Wolfenstein 3D", "Quake", "Half-Life"},
	},
}
