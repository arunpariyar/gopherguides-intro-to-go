package demo

import (
	"testing"
)

func Test_clause_String(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		cls  Clauses
		exp  string
	}{
		{
			name: "matching clauses",
			cls:  Clauses{"name": "Sir Canon Doyle", "job": "Investigator"},
			exp:  `"job" = "Investigator" and "name" = "Sir Canon Doyle"`},
		{
			name: "unmatching clauses",
			cls:  Clauses{"spiderman": "peter parker", "batman": "bruce wayne"},
			exp:  `"batman" = "bruce wayne" and "spiderman" = "peter parker"`},
		{
			name: "empty clauses",
			cls:  Clauses{},
			exp:  ``},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.cls.String()

			if act != tt.exp {
				t.Fatalf("expected %q, got %q", tt.exp, act)
			}
		})
	}
}

func Test_clause_Match(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		cls  Clauses
		mdl  Model
		exp  bool
	}{
		{
			name: "true case",
			cls:  Clauses{"name": "Sir Canon Doyle", "job": "Investigator"},
			mdl:  Model{"name": "Sir Canon Doyle", "job": "Investigator"},
			exp:  true},
		{
			name: "false case",
			cls:  Clauses{"spiderman": "peter parker", "batman": "bruce wayne"},
			mdl:  Model{"spiderman": "eddy green", "batman": "damien wayne"},
			exp:  false},
		{
			name: "empty case",
			cls:  Clauses{},
			mdl:  Model{},
			exp:  true,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.cls.Match(tt.mdl)

			if res != tt.exp {
				t.Fatalf("%s expected %t got %t", tt.name, tt.exp, res)
			}
		})
	}
}
