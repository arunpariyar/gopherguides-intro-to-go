package week06

import (
	"context"
	"os/signal"
	"syscall"
	"testing"
	"time"
)



func Test_Manager_Start_Fail(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	exp := ErrInvalidEmployeeCount(0)

	act := m.Start(ctx, 0)

	if act.Error() != exp.Error() {
		t.Fatalf("expected %q got %q", exp.Error(), act.Error())
	}
}

func Test_Manager_Start_Success(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	exp := 10

	err := m.Start(ctx, 3)

	if err == nil {
		go func() {
			m.Assign(ctx, &Product{Quantity: 10})
		}()
		act := <-m.completed

		if act.Product.Quantity != 10 {
			t.Fatalf("expected %q got %q", act.Product.Quantity, exp)
		}
	}
}

func Test_Manager_Assign_Stopped(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()
	exp := ErrManagerStopped{}
	ctx := context.Background()
	//stopping the manager
	m.Stop()

	act := m.Assign(ctx,&Product{})

	if act.Error() != exp.Error() {
		t.Fatalf("expected %q got %q", exp.Error(), act.Error())
	}
}

func Test_Manager_Assign_Success(t *testing.T) {
	t.Parallel()
	m := NewManager()

	p1 := &Product{Quantity: 1}
	p2 := &Product{Quantity: 2}
	p3 := &Product{Quantity: 3}
	ctx := context.Background()
	exp := 6
	act := 0

	go func() {
		m.Assign(ctx, p1, p2, p3)
		close(m.jobs)
	}()
	//aggregating the total product quantity to compare with total products quatity "exp"
	for v := range m.jobs {
		act += v.Quantity
	}
	if act != exp {
		t.Fatalf("expected %v got %v", exp, act)
	}
}

func Test_Manager_Complete_Success(t *testing.T) {
	t.Parallel()
	m := NewManager()

	e := Employee(5)
	p := &Product{
		Quantity: 10,
		builtBy:  e,
	}
	exp := CompletedProduct{
		Product: Product{
			Quantity: 10,
			builtBy:  e,
		},
		Employee: e,
	}

	go func() {
		m.Complete(e, p)
		m.Stop()
	}()

	act := <-m.completed
	if act != exp {
		t.Fatalf("expected %q got %q", exp, act)
	}
}

func Test_Manager_Complete_Fail(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		p    *Product
		e    Employee
		exp  error
	}{
		{
			name: "Invalid Employee Error",
			p: &Product{
				Quantity: 10,
			},
			e:   Employee(0),
			exp: ErrInvalidEmployee(0),
		},
		{
			name: "Product not built error",
			p: &Product{
				Quantity: 10,
			},
			e:   Employee(1),
			exp: ErrProductNotBuilt("product is not built: {10 0}"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager()
			defer m.Stop()

			act := m.Complete(tt.e, tt.p)
			if act.Error() != tt.exp.Error() {
				t.Fatalf("expected %v got %v", tt.exp.Error(), act.Error())
			}
		})
	}
}

func Test_Manager_Completed(t *testing.T) {
	t.Parallel()

	m := NewManager()
	defer m.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e := Employee(3)
	exp := CompletedProduct{
		Product: Product{
			Quantity: 10,
			builtBy:  e,
		},
		Employee: e,
	}

	go e.work(ctx, m)

	go func() {
		m.Assign(ctx, &Product{Quantity: 10})
	}()

	act := <-m.completedCh()

	if act != exp {
		t.Fatalf("expected %q got %q", exp, act)
	}
}

func Test_Manager_Done(t *testing.T) {
	m := NewManager()

	exp := true

	m.Stop()
	act := m.stopped

	if act != exp {
		t.Fatalf("expected %t got %t", exp, act)
	}
}

func Test_Run_Successful_Output_Tested(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	count := 5
	p := []*Product{
		&Product{Quantity: 1},
		&Product{Quantity: 2},
		&Product{Quantity: 3},
		&Product{Quantity: 4},
		&Product{Quantity: 5},
	}

	act, err := Run(ctx, count, p...)

	if err != nil {
		t.Fatalf("expected %v got %v", nil, act)
	}

	if len(act) != 5 {
		t.Fatalf("expected %v got %v", count, len(act))
	}

}

func Test_Run_Error(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	count := 5
	p := []*Product{
		&Product{Quantity: 1},
		&Product{Quantity: 2},
		&Product{Quantity: 3},
		&Product{Quantity: 4},
		&Product{Quantity: 0},
	}
	exp := ErrInvalidQuantity(0)

	_, err := Run(ctx, count, p...)

	if err.Error() != exp.Error() {
		t.Fatalf("expected %v got %v", exp.Error(), err.Error())
	}
}

func Test_Run_With_TimeOut(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	p := &Product{Quantity: 25000}
	count := 1

	Run(ctx, count , p)

	<-ctx.Done()

	if ctx.Err().Error() != "context deadline exceeded" {
		t.Fatalf("expected %v got %v", "context deadline exceeded", ctx.Err().Error())
	}

}

func Test_Run_Interrupted_Signal(t *testing.T){
	t.Parallel()

	const TEST_SIGNAL = syscall.SIGUSR2
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	sigCtx, cancel := signal.NotifyContext(ctx, TEST_SIGNAL)
	defer cancel()

	p := &Product{Quantity: 300000}
	count := 1

	go Run(sigCtx, count, p)

	go func(){
		time.Sleep(time.Second)
		syscall.Kill(syscall.Getpid(), TEST_SIGNAL)
	}()

	exp := context.Canceled.Error()

	<-sigCtx.Done()
	
	if  sigCtx.Err().Error() != exp {
		t.Fatalf("expected: %v got %v", exp, sigCtx.Err().Error())
	}

}
