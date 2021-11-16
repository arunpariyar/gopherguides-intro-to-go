package week08

import (
	"context"
	"fmt"
	"testing"
)

// snippet: example
// ATTENTION: YOU ARE NOT ALLOWED TO SUBMIT THIS
// TEST AS PART OF YOUR ASSIGNMENT!!
//
// This test is meant to demonstrate how to use the
// application.
//
// YOU MUST DELETE THIS TEST !!BEFORE!! YOU SUBMIT!!
func Test_Manager_Demonstration(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	m := &Manager{}

	ctx, err := m.Start(ctx, 5)
	if err != nil {
		t.Fatal("unexpected error")
	}

	for i := 0; i < 10; i++ {
		go m.Assign(ProductA)
		go m.Assign(ProductB)
	}

	var completed []CompletedProduct

	go func() {
		fmt.Println("waiting for a completed product")

		for cp := range m.Completed() {
			completed = append(completed, cp)

			if len(completed) >= 20 {
				m.Stop()
			}
		}
	}()

	fmt.Println("waiting for the ctx to be cancelled")
	<-ctx.Done()

	fmt.Println("validating output")
	if len(completed) != 20 {
		t.Fatalf("got %v, expected %v", len(completed), 20)
	}
	fmt.Println("validated")
}

// snippet: example
