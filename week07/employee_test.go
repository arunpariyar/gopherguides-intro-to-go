package week06

import (
	"context"
	"testing"
)

func Test_employee_IsValid(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		e    Employee
		exp  error
	}{
		{
			name: "valid case",
			e:    10,
			exp:  nil,
		},
		{
			name: "zero case",
			e:    0,
			exp:  ErrInvalidEmployee(0),
		},
		{
			name: "negative case",
			e:    -1,
			exp:  ErrInvalidEmployee(-1),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.e.IsValid()
			if act != nil {
				if act.Error() != tt.exp.Error() {
					t.Fatalf("%s: expected %q got %q", tt.name, tt.exp.Error(), act.Error())
				}
			}
		})
	}
}

func Test_Employee_Work_Error(t *testing.T) {
	t.Parallel()

	m := NewManager()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e := Employee(0)
	exp := ErrInvalidEmployee(0)

	go e.work(ctx, m)

	go func() {
		m.Assign(&Product{Quantity: 1})
	}()

	act := <-m.Errors()

	if act.Error() != exp.Error() {
		t.Fatalf("expected %q got %q", exp.Error(), act.Error())
	}
}

func Test_Employee_Work_Success(t *testing.T) {
	t.Parallel()
	m := NewManager()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	e := Employee(5)
	exp := CompletedProduct{
		Product: Product{
			Quantity: 10,
			builtBy:  5,
		},
		Employee: 5}

	go e.work(ctx, m)

	go func() {
		m.Assign(&Product{Quantity: 10})
	}()

	act := <-m.completed

	if act != exp {
		t.Fatalf("expected %q got %q", exp, act)
	}
}
