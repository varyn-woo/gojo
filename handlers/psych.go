package handlers

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

type PsychState struct {
	Questions map[string]string
	Answers   map[string]string
}

var currState = PsychState{
	Questions: make(map[string]string),
	Answers:   make(map[string]string),
}
var psychLock = sync.RWMutex{}

type PsychTextInputType string

const (
	PsychTextInputTypeQuestion PsychTextInputType = "question"
	PsychTextInputTypeAnswer   PsychTextInputType = "answer"
)

type PsychTextInput struct {
	Player    string `json:"player"`
	Text      string `json:"text"`
	InputType string `json:"inputType"`
}

func HandleAddPsychTextInput(c *gin.Context) {
	var input PsychTextInput
	c.BindJSON(&input)
	log.Printf("adding text input: %s from player: %s for type: %s", input.Text, input.Player, input.InputType)
	if currState.Questions[input.Player] != "" {
		c.JSON(400, gin.H{"error": "text already exists for this player"})
		return
	}
	if input.Player == "" || input.Text == "" {
		c.JSON(400, gin.H{"error": "player and question must be provided"})
		return
	}

	psychLock.Lock()
	defer psychLock.Unlock()
	switch input.InputType {
	case string(PsychTextInputTypeQuestion):
		currState.Questions[input.Player] = input.Text
	case string(PsychTextInputTypeAnswer):
		currState.Answers[input.Player] = input.Text
	default:
		c.JSON(400, gin.H{"error": "invalid input type"})
		return
	}
	log.Printf("text input added: %s", input.Text)
	c.JSON(200, currState)
}

func HandleGetPsychState(c *gin.Context) {
	psychLock.RLock()
	defer psychLock.RUnlock()

	c.JSON(200, currState)
}
