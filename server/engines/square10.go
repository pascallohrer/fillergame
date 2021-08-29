package engines

import (
	"fmt"
	"math/rand"

	"github.com/pascallohrer/fillergame/api"
	"github.com/pascallohrer/fillergame/server/util"
)

//A Square10 is a 10-by-10 square grid for 2 players
type Square10 struct {
	*coreEngine
}

// NewSquare10 creates and initializes an engine for the Square10 board layout.
func NewSquare10() *Square10 {
	randomTiles := [100]api.Color{}
	for id := range randomTiles {
		randomTiles[id] = api.Color(rand.Intn(len(api.Color_name)))
	}
	for _, next := range []int{1, 10, 89, 98} {
		for randomTiles[next] == randomTiles[0] || randomTiles[next] == randomTiles[99] {
			randomTiles[next]++
			if randomTiles[next] >= api.Color(len(api.Color_name)) {
				randomTiles[next] = api.Color(0)
			}
		}
	}
	ownership := make([]int, 100, 100)
	for centerID := 1; centerID < 99; centerID++ {
		ownership[centerID] = contested
	}
	ownership[99] = 1
	return &Square10{&coreEngine{
		tiles:      randomTiles[:],
		players:    []bool{false, false},
		turnPlayer: -1,
		ownership:  ownership,
	}}
}

// CurrentBoard returns a slice of colors representing the current board state.
func (e *Square10) CurrentBoard(referencePlayer int) []api.Color {
	referenceBoard := [100]api.Color{}
	for tile, color := range e.tiles {
		targetID := referencePlayer*99 - tile*(referencePlayer*2-1)
		referenceBoard[targetID] = color
	}
	return referenceBoard[:]
}

// Reference transforms the modifiers into the perspective of refPlayer.
func (e *Square10) Reference(modifiers map[int]api.Color, refPlayer int) map[int]api.Color {
	referenceBoard := make(map[int]api.Color, len(modifiers))
	for tile, color := range modifiers {
		targetID := refPlayer*99 - tile*(refPlayer*2-1)
		referenceBoard[targetID] = color
	}
	return referenceBoard
}

// ProcessMove changes player's Color to newColor and returns a slice of all changed tiles
func (e *Square10) ProcessMove(player int, newColor api.Color) (map[int]api.Color, error) {
	if player != e.turnPlayer {
		return nil, fmt.Errorf("invalid player, this should be caught in game.go")
	}
	if newColor == e.tiles[0] || newColor == e.tiles[99] {
		return nil, fmt.Errorf("illegal move: no player's current color may be chosen")
	}
	playerHome := player * 99
	modifiers := e.changeColor(player, playerHome, newColor, neighbours)
	e.rebuildBuckets(neighbours)
	e.turnPlayer = 1 - player
	return modifiers, nil
}

func neighbours(tile int) util.IntSet {
	validTiles := util.IntSet{}
	if tile%10 != 0 {
		validTiles.Add(tile - 1)
	}
	if tile%10 != 9 {
		validTiles.Add(tile + 1)
	}
	if tile > 9 {
		validTiles.Add(tile - 10)
	}
	if tile < 90 {
		validTiles.Add(tile + 10)
	}
	return validTiles
}

// DetermineWinner returns the player ID of the winner or -1 if there is no winner yet.
// If both players control half the board, the player going second (playerID 1) wins.
func (e *Square10) DetermineWinner() int {
	score := []int{0, 0}
	for _, owner := range e.ownership {
		if owner == contested {
			continue
		}
		score[owner]++
	}
	for player, points := range score {
		if points+player > 50 {
			return player
		}
	}
	return -1
}

func (e *Square10) String() string {
	return fmt.Sprint(e.coreEngine)
}
