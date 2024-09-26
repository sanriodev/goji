package definitions

type Emoji struct {
	Content string `json:"content"`
}

type Favorites struct {
	Emojis []Emoji `json:"emojis"`
}
