package week07

import (
	"context"
	"fmt"
	"testing"
)

// snippet: example
// TODO: Implement test cases for the Run function.
// func Test_Run(t *testing.T) {
// 	t.Parallel()

// 	// Tests you will need to write:

// 	// TODO: timeout after 5 seconds if nothing happens
// 	// TODO: interruption by a signal
// 	// TODO: Run returns when the products are completed
// 	// TODO: test that the output is correct

// }

// snippet: example

func Test_Start(t *testing.T) {
	
	
	m := NewManager()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	defer m.Stop()
	e := 1

	m.Start(ctx, e)

	// cancel()

	go func() {
		m.Assign(&Product{
			Quantity: 5,
		})
	}()

	

	select {
	case p := <-m.completed:
		fmt.Println(p)
	case e := <-m.errs:
		fmt.Println(e)
	}



}
