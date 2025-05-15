package goevent

import (
	"context"
	"fmt"
	"testing"
)

func TestEvent(t *testing.T) {

	Event().Register("test", func(ctx context.Context, args ...interface{}) {

		fmt.Printf("event callback is runing: %#v", args[0])
	})

	Event().Call("test", context.Background(), []int{1, 2, 3, 4, 5})
}
