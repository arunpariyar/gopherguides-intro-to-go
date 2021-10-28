package demo

import "testing"

func Test_Add(t *testing.T) {
	tests := map[string]struct {
		a    int
		b    int
		want int
	}{
		"simple": {a: 1, b: 2, want: 3},
		"negative":    {a: -2, b: 5, want: 3},
		// "zero values": {a: 0, b: 0, want: 0},
	}

	for name, tc := range tests {
		got, err := Add(tc.a, tc.b)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Fatalf("%q expected %d got %d", name, tc.want, got)
		}
	}
}
