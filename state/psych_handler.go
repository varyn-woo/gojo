package state

import (
	"fmt"
	"gojo/gen"
	"gojo/ui"
	"math/rand/v2"
	"sort"
	"strings"
	"sync"
	"time"
)

type PsychHandler struct {
	questions      []*gen.TextInput
	answers        map[string]*gen.TextInput
	votes          map[string]int
	playerSet      []string // set of player names to draw from when filling in a question
	playerSetIndex int      // index of player in the randomized set
	ready          int      // count of players ready to move on from results summary
	stage          gameStage
	timerDuration  time.Duration
	questionIndex  int
	lock           sync.RWMutex
}

type inputType string

const (
	InputTypeQuestion inputType = "question"
	InputTypeAnswer   inputType = "answer"
)

type gameStage int

const (
	questionInput gameStage = iota
	answerInput
	voting
	votingResults
)

const (
	DEFAULT_TIMER  = time.Minute * 2
	NEXT_BUTTON_ID = "resultButton"
)

func (p *PsychHandler) Sync() error {
	// update game state params based on Psych specific state
	p.lock.Lock()
	defer p.lock.Unlock()
	timerDone := game.GetElapsedTime() > p.timerDuration
	players := game.GetPlayers()
	if len(p.playerSet) == 0 || p.playerSetIndex >= len(p.playerSet) {
		// if players haven't been randomized or this randomization is used up, randomize them
		p.playerSetIndex = 0
		p.playerSet = make([]string, len(players))
		rand.Shuffle(len(players), func(i int, j int) {
			p.playerSet[i] = players[j].DisplayName
			p.playerSet[j] = players[i].DisplayName
		})
	}
	switch p.stage {
	case questionInput:
		if timerDone {
			// questions are done collecting, move to answer gathering
			p.stage = answerInput
			p.questionIndex = 0
			game.EndWait()
			game.ResetTimer()
		}
	case answerInput:
		if len(p.answers) == len(players) {
			// if every player has submitted, move on to voting
			p.stage = voting
			p.votes = make(map[string]int)
			for _, player := range players {
				p.votes[player.Id] = 0
			}
			game.EndWait()
			game.ResetTimer()
		}
	case voting:
		// tally up votes and if all votes are in, proceed with summary + winner
		totalVotes := 0
		maxVotes := 0
		var winnerId string
		for k, v := range p.votes {
			if v > maxVotes {
				winnerId = k
			}
			totalVotes += v
		}
		if totalVotes == len(players) {
			p.stage = votingResults
			game.IncrPlayerScore(winnerId, 1)
		}
	case votingResults:
		// do nothing, transition is triggered by User Input
	}
	p.runGameStateUpdates()
	return nil
}

func (p *PsychHandler) runGameStateUpdates() {
	elapsedTime := game.GetElapsedTime()
	game.lock.Lock()
	defer game.lock.Unlock()

	// update time remaining
	game.state.TimeRemaining = int32(elapsedTime / time.Second)

	switch p.stage {
	case questionInput:
		questionField := ui.MakeTextInput("questionField", &gen.TextField{
			Label:       "Enter questions with 'player' instead of real names. (only 1 player insert supported currently)",
			Placeholder: "What would player choose as their theme song?",
			InputType:   string(InputTypeQuestion),
		})
		game.state.UiElements = []*gen.UiElement{questionField}
	case answerInput:
		currentQuestion := p.questions[p.questionIndex].Text
		currentQuestion = strings.Replace(currentQuestion, "player", p.playerSet[p.playerSetIndex], 1)
		p.playerSetIndex++
		answerField := ui.MakeTextInput("answerField", &gen.TextField{
			Label:       currentQuestion,
			Placeholder: "your answer here",
			InputType:   string(InputTypeAnswer),
		})
		game.state.UiElements = []*gen.UiElement{answerField}
	case voting:
		currentQuestion := p.questions[p.questionIndex].Text
		options := make(map[string]string)
		for k, v := range p.answers {
			options[k] = v.Text
		}
		questionPrompt := ui.MakeSimpleText("questionPrompt", currentQuestion)
		votingOptions := ui.MakeVotingOptions("voting", options)
		game.state.UiElements = []*gen.UiElement{questionPrompt, votingOptions}
	case votingResults:
		voteList := sortedKeysByValue(p.votes)
		resultStrings := make([]string, len(voteList))
		for i, voteStat := range voteList {
			player := game.players[voteStat.Key]
			resultStrings[i] = fmt.Sprintf("%s - %d points", player.DisplayName, voteStat.Value)
		}
		title := ui.MakeSimpleText("resultText", "Voting Results")
		resultUi := ui.MakeStringList("resultList", resultStrings)
		nextButton := ui.MakeSimpleButton(NEXT_BUTTON_ID, "Next Round")
		game.state.UiElements = []*gen.UiElement{title, resultUi, nextButton}
	}
}

func (p *PsychHandler) HandleUserInput(input *gen.UserInputRequest) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	switch t := input.Request.(type) {
	case *gen.UserInputRequest_TextInputRequest:
		switch p.stage {
		case questionInput:
			if t.TextInputRequest.InputType != string(InputTypeQuestion) {
				return ErrOutOfSync
			}
			p.questions = append(p.questions, t.TextInputRequest)
		case answerInput:
			if t.TextInputRequest.InputType != string(InputTypeAnswer) {
				return ErrOutOfSync
			}
			p.answers[input.PlayerId] = t.TextInputRequest
		}
		// mark player as waiting since they submitted
		game.SetPlayerWaiting(input.PlayerId, true)
	case *gen.UserInputRequest_ButtonPressRequest:
		if p.stage != votingResults {
			return ErrOutOfSync
		}
		p.ready++
		if p.ready == len(game.players) {
			// move to next question and clear out the old answers and votes
			p.questionIndex++
			p.answers = make(map[string]*gen.TextInput)
			p.votes = make(map[string]int)
			// moving on from voting results page
			p.stage = answerInput
			p.ready = 0
		}
	case *gen.UserInputRequest_VoteRequest:
		if p.stage != voting {
			return ErrOutOfSync
		}
		p.votes[input.PlayerId] += int(t.VoteRequest.Rank)
		// mark player as waiting since they submitted
		game.SetPlayerWaiting(input.PlayerId, true)
	default:
		return ErrInvalidInput
	}
	return nil
}

func sortedKeysByValue[K comparable](m map[K]int) [](struct {
	Key   K
	Value int
}) {
	ss := []struct {
		Key   K
		Value int
	}{}
	for k, v := range m {
		ss = append(ss, struct {
			Key   K
			Value int
		}{Key: k, Value: v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}
