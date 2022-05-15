package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx1 := context.Background()
	select {
	case <-ctx1.Done():
		fmt.Println("----------------")
	default:
	}

	ctx, cancel := context.WithCancel(context.Background())

	ctxt := context.WithValue(ctx, "han", "tuo")
	fmt.Println(ctxt.Value("han"))
	cancel()

	fmt.Println(ctxt.Value("han"))
}
