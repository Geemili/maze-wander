package game

import (
	"time"
)

type Conversation struct {
	Dialog []string
	Index  int
}

func (c Conversation) Act(mainLoop chan string) {
	if c.Index < len(c.Dialog) {
		go typeString(mainLoop, c.Dialog[c.Index])
		c.Index++
	}
}

func typeString(mainLoop chan string, text string) {
	for idx, _ := range text {
		mainLoop <- text[:idx+1]
		time.Sleep(50 * time.Millisecond)
	}
}
