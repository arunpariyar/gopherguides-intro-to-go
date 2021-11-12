package week06

import (
	"testing"
)

func Test_Manager_Start_Fail(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()

	exp := ErrInvalidEmployeeCount(0)

	act := m.Start(0)

	if act.Error() != exp.Error() {
		t.Fatalf("expected %q got %q", exp.Error(), act.Error())
	}
}

func Test_Manager_Start_Success(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()

	exp := 10

	err := m.Start(3)

	if err == nil {
		go func() {
			m.Assign(&Product{Quantity: 10})
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

	exp := ErrManagerStopped{}

	//stopping the manager
	m.Stop()

	act := m.Assign(&Product{})

	if act.Error() != exp.Error() {
		t.Fatalf("expected %q got %q", exp.Error(), act.Error())
	}
}

func Test_Manager_Assign_Success(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()

	p1 := &Product{Quantity: 1}
	p2 := &Product{Quantity: 2}
	p3 := &Product{Quantity: 3}

	exp := 6
	act := 0

	go func() {
		m.Assign(p1, p2, p3)
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

	e := Employee(3)
	exp := CompletedProduct{
		Product: Product{
			Quantity: 10,
			builtBy:  e,
		},
		Employee: e,
	}

	go e.work(m)

	go func() {
		m.Assign(&Product{Quantity: 10})
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
