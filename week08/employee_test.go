package week08

import (
	"context"
	"testing"
	"time"
)

func Test_Employee_isValid(t *testing.T) {
	table := []struct {
		name string
		e    Employee
		exp  error
	}{
		{
			name: "success",
			e:    5,
			exp:  nil,
		},
		{
			name: "error",
			e:    0,
			exp:  ErrInvalidEmployee(0),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.e.IsValid()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Employee_Work_WithTimeOutContext(t *testing.T) {
	//create a new manager
	m := &Manager{}
	e := Employee(5)
	//create a context with timeOut cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	//set m.cancel to cancel
	m.cancel = cancel
	defer m.cancel()

	e.work(ctx, m)

	v := <-ctx.Done()

	if v != struct{}{} {
		t.Fatalf("expected %v got %v", v, struct{}{})
	}
}

func Test_Employee_Work_Success(t *testing.T) {

	m := &Manager{}
	m.Warehouse = &Warehouse{
		materials: Materials{
			Metal: 5,
			// Oil:     2,
			Plastic: 5,
			// Wood:    4,
		},
		cap: 10,
	}

	count := 5

	ctx := context.Background()
	p := &Product{
		Materials: Materials{
			Oil:  2,
			Wood: 2,
		},
	}

	go m.Start(ctx, count)

	go m.Assign(p)

	cp := <-m.completedCh()

	if cp.Employee != p.builtBy {
		t.Fatalf("expected %v got %v", p.builtBy, cp.Employee)
	}
}
