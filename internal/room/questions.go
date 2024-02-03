package room

type Quiz struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Variants []string `json:"variants"`
}

var Quizes = []Quiz{
	Quiz{
		Question: "Какая git команда отправляет изменения в удаленный репозиторий?",
		Answer:   "git push",
		Variants: []string{"git push", "git pull", "git commit", "git add"},
	},
	Quiz{
		Question: "Какая git команда создает новую ветку?",
		Answer:   "git branch",
		Variants: []string{"git branch", "git checkout", "git merge", "git add"},
	},
	Quiz{
		Question: "Какая git команда переключает ветки?",
		Answer:   "git checkout",
		Variants: []string{"git checkout", "git branch", "git merge", "git add"},
	},
	Quiz{
		Question: "... - это принцип, позволяющий объектам разных типов обрабатываться как объекты одного типа.",
		Answer:   "Полиморфизм",
		Variants: []string{"Инкапсуляция", "Полиморфизм", "Наследование", "Абстракция"},
	},
	Quiz{
		Question: "... - это простейшая форма искусственной нейронной сети, модель бинарного классификатора.",
		Answer:   "Персептрон",
		Variants: []string{"Сеть Хопфилда", "Сеть Кохонена", "Персептрон", "Сеть Болцмана"},
	},
}
