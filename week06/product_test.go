package week06

import (
	"testing"
)

func Test_Product_BuiltBy(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		p    Product
		exp  Employee
	}{
		{
			name: "with value",
			p: Product{
				Quantity: 10,
				builtBy:  5,
			},
			exp: 5,
		},
		{
			name: "zero value",
			p:    Product{},
			exp:  0,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.BuiltBy()
			if act != tt.exp {
				t.Fatalf("expected %q got %q", tt.exp, act)
			}
		})
	}
}

func Test_Product_Build(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		p    *Product
		e    Employee
		exp  error
	}{
		{
			name: "happy case",
			p: &Product{
				Quantity: 10,
			},
			e:   5,
			exp: nil,
		},
		{
			name: "no product",
			p: &Product{
				Quantity: 0,
			},
			e:   5,
			exp: ErrInvalidQuantity(0),
		},
		{
			name: "no employee",
			p: &Product{
				Quantity: 5,
			},
			e:   0,
			exp: ErrInvalidEmployee(0),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.Build(tt.e)

			if act != tt.exp {
				t.Fatalf("expected %q got %q", tt.exp, act)
			}
		})
	}
}

func Test_Product_IsValid(t *testing.T) {
	t.Parallel()
	table := []struct {
		name string
		p    Product
		exp  error
	}{
		{
			name: "valid case",
			p: Product{
				Quantity: 10,
			},
			exp: nil,
		},
		{
			name: "zero case",
			p: Product{
				Quantity: 0,
			},
			exp: ErrInvalidQuantity(0),
		},
		{
			name: "negative case",
			p: Product{
				Quantity: -1,
			},
			exp: ErrInvalidQuantity(-1),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.IsValid()

			if act != tt.exp {
				t.Fatalf("expected %q got %q", tt.exp, act)
			}
		})
	}
}

func Test_Product_IsBuilt(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		p    Product
		exp  error
	}{
		{
			name: "happy case",
			p: Product{
				Quantity: 10,
				builtBy:  5,
			},
			exp: nil,
		},
		{
			name: "invalid quantity",
			p: Product{
				Quantity: 0,
			},
			exp: ErrInvalidQuantity(0),
		},
		{
			name: "product not built",
			p: Product{
				Quantity: 1,
			},
			exp: ErrProductNotBuilt("product is not built: {1 0}"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.p.IsBuilt()

			if act != tt.exp {
				t.Fatalf("expected %q got %q", tt.exp, act)
			}
		})
	}
}
