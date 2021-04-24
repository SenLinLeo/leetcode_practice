package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/**
    1. 并发交给调用者
    2. goroutinue什么时候退出
    3. 能够控制退出
**/
func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	fmt.Printf("goroutinue数量1：%d\n", runtime.NumGoroutine())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	fmt.Printf("goroutinue数量2：%d\n", runtime.NumGoroutine())
	defer cancel()
	tr.Shutdown(ctx)
}

func NewTracker() *Tracker {
	return &Tracker{
		ch : make(chan string, 10),
	}
}

type Tracker struct {
	ch chan string
	stop chan struct{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	    case t.ch <- data:
			return nil
		case <-ctx.Done():
			return ctx.Err()
	}
}

// 消费
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}

	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <- t.stop:
		case <-ctx.Done():
	}
}