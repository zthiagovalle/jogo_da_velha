package entity

type HashGame struct {
	Matriz [3][3]string `json:"matriz"`
}

func NewHashGame() *HashGame {
	return &HashGame{}
}
