package week08

import (
	"context"
	"fmt"
	"testing"
)

func Test_Warehouse_Start(t *testing.T) {
	t.Parallel()
	t.Skip()
	rootCtx := context.Background()

	exp, cancel := context.WithCancel(rootCtx)
	defer cancel()

	w := Warehouse{
		cap: 10,
		materials: Materials{
			Metal: 10,
		},
	}

	act := w.Start(rootCtx)

	fmt.Println(act.Err())
	fmt.Println(exp)

	//how can we actually compare contexts ?

}

func Test_Warehouse_Retreve(t *testing.T) {
	t.Parallel()

	w := Warehouse{
		cap: 10,
		materials: Materials{
			Metal: 10,
		},
	}

	exp := w.materials[Metal] - 5

	w.Retrieve(Metal, 5)

	act := w.materials[Metal]

	if act != exp {
		t.Fatalf("expected %v got %v", act, exp)
	}

}

func Test_Warehouse_fill(t *testing.T) {
	w := &Warehouse{}
	ctx := w.fill(Metal)
	exp := 10

	//listen until the context to signal warehouse is fillled.
	<-ctx.Done()

	if w.materials[Metal] != exp {
		t.Fatalf("expected %v got %v", exp, len(w.materials))
	}
}

func Test_Warehouse_Stop(t *testing.T) {
	w := &Warehouse{}

	w.Start(context.Background())

	w.Stop()

	if len(w.materials) != 0 && w.cap != 0 {
		t.Fatalf("expected w.material: %v got %v w.cap %v got %v", 0, 0, len(w.materials), w.cap)

	}
}
