package game

import (
	"time"
)

type Conversation struct {
	DialogBox   *MessageBox
	Dialog      []string
	Index       int
	isExecuting bool
	actionChan  chan action
}

type action struct {
	code int
	text string
}

func NewConversation(dialogBox *MessageBox, dialog ...string) *Conversation {
	return &Conversation{
		dialogBox,
		dialog,
		0, false, make(chan action, 1),
	}
}

func (c *Conversation) Act() {
	if !c.isExecuting && c.Index < len(c.Dialog) {
		go typeString(c.actionChan, c.Dialog[c.Index])
		c.Index++
		c.isExecuting = true
	}
}

func (c *Conversation) Tick() {
	if c.isExecuting {
		action := <-c.actionChan
		switch action.code {
		case 1:
			c.DialogBox.Message = action.text
		case 2:
			c.isExecuting = false
		default:
			c.DialogBox.Message = "Unknown action code " + string(action.code)
		}
	}
}

func typeString(msg chan action, text string) {
	for idx, _ := range text {
		msg <- action{1, text[:idx+1]} // Update dialog box
		time.Sleep(50 * time.Millisecond)
	}
	msg <- action{2, ""} // End typing
}
