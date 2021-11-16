package week08

import (
	"testing"
)

func Test_Completed_Product_IsValid(t *testing.T) {
	table := []struct {
		name string
		cp   CompletedProduct
		exp  error
	}{
		{
			name: "valid case",
			cp: CompletedProduct{
				Product: Product{
					Materials: Materials{
						Metal:   1,
						Oil:     2,
						Plastic: 3,
						Wood:    4,
					},
					builtBy: 10,
				},
				Employee: 10,
			},
			exp: nil,
		},
		{
			name: "employee invalid case",
			cp: CompletedProduct{
				Product: Product{
					Materials: Materials{},
					builtBy: 10,
				},
				Employee: 10,
			},
			exp: ErrInvalidMaterials(0),
		},
		{
			name: "product not built",
			cp: CompletedProduct{
				Product: Product{
					Materials: Materials{
						Metal:   1,
						Oil:     2,
						Plastic: 3,
						Wood:    4,
					},	
				},
				Employee: 10,
			},
			exp: ErrProductNotBuilt("product is not built: [{metal:1x}, {oil:2x}, {plastic:3x}, {wood:4x}]"),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.cp.IsValid()
			if act != tt.exp{
				t.Fatalf("expected %v got %v",tt.exp, act)
			}
		})
	}
}
