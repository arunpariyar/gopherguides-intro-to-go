package week08

import (
	"context"
	"testing"
)

func Test_Manager_Start(t *testing.T) {
	t.Skip()
	table := []struct {
		name    string
		m       *Manager
		count   int
		rootCtx context.Context
		exp     error
	}{
		{
			name:    "success",
			m:       &Manager{},
			count:   5,
			rootCtx: context.Background(),
			exp:     nil,
		},
		{
			name:    "with error",
			m:       &Manager{},
			count:   0,
			rootCtx: context.Background(),
			exp:     ErrInvalidEmployeeCount(0),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.m.Start(tt.rootCtx, tt.count)
			if err != tt.exp {
				//stopping manager before fatal to avoid any resource leakage∏
				tt.m.Stop()
				t.Fatalf("expected %v got %v", tt.exp, err)

			}
		})
	}
}

func Test_Manager_Assign_Stopped(t *testing.T) {
	m := &Manager{}
	p := &Product{
		Materials: Materials{
			Metal:   1,
			Oil:     2,
			Plastic: 3,
			Wood:    4,
		},
	}
	exp := ErrManagerStopped{}
	m.Start(context.Background(), 1)

	m.Stop()
	act := m.Assign(p)

	if act != exp {
		t.Fatalf("expected %v got %v", exp.Error(), act.Error())
	}
}

func Test_Manager_Assign(t *testing.T) {
	table := []struct {
		name    string
		m       *Manager
		ps      *Product
		rootCtx context.Context
		count   int
		exp     error
	}{
		{
			name: "invalid materials",
			m:    &Manager{},
			ps: &Product{
				Materials: Materials{},
			},
			count:   5,
			rootCtx: context.Background(),
			exp:     ErrInvalidMaterials(0),
		},
		{
			name: "invalid materials",
			m:    &Manager{},
			ps: &Product{
				Materials: Materials{
					Metal:   1,
					Oil:     2,
					Plastic: 3,
					Wood:    4,
				},
			},
			count:   5,
			rootCtx: context.Background(),
			exp:     nil,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Start(tt.rootCtx, tt.count)

			act := tt.m.Assign(tt.ps)

			if act != tt.exp {
				tt.m.Stop()
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Manager_Complete_Success(t *testing.T) {
	m := &Manager{}
	e := Employee(1)
	p := &Product{
		Materials: Materials{
			Metal:   1,
			Oil:     2,
			Plastic: 3,
			Wood:    4,
		},
		builtBy: 5,
	}

	exp := CompletedProduct{
		Employee: 1,
		Product: Product{
			Materials: Materials{
				Metal:   1,
				Oil:     2,
				Plastic: 3,
				Wood:    4,
			},
			builtBy: 1,
		},
	}

	go func() {
		m.Start(context.Background(), 5)
		m.Complete(e, p)
	}()

	//channel to listen for completed products
	cc := m.Completed()
	//receive from channel
	act := <-cc
	//checking to make sure that the Employee was reassigned
	if act.Employee != exp.Employee {
		t.Fatalf("expected %v got %v", exp.Employee, act.Employee)
	}

}
func Test_Manager_Complete_Fail(t *testing.T) {
	table := []struct {
		name string
		m    *Manager
		e    Employee
		p    *Product
		exp  error
	}{
		{
			name: "employee not valid",
			m:    &Manager{},
			e:    0,
			p: &Product{
				Materials: Materials{
					Metal: 1,
				},
			},
			exp: ErrInvalidEmployee(0),
		},
		{
			name: "product not built",
			m:    &Manager{},
			e:    1,
			p: &Product{
				Materials: Materials{
					Metal: 1,
				},
			},
			exp: ErrProductNotBuilt("product is not built: [{metal:1x}]"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.m.Complete(tt.e, tt.p)

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}
