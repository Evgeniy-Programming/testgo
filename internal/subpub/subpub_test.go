package subpub

import (
    "context"
    "testing"
    "time"
)

func TestPubSub(t *testing.T) {
    bus := NewSubPub()
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Тест подписки и публикации
    sub, err := bus.Subscribe("test", func(msg interface{}) {
        // обработка сообщения
    })
    if err != nil {
        t.Fatalf("Subscribe failed: %v", err)
    }

    err = bus.Publish("test", "message")
    if err != nil {
        t.Fatalf("Publish failed: %v", err)
    }

    // Тест graceful shutdown
    go func() {
        time.Sleep(100 * time.Millisecond)
        cancel()
    }()

    err = bus.Close(ctx)
    if err != nil {
        t.Fatalf("Close failed: %v", err)
    }
}