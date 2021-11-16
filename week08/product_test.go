package week08

import (
	"testing"
)

func Test_Product_String(t *testing.T) {
	table := []struct {
		name string
		p    *Product
		exp  string
	}{
		{
			name: "ProductA",
			p:    ProductA,
			exp:  "[{oil:3x}, {wood:2x}]",
		},
		{
			name: "Empty",
			p:    &Product{},
			exp:  "[]",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.String()
			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Product_BuiltBy(t *testing.T) {
	table := []struct {
		name string
		p    *Product
		exp  Employee
	}{
		{
			name: "built",
			p: &Product{
				Materials: Materials{
					Metal:   1,
					Oil:     2,
					Plastic: 3,
					Wood:    4,
				},
				builtBy: 10,
			},
			exp: 10,
		},
		{
			name: "not built",
			p:    &Product{},
			exp:  0,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.BuiltBy()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Product_Build(t *testing.T) {
	table := []struct {
		name string
		p    *Product
		e    Employee
		w    *Warehouse
		exp  error
	}{
		{
			name: "valid case",
			e:    2,
			p:    ProductA,
			w:    &Warehouse{},
			exp:  nil,
		},
		{
			name: "product invalid",
			e:    2,
			p:    &Product{},
			w:    &Warehouse{},
			exp:  ErrInvalidMaterials(0),
		},
		{
			name: "employee invalid",
			e:    0,
			p:    ProductA,
			w:    &Warehouse{},
			exp:  ErrInvalidEmployee(0),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.Build(tt.e, tt.w)
			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Product_IsBuilt(t *testing.T) {
	table := []struct {
		name string
		p    *Product
		exp  error
	}{
		{
			name: "valid case",
			p: &Product{
				Materials: Materials{
					Metal:   1,
					Oil:     2,
					Plastic: 3,
					Wood:    4,
				},
				builtBy: 10,
			},
			exp: nil,
		},
		{
			name: "invalid product",
			p: &Product{
				Materials: Materials{},
				builtBy:   10,
			},
			exp: ErrInvalidMaterials(0),
		},
		{
			name: "not built",
			p: &Product{
				Materials: Materials{
					Metal:   1,
					Oil:     2,
					Plastic: 3,
					Wood:    4,
				},
			},
			exp: ErrProductNotBuilt("product is not built: [{metal:1x}, {oil:2x}, {plastic:3x}, {wood:4x}]"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.IsBuilt()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}
