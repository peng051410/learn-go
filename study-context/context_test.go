package studycontext

import (
	"context"
	"log"
	"testing"
	"time"
)


func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	go handle(ctx, 500 * time.Microsecond)

	select {
		case <-ctx.Done():
			log.Println("main exit",ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle exit")
	case <-time.After(duration):
		log.Println("process request with", duration)
	}
}

func TestContextWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	go handle(ctx, 1500 * time.Microsecond)

	select {
		case <-ctx.Done():
			log.Println("main exit",ctx.Err())
	}
}