package gserver_test

import (
	"encoding/json"
	"gojo/gserver"
	"gojo/state"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// ServerTestSuite is a suite of tests for the game server
type ServerTestSuite struct {
	suite.Suite
	game   *state.Game
	server *httptest.Server
}

// SetupSuite runs once before any tests run
func (s *ServerTestSuite) SetupSuite() {
	// get a test server (httptest will run the server on any empty port)
	// this ensures there is no port conflict with a running server
	s.server = httptest.NewServer(gserver.GetGameServer().Handler())

	// make sure no global game is set
	_, err := state.GetGame()
	require.Error(s.T(), err)

	// initialize the game
	s.game = state.NewGame()
}

func (s *ServerTestSuite) TearDownSuite() {
	// close the test server when we're done testing
	s.server.Close()
}

func TestGameServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (s *ServerTestSuite) TestGetSetGame() {
	// make a GET request to /game_get
	resp, err := s.server.Client().Get(s.server.URL + "/game_get")
	require.NoError(s.T(), err)

	// check that the response status code is 200 (success)
	require.Equal(s.T(), 200, resp.StatusCode)
	// check that game is equal to the current game
	var game state.Game
	json.NewDecoder(resp.Body).Decode(&game)
	require.Equal(s.T(), s.game.Name, game.Name)
	require.Equal(s.T(), s.game.Players, game.Players)

	// make a POST request to /game_set
	req, _ := http.NewRequest(
		"POST",
		s.server.URL+"/game_set",
		strings.NewReader(`{"game": "psych"}`),
	)
	// set CORS origin header so it doesn't get rejected
	req.Header.Add("Origin", "test")
	require.NoError(s.T(), err)
	resp, _ = s.server.Client().Do(req)

	// check that the response status code is 200 (success)
	require.Equal(s.T(), 200, resp.StatusCode)
	// check that game state has changed to reflect the newly set game
	require.Equal(s.T(), "psych", s.game.Name)
}
