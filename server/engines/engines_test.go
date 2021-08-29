package engines

import (
	"testing"

	"github.com/pascallohrer/fillergame/api"
	"gotest.tools/assert"
)

func TestSquare10RepresentativeGame(t *testing.T) {
	engine := NewSquare10()
	t.Run("References for both players are complementary", func(t *testing.T) {
		board_player_0 := engine.CurrentBoard(0)
		board_player_1 := engine.CurrentBoard(1)
		for n := 0; n < 100; n++ {
			assert.Equal(t, board_player_0[n], board_player_1[99-n])
		}
	})
	engine.tiles = buildSq10(1, 2, 3, 4, 14, 24, 34, 44, 54, 64, 74, 84, 94, 95, 96, 97, 98)
	t.Run("Game is not running after initialization", func(t *testing.T) {
		assert.Assert(t, !engine.IsRunning())
	})
	t.Run("Game is not running after first player joins", func(t *testing.T) {
		engine.AddPlayer()
		assert.Assert(t, !engine.IsRunning())
	})
	t.Run("Game is running after second player joins", func(t *testing.T) {
		engine.AddPlayer()
		assert.Assert(t, engine.IsRunning())
	})
	t.Run("Additional Player cannot join", func(t *testing.T) {
		_, err := engine.AddPlayer()
		assert.ErrorContains(t, err, "full")
	})
	t.Run("Player ID 0 goes first", func(t *testing.T) {
		assert.Equal(t, engine.TurnPlayer(), 0)
	})
	t.Run("Opponent cannot make a move", func(t *testing.T) {
		_, err := engine.ProcessMove(1, api.Color_GREEN)
		assert.ErrorContains(t, err, "invalid player")
	})
	t.Run("No winner in the initial state", func(t *testing.T) {
		assert.Equal(t, engine.DetermineWinner(), -1)
	})
	t.Run("Player cannot choose current color", func(t *testing.T) {
		_, err := engine.ProcessMove(0, api.Color_PURPLE)
		assert.ErrorContains(t, err, "illegal move")
	})
	t.Run("Player cannot choose opponent's color", func(t *testing.T) {
		_, err := engine.ProcessMove(0, api.Color_YELLOW)
		assert.ErrorContains(t, err, "illegal move")
	})
	mod, err := engine.ProcessMove(0, api.Color_RED)
	t.Run("Turn Player can make successful move", func(t *testing.T) {
		assert.NilError(t, err)
	})
	t.Run("Modifiers are complementary after call to Reference()", func(t *testing.T) {
		opponent_mod := engine.Reference(mod, 1)
		for n, color := range mod {
			opponent_color, exists := opponent_mod[99-n]
			assert.Assert(t, exists)
			assert.Equal(t, color, opponent_color)
		}
	})
	t.Run("Player has not won by choosing color with large disjoint regions", func(t *testing.T) {
		assert.Equal(t, engine.DetermineWinner(), -1)
	})
	t.Run("Turn Player changed after successful move", func(t *testing.T) {
		assert.Equal(t, engine.TurnPlayer(), 1)
	})
	t.Run("Player wins after choosing divider color to cut off winning region", func(t *testing.T) {
		engine.ProcessMove(1, api.Color_GREEN)
		assert.Equal(t, engine.DetermineWinner(), 1)
	})
}

func buildSq10(greenTiles ...int) []api.Color {
	emptyBoard := make([]api.Color, 100) // default color is RED
	for _, tile := range greenTiles {
		emptyBoard[tile] = api.Color_GREEN
	}
	emptyBoard[0] = api.Color_PURPLE
	emptyBoard[99] = api.Color_YELLOW
	return emptyBoard
}
