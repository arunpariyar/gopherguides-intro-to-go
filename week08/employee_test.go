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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	m.cancel = cancel
	defer m.cancel()

	e.work(ctx, m)
	//listen for done
	v := <-ctx.Done()
	// compare to an empty struct
	if v != struct{}{} {
		t.Fatalf("expected %v got %v", v, struct{}{})
	}
}

func Test_Employee_Work_Success(t *testing.T) {
	m := &Manager{}
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
	//listen for the completed product
	cp := <-m.completedCh()

	if cp.Employee != p.builtBy {
		t.Fatalf("expected %v got %v", p.builtBy, cp.Employee)
	}
}

func Test_Employee_Work_Error(t *testing.T) {
	m := &Manager{}
	count := 5
	ctx := context.Background()
	p := &Product{
		Materials: Materials{},
	}

	e := Employee(5)
	exp := ErrInvalidMaterials(0)

	ctx, err := m.Start(ctx, count)
	if err != nil {
		t.Fatalf("expected %v got %v", nil, err)
	}

	go e.work(ctx, m)

	m.Jobs() <- p
	//listen for the error
	select {
	case act := <-m.Errors():
		if act != exp {
			t.Fatalf("expected %v got %v", exp, act)
		}
	case <-ctx.Done():
	}

}
