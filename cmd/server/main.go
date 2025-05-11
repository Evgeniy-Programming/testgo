package main

import (
    "context"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/yourusername/pubsub-service/internal/server"
    "github.com/yourusername/pubsub-service/internal/subpub"
    "google.golang.org/grpc"
)

func main() {
    // Загрузка конфигурации
    cfg, err := server.LoadConfig("configs/config.yaml")
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    // Инициализация шины событий
    bus := subpub.NewSubPub()

    // Создание gRPC сервера
    srv := server.NewServer(bus)
    grpcServer := grpc.NewServer()
    pb.RegisterPubsubServer(grpcServer, srv)

    // Graceful shutdown
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        sig := make(chan os.Signal, 1)
        signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
        <-sig
        cancel()
    }()

    // Запуск сервера
    lis, err := net.Listen("tcp", ":"+cfg.Server.Port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Printf("server listening at %v", lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}