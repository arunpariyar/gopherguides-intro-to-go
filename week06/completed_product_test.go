package week06

import (
	"testing"
)

func Test_completed_product_IsValid(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		cp   CompletedProduct
		exp  error
	}{
		{
			name: "all valid",
			cp: CompletedProduct{
				Product: Product{
					Quantity: 10,
					builtBy:  5,
				},
				Employee: 5,
			},
			exp: nil,
		},
		{
			name: "invalid employee",
			cp: CompletedProduct{
				Product: Product{
					Quantity: 10,
				},
				Employee: 0,
			},
			exp: ErrInvalidEmployee(0),
		},
		{
			name: "invalid product",
			cp: CompletedProduct{
				Product: Product{
					Quantity: 0,
				},
				Employee: 5,
			},
			exp: ErrInvalidQuantity(0),
		},
		{
			name: "product not built",
			cp: CompletedProduct{
				Product: Product{
					Quantity: 1,
				},
				Employee: 5,
			},
			exp: ErrProductNotBuilt("product is not built: {1 0}"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.cp.IsValid()
			if act != tt.exp {
				t.Fatalf("expected %q got %q", act, tt.exp)
			}
		})
	}

}
