package engines

import (
	"fmt"
	"math/rand"

	"github.com/pascallohrer/fillergame/api"
	"github.com/pascallohrer/fillergame/server/util"
)

type coreEngine struct {
	tiles      []api.Color
	players    []bool
	turnPlayer int
	ownership  []int
}

const contested = -1

// AddPlayer adds a new player to the game and returns that added player's internal id.
func (e *coreEngine) AddPlayer() (int, error) {
	maxPlayers := len(e.players)
	var availableSlots []int
	for id, registered := range e.players {
		if !registered {
			availableSlots = append(availableSlots, id)
		}
	}
	openSlots := len(availableSlots)
	if openSlots == 0 {
		return 0, fmt.Errorf("Game already full! Only %d players allowed", maxPlayers)
	}
	newID := availableSlots[rand.Intn(openSlots)]
	e.players[newID] = true
	if e.IsRunning() {
		e.turnPlayer = 0
	}
	return newID, nil
}

// IsRunning returns false while the game is waiting for new players, then true.
func (e *coreEngine) IsRunning() bool {
	for _, registered := range e.players {
		if !registered {
			return false
		}
	}
	return true
}

// TurnPlayer returns the id of the player whose turn it currently is.
func (e *coreEngine) TurnPlayer() int {
	return e.turnPlayer
}

func (e *coreEngine) changeColor(
	player int,
	startingTile int,
	newColor api.Color,
	neighbours func(int) util.IntSet,
) map[int]api.Color {
	modifiedTiles := make(map[int]api.Color)
	currentColor := e.tiles[startingTile]
	e.tiles[startingTile] = newColor
	modifiedTiles[startingTile] = newColor
	for neighbourID := range neighbours(startingTile) {
		if e.tiles[neighbourID] == currentColor {
			for subID, subColor := range e.changeColor(player, neighbourID, newColor, neighbours) {
				modifiedTiles[subID] = subColor
			}
		}
	}
	for tile := range modifiedTiles {
		e.acquireTiles(tile, player, neighbours)
	}
	return modifiedTiles
}

func (e *coreEngine) acquireTiles(startingTile int, player int, neighbours func(int) util.IntSet) {
	for neighbourID := range neighbours(startingTile) {
		if e.ownership[neighbourID] == player {
			continue
		}
		if e.tiles[neighbourID] != e.tiles[startingTile] {
			continue
		}
		e.ownership[neighbourID] = player
		e.acquireTiles(neighbourID, player, neighbours)
	}
}

func (e *coreEngine) rebuildBuckets(neighbours func(int) util.IntSet) {
	contestedTiles := util.FilterInt(e.ownership, contested)
	for len(contestedTiles) > 0 {
		newBucket := e.buildBucket(util.IntSet{}, contestedTiles.Pop(), neighbours)
		contestedTiles.Remove(newBucket)
		if bucketOwner := e.classifyBucket(newBucket, neighbours); bucketOwner != contested {
			for tile := range newBucket {
				e.ownership[tile] = bucketOwner
			}
		}
	}
}

func (e *coreEngine) buildBucket(bucket util.IntSet, tile int, neighbours func(int) util.IntSet) util.IntSet {
	bucket.Add(tile)
	for neighbourID := range neighbours(tile) {
		if e.ownership[neighbourID] != contested {
			continue
		}
		if bucket.Contains(neighbourID) {
			continue
		}
		bucket = e.buildBucket(bucket, neighbourID, neighbours)
	}
	return bucket
}

func (e *coreEngine) classifyBucket(bucket util.IntSet, neighbours func(int) util.IntSet) int {
	bucketOwner := contested
	for tile := range bucket {
		for neighbour := range neighbours(tile) {
			neighbourOwner := e.ownership[neighbour]
			if neighbourOwner == contested || neighbourOwner == bucketOwner {
				continue
			}
			if bucketOwner != contested {
				return contested
			}
			bucketOwner = neighbourOwner
		}
	}
	return bucketOwner
}
