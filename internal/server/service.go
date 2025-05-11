package server

import (
    "context"
    "log"

    pb "github.com/yourusername/pubsub-service/pkg/pb"
    "github.com/yourusername/pubsub-service/internal/subpub"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type Server struct {
    pb.UnimplementedPubsubServer
    bus subpub.SubPub
}

func (s *Server) Subscribe(req *pb.SubscribeRequest, stream pb.Pubsub_SubscribeServer) error {
    // ... реализация streaming подписки
}

func (s *Server) Publish(ctx context.Context, req *pb.PublishRequest) (*emptypb.Empty, error) {
    // ... реализация публикации
}