package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pascallohrer/fillergame/api"
)

// Server is the implementation of the GRPC server interface.
// It links GRPC methods to the corresponding internal methods on the engine.
// It also catches some errors that can be recognized from outside the engine.
type Server struct {
	currentGame *fillerGame
}

func (s *Server) JoinGame(request *api.JoinRequest, stream api.FillerGame_JoinGameServer) error {
	token, err := createToken(request.Username, request.PasswordHash)
	if err != nil {
		return streamStatus(stream, &api.Status{
			StatusCode: api.StatusCode_DEFAULT,
			Message:    fmt.Sprintf("Authentication Error: %v", err),
		})
	}
	if s.currentGame == nil {
		s.currentGame = newGame(request.Mode)
	}
	if err := s.currentGame.validateMode(request.Mode); err != nil {
		return streamStatus(stream, &api.Status{
			StatusCode: api.StatusCode_UNAVAILABLE_MODE,
			Message:    fmt.Sprintf("Invalid Mode: %v", err),
		})
	}
	gameStream := make(chan *api.BoardState)
	if err := s.currentGame.addPlayer(token, gameStream); err != nil {
		return streamStatus(stream, &api.Status{
			StatusCode: api.StatusCode_DEFAULT,
			Message:    fmt.Sprintf("Unable to join: %v", err),
		})
	}
	err = streamStatus(stream, &api.Status{
		StatusCode: api.StatusCode_AUTH_TOKEN,
		Message:    token,
	})
	if err != nil {
		return err
	}
	fmt.Printf("New player joined the game. token: %v\n", token)
	for message := range gameStream {
		if err := stream.Send(message); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) MakeMove(ctx context.Context, request *api.MoveRequest) (*api.MoveAcknowledgment, error) {
	s.currentGame.processMove(request.AuthToken, request.NewColor)
	return &api.MoveAcknowledgment{}, nil
}

func (s *Server) RebuildBoard(request *api.RebuildRequest, stream api.FillerGame_RebuildBoardServer) error {
	player, exists := s.currentGame.players[request.AuthToken]
	if !exists {
		return streamStatus(stream, &api.Status{
			StatusCode: api.StatusCode_INVALID_TOKEN,
			Message:    "Your token was not recognized. Please use the correct token.",
		})
	}
	currentBoard := &api.BoardState{
		Status: s.currentGame.status(request.AuthToken),
		Tiles:  s.currentGame.boardState(request.AuthToken),
	}
	if err := stream.Send(currentBoard); err != nil {
		return err
	}
	gameStream := make(chan *api.BoardState)
	player.setStream(gameStream)
	for message := range gameStream {
		if err := stream.Send(message); err != nil {
			return err
		}
	}
	return nil
}

// PrintDebug is supposed to run as a goroutine during server operation. It reads command-line commands to print debug information.
func (s *Server) PrintDebug() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if err := input.Err(); err != nil {
			log.Fatalf("PrintDebug: command couldn't be read: %v", err)
			return
		}
		switch input.Text() {
		case "d":
			fmt.Println(s.currentGame)
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("currentGame not set.")
					go s.PrintDebug()
				}
			}()
			fmt.Println(s.currentGame.gameEngine)
		}
	}
}
