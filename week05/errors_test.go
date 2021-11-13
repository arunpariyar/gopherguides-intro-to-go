package demo

import (
	"testing"
)

func Test_Errors_All_ErrTableNotFound(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		e    ErrTableNotFound
		tbl  string
		exp  string
	}{
		{
			name: "with table",
			e: ErrTableNotFound{
				table: "test",
			},
			tbl: "test",
			exp: "table not found test",
		},
		{
			name: "without table",
			e:    ErrTableNotFound{},
			tbl:  "",
			exp:  "table not found ",
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			if ok := IsErrTableNotFound(tt.e); ok {

				act := tt.e.Error()
				//checking the error message
				if act != tt.exp {
					t.Errorf("expected %q got %q", tt.exp, act)
				}
				//checking the table returned
				tbl := tt.e.TableNotFound()
				if tbl != tt.tbl {
					t.Errorf("expected %q got %q", tt.tbl, tbl)
				}
			}
		})
	}
}

func Test_Errors_Clauses(t *testing.T) {
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
				t.Fatalf("expected %#v got %#v ", tt.exp.String(), act.String())
			}
		})
	}
}

func Test_Errors_RowNotFound(t *testing.T) {
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

func Test_Errors_errNoRows_Is(t *testing.T) {
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

func Test_Errors_isErrNoRows(t *testing.T) {
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
				t.Fatalf(" expected %t got %t", tt.exp, act)
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
				t.Fatalf("expected %t got %t", tt.exp, act)
			}
		})
	}
}
