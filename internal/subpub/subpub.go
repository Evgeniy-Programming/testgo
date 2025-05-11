package subpub

import (
    "context"
    "sync"
)

type subPub struct {
    mu          sync.RWMutex
    subscribers map[string][]*subscriber
    closed      bool
}

type subscriber struct {
    ch     chan interface{}
    ctx    context.Context
    cancel context.CancelFunc
}

func NewSubPub() SubPub {
    return &subPub{
        subscribers: make(map[string][]*subscriber),
    }
}

func (s *subPub) Subscribe(subject string, cb MessageHandler) (Subscription, error) {
    // ... реализация подписки
}

func (s *subPub) Publish(subject string, msg interface{}) error {
    // ... реализация публикации
}

func (s *subPub) Close(ctx context.Context) error {
    // ... реализация graceful shutdown
}