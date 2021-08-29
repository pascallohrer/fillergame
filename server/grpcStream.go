package server

import (
	"github.com/pascallohrer/fillergame/api"
)

type grpcStream interface {
	Send(*api.BoardState) error
}

// internal convenience method to send a status message with empty modifier list to a grpc stream
func streamStatus(stream grpcStream, status *api.Status) error {
	errorBoard := &api.BoardState{
		Status: status,
		Tiles:  []*api.Modifier{},
	}
	if err := stream.Send(errorBoard); err != nil {
		return err
	}
	return nil
}
