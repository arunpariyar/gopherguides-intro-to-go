package week08

import (
	"testing"
)

func Test_Errors_All(t *testing.T) {
	table := []struct {
		name string
		e    error
		exp  string
	}{
		{
			name: "invalid materials",
			e:    ErrInvalidMaterials(0),
			exp:  "materials must be greater than 0, got 0",
		},
		{
			name: "product not built",
			e:    ErrProductNotBuilt("test"),
			exp:  "test",
		},
		{
			name: "invalid employee",
			e:    ErrInvalidEmployee(0),
			exp:  "invalid employee number: 0",
		},
		{
			name: "invalid employee count",
			e:    ErrInvalidEmployeeCount(0),
			exp:  "invalid employee count: 0",
		},
		{
			name: "manager stopped",
			e:    ErrManagerStopped{},
			exp:  "manager is stopped",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.e.Error()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}

		})
	}
}
