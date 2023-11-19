package testdata

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 监听Context的取消信号
			fmt.Println("Worker received cancellation signal")
			return
		default:
			// 执行工作任务
			fmt.Println("Worker is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestContext(t *testing.T) {

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	WithVal(context.WithValue(ctx, Val1{}, "hhhhh"))

	go worker(ctx)

	time.Sleep(3 * time.Second)

	cancel() // 发送取消信号给Context

	time.Sleep(1 * time.Second)

}

func WithVal(ctx context.Context) {
	value := ctx.Value(Val1{}).(string)
	fmt.Println(value)

}

type Val1 struct {
	Name string
}
