package server

import (
	"context"

	"github.com/minish144/go-sms-api/gen/pb"
)

type ApiServiceServer struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) SendMessage(ctx context.Context, in *pb.Messages_SendRequest) (*pb.Messages_SendResponse, error) {
	return nil, nil
}
