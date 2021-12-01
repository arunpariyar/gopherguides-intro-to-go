package week08

import (
	"testing"
	"time"
)

func Test_Material_String(t *testing.T) {
	t.Parallel()

	var Mud Material = "mud"

	table := []struct {
		name string
		m    Material
		exp  string
	}{
		{
			name: "metal",
			m:    Metal,
			exp:  "metal",
		},
		{
			name: "oil",
			m:    Oil,
			exp:  "oil",
		}, {
			name: "mud",
			m:    Mud,
			exp:  "mud",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.m.String()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Material_Duration(t *testing.T) {
	t.Parallel()
	var Mud Material = "mud"
	table := []struct {
		name string
		m    Material
		exp  time.Duration
	}{
		{
			name: "metal",
			m:    Metal,
			exp:  5 * time.Millisecond,
		},
		{
			name: "plastic",
			m:    Plastic,
			exp:  7 * time.Millisecond,
		},
		{
			name: "mud",
			m:    Mud,
			exp:  3 * time.Millisecond,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.m.Duration()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}

func Test_Materials_Duration(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		ms   Materials
		exp  time.Duration
	}{
		{
			name: "one item",
			ms: Materials{
				Metal: 5,
			},
			exp: 25 * time.Millisecond,
		},
		{
			name: "three item",
			ms: Materials{
				Metal:   10,
				Plastic: 20,
				Wood:    10,
			},
			exp: 230 * time.Millisecond,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.ms.Duration()

			if act != tt.exp {
				t.Fatalf("expected %v got %v", tt.exp, act)
			}
		})
	}
}
