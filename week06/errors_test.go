package week06

import (
	"testing"
)

func Test_All_Errors(t *testing.T) {
	table := []struct {
		name string
		e    error
		exp  string
	}{
		{
			name: "invalid quantity error",
			e:    ErrInvalidQuantity(-1),
			exp:  "quantity must be greater than 0, got -1",
		},
		{
			name: "product not built error",
			e:    ErrProductNotBuilt("error"),
			exp:  "error",
		},
		{
			name: "invalid employee error",
			e:    ErrInvalidEmployee(0),
			exp:  "invalid employee number: 0",
		},
		{
			name: "invalid employee count",
			e:    ErrInvalidEmployeeCount(0),
			exp:  "invalid employee count: 0",
		},
		{
			name: "manager stopped error",
			e:    ErrManagerStopped{},
			exp:  "manager is stopped",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.e.Error()

			if act != tt.exp {
				t.Fatalf("expected %q got %q", tt.exp, act)
			}
		})
	}
}
