package client

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/pascallohrer/fillergame/api"

	"google.golang.org/grpc"
)

const title = " Welcome to FillerGame!\n" +
	" Please enter your username and password, separated by one space:\n"

var (
	status string
	token  string
)

// Run executes the client by trying to connect to host, setting up, then running the game loop
func Run(host string) error {
	connection, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer connection.Close()
	client := api.NewFillerGameClient(connection)
	fmt.Print(title)
	field := field{}
	err = inputLoop(func(command string) error {
		return execute(command, client, &field)
	})
	if err != nil {
		return fmt.Errorf("Error processing input: %v", err)
	}
	return nil
}

func execute(command string, client api.FillerGameClient, currentField *field) error {
	commandParts := strings.SplitN(command, " ", -1)
	switch status {
	case "":
		username := commandParts[0]
		password := commandParts[1]
		stream, err := client.JoinGame(context.Background(), &api.JoinRequest{
			Mode:         api.GameBoard_SQUARE10,
			Username:     username,
			PasswordHash: password,
		})
		if err != nil {
			return fmt.Errorf("Unable to join game: %v", err)
		}
		fmt.Println(" JoinRequest sent to server. Waiting for response...")
		response, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("Error reading first server response: %v", err)
		}
		if response.Status.StatusCode != api.StatusCode_AUTH_TOKEN {
			return fmt.Errorf("Server error. Expected token, got %v", response.Status.Message)
		}
		token = response.Status.Message
		fmt.Printf(" Joined game. Token: %v\n", token)
		status = "waiting"
		go readStream(stream, currentField)
	case "waiting":
		fmt.Println(" It's not your turn!")
	case "move":
		newColor, valid := colors[commandParts[0]]
		if !valid {
			return nil
		}
		status = "waiting"
		client.MakeMove(context.Background(), &api.MoveRequest{
			AuthToken: token,
			NewColor:  newColor,
		})
	default:
		return fmt.Errorf("Unexpected status: %v", status)
	}
	return nil
}

func readStream(stream api.FillerGame_JoinGameClient, currentField *field) error {
	for {
		boardState, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Error reading stream: %v", err)
		}
		if boardState.Status.StatusCode == api.StatusCode_YOUR_TURN ||
			boardState.Status.StatusCode == api.StatusCode_ILLEGAL_MOVE {
			status = "move"
		}
		for _, modifier := range boardState.Tiles {
			currentField.update(int(modifier.TileId), modifier.NewColor)
		}
		printUpdate(currentField, boardState.Status.Message)
	}
}
