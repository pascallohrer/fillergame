package server

import (
	"fmt"

	"github.com/pascallohrer/fillergame/api"
	"github.com/pascallohrer/fillergame/server/engines"
)

/* New supported Game Modes must be registered here!
 * Key: Enum value from the protobuf-generated file
 * Value: New-function that returns an engine which implements gameEngine
 */
var gameModes = map[api.GameBoard]func() gameEngine{
	api.GameBoard_SQUARE10: func() gameEngine { return engines.NewSquare10() },
	//api.GameBoard_$MODE$: func() gameEngine { return engines.New$Mode$() },
}

type gameEngine interface {
	AddPlayer() (int, error)
	IsRunning() bool
	TurnPlayer() int
	DetermineWinner() int
	CurrentBoard(referencePlayer int) []api.Color
	Reference(modifiers map[int]api.Color, refPlayer int) map[int]api.Color
	ProcessMove(player int, newColor api.Color) (map[int]api.Color, error)
}

type fillerGame struct {
	mode    api.GameBoard
	players map[string]player
	gameEngine
}

type player struct {
	id     int
	stream chan *api.BoardState
}

func (p *player) setStream(newStream chan *api.BoardState) {
	if p.stream != nil {
		close(p.stream)
	}
	p.stream = newStream
}

func newGame(mode api.GameBoard) *fillerGame {
	return &fillerGame{mode, map[string]player{}, gameModes[mode]()}
}

func (game *fillerGame) validateMode(mode api.GameBoard) error {
	if game.mode != mode {
		return fmt.Errorf("requested mode %v does not match the game's mode %v", mode, game.mode)
	}
	return nil
}

func (game *fillerGame) addPlayer(joiningPlayer string, stream chan *api.BoardState) error {
	playerID, err := game.AddPlayer()
	if err != nil {
		return err
	}
	game.players[joiningPlayer] = player{playerID, stream}
	go func() {
		stream <- &api.BoardState{
			Status: game.status(joiningPlayer),
			Tiles:  game.boardState(joiningPlayer),
		}
	}()
	if game.IsRunning() {
		fmt.Println("Game started!")
		go game.sendStreams(nil)
	}
	return nil
}

func (game *fillerGame) status(referencePlayer string) *api.Status {
	winner := game.DetermineWinner()
	if game.players[referencePlayer].id == winner {
		return &api.Status{
			StatusCode: api.StatusCode_GAME_OVER,
			Message:    "You win!",
		}
	}
	if winner != -1 {
		return &api.Status{
			StatusCode: api.StatusCode_GAME_OVER,
			Message:    "You lose.",
		}
	}
	turnPlayer := game.TurnPlayer()
	if turnPlayer == -1 {
		return &api.Status{
			StatusCode: api.StatusCode_WAITING_FOR_PLAYERS,
			Message:    "Waiting for more players to join the game.",
		}
	}
	if game.players[referencePlayer].id == turnPlayer {
		return &api.Status{
			StatusCode: api.StatusCode_YOUR_TURN,
			Message:    "It's your turn.",
		}
	}
	return &api.Status{
		StatusCode: api.StatusCode_OPPONENTS_TURN,
		Message:    "Waiting for an opponent's turn.",
	}
}

func (game *fillerGame) boardState(referencePlayer string) []*api.Modifier {
	result := make([]*api.Modifier, 0)
	for id, color := range game.CurrentBoard(game.players[referencePlayer].id) {
		result = append(result, &api.Modifier{
			TileId:   int32(id),
			NewColor: color,
		})
	}
	return result
}

func (game *fillerGame) processMove(
	movingPlayer string,
	newColor api.Color,
) {
	player := game.players[movingPlayer]
	if player.id != game.TurnPlayer() {
		player.stream <- &api.BoardState{
			Status: &api.Status{
				StatusCode: api.StatusCode_NOT_YOUR_TURN,
				Message:    "It is not your turn. Move ignored.",
			},
			Tiles: []*api.Modifier{},
		}
		return
	}
	modifiers, err := game.ProcessMove(player.id, newColor)
	if err != nil {
		player.stream <- &api.BoardState{
			Status: &api.Status{
				StatusCode: api.StatusCode_ILLEGAL_MOVE,
				Message:    "Cannot choose a player's current color. Move ignored.",
			},
			Tiles: []*api.Modifier{},
		}
		return
	}
	game.sendStreams(modifiers)
}

func (game *fillerGame) sendStreams(modifiers map[int]api.Color) {
	for token, player := range game.players {
		refModifiers := game.Reference(modifiers, player.id)
		modifierObjects := make([]*api.Modifier, 0, len(refModifiers))
		for tile, color := range refModifiers {
			modifierObjects = append(modifierObjects, &api.Modifier{
				TileId:   int32(tile),
				NewColor: color,
			})
		}
		player.stream <- &api.BoardState{
			Status: game.status(token),
			Tiles:  modifierObjects,
		}
	}
}
