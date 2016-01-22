package game

type MessageBox struct {
	X, Y    int
	W, H    int
	Message string
	Visible bool
}

func NewMessageBox() *MessageBox {
	return &MessageBox{
		0, 36,
		80, 6,
		"",
		true,
	}
}
