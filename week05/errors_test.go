package demo

import (
	"testing"
)

func Test_errors_Error(t *testing.T) {
	t.Parallel()
	e := ErrTableNotFound{
		table: "test",
	}

	exp := `table not found test`

	if ok := IsErrTableNotFound(e); ok {

		act := e.Error()

		if act != exp {
			t.Errorf("expected %q got %q", exp, act)
		}
	}
}

func Test_errors_Table_Not_Found(t *testing.T) {
	t.Parallel()
	e := ErrTableNotFound{
		table: "test",
	}

	exp := "test"

	if ok := IsErrTableNotFound(e); ok {

		act := e.TableNotFound()

		if act != exp {
			t.Errorf("expected %q got %q", exp, act)
		}
	}
}

func Test_errors_ErrTableNotFound_Is(t *testing.T) {
	t.Parallel()
	e := ErrTableNotFound{
		table: "test",
	}
	exp := true

	act := e.Is(ErrTableNotFound{})
	if act != true {
		t.Fatalf("expected %t got %t", exp, act)
	}
}

func Test_errors_Clauses(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		e    errNoRows
		exp  Clauses
	}{
		{
			name: "not empty",
			e: errNoRows{
				clauses: Clauses{"gps": "garmin"},
				table:   "orders",
			},
			exp: Clauses{"gps": "garmin"}},
		{
			name: "empty",
			e: errNoRows{
				clauses: Clauses{},
				table:   "orders",
			},
			exp: Clauses{}},
		{
			name: "clause nil",
			e: errNoRows{
				clauses: nil,
				table:   "orders",
			},
			exp: Clauses{}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.e.Clauses()

			if act.String() != tt.exp.String() {
				t.Fatalf("%s expected %#v got %#v ", tt.name, tt.exp.String(), act.String())
			}
		})
	}
}

func Test_errors_RowNotFound(t *testing.T) {
	t.Parallel()

	cls1 := Clauses{"gps": "garmin"}
	tn1 := "orders"

	e1 := errNoRows{
		clauses: cls1,
		table:   tn1,
	}

	tbl, cls := e1.RowNotFound()

	if tbl != tn1 {
		t.Fatalf("expected %#v got %#v", tn1, tbl)
	}

	if cls1.String() != cls.String() {
		t.Fatalf("expected %#v got %#v", cls1.String(), cls.String())
	}
}

func Test_errors_errNoRows_Is(t *testing.T) {
	t.Parallel()
	cls1 := Clauses{"gps": "garmin"}
	tn1 := "orders"

	e := errNoRows{
		clauses: cls1,
		table:   tn1,
	}
	exp := true

	act := e.Is(&errNoRows{})
	if act != exp {
		t.Fatalf("expected %t got %t", exp, act)
	}
}

func Test_errors_isErrNoRows(t *testing.T) {
	t.Parallel()
	table := []struct {
		name string
		err  error
		exp  bool
	}{
		{
			name: "true case",
			err:  &errNoRows{},
			exp:  true},
		{
			name: "false case",
			err:  &ErrTableNotFound{},
			exp:  false},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := IsErrNoRows(tt.err)
			if act != tt.exp {
				t.Fatalf("%s expected %t got %t", tt.name, tt.exp, act)
			}
		})
	}
}

func Test_errors_AsErrNoRows(t *testing.T) {
	t.Parallel()
	table := []struct {
		name string
		e    error
		exp  bool
	}{
		{
			name: "expected true",
			e:    &errNoRows{},
			exp:  true},
		{
			name: "expected false",
			e:    ErrTableNotFound{},
			exp:  false},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, act := AsErrNoRows(tt.e)

			if act != tt.exp {
				t.Fatalf("%s expected %t got %t", tt.name, tt.exp, act)
			}
		})
	}
}
