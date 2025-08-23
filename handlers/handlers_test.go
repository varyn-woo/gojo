package handlers_test

import (
	"gojo/gen"
	"gojo/handlers"
	"reflect"
	"testing"
)

func TestRouteUserInput(t *testing.T) {
	// Test cases for RouteUserInput function
	tests := []struct {
		name     string
		input    *gen.UserInputRequest
		expected func() *gen.ServerResponse
	}{
		{
			name: "New Game Request",
			input: &gen.UserInputRequest{
				Request: &gen.UserInputRequest_NewGameRequest{
					NewGameRequest: "game",
				},
			},
			expected: handlers.MakeGameStateResponse,
		},
		{
			name: "Player Add Request",
			input: &gen.UserInputRequest{
				Request: &gen.UserInputRequest_PlayerAddRequest{
					PlayerAddRequest: &gen.PlayerAddRequest{
						PlayerId:    "123",
						DisplayName: "Test Player",
					},
				},
			},
			expected: handlers.MakeAcknowledgementResponse,
		},
		{
			name: "Start Game Request",
			input: &gen.UserInputRequest{
				Request: &gen.UserInputRequest_StartGameRequest{},
			},
			expected: handlers.MakeGameStateResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := handlers.RouteUserInput(tt.input)
			expected := tt.expected()
			if !reflect.DeepEqual(response, expected) {
				t.Errorf("expected %v, got %v", expected, response)
			}
		})
	}
}
