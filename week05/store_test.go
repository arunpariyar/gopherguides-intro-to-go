package demo

import (
	"fmt"
	"testing"
)

func Test_Store_db(t *testing.T) {
	t.Parallel()
	s := &Store{
		data: nil,
	}

	db := s.db()

	if len(db) != 0 {
		t.Fatalf("expected 0 got %v", len(db))
	}
}

func Test_Store_All(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  string
		store *Store
		tn    string
		mods  Models
	}{
		{name: "store with orders",
			store: &Store{
				data: data{
					"orders": {
						{"desktop": "dell"},
						{"mobile": "apple"},
						{"smartwatch": "xioami"},
						{"gps": "garmin"},
					},
				},
			},
			tn: "orders",
			mods: Models{
				{"desktop": "dell"},
				{"mobile": "apple"},
				{"smartwatch": "xioami"},
				{"gps": "garmin"},
			},
		},
		{name: "empty store",
			store: &Store{
				data: data{
					"orders": {},
				},
			},
			tn:   "orders",
			mods: Models{},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			mods, err := tt.store.All(tt.tn)

			if err != nil {
				if ok := IsErrTableNotFound(err); !ok {
					t.Fatalf("expected error type %s got %T", "ErrTableNotFound", err)
				}
			}

			assertModels(t, mods, tt.mods)
		})
	}
}

func Test_Store_Len(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  string
		store *Store
		tn    string
		len   int
		exp   error
	}{
		{
			name:  "empty store",
			store: &Store{},
			tn:    "users",
			len:   0,
		},
		{name: "store with orders",
			store: &Store{
				data: data{
					"orders": {
						{"id": 1, "item": "desktop"},
						{"id": 2, "item": "laptop"},
					},
				},
			},
			tn:  "orders",
			len: 2,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			len, err := tt.store.Len(tt.tn)
			if err != nil {
				if ok := IsErrTableNotFound(err); !ok {
					t.Fatalf(" expected error type %s got %T", "ErrTableNotFound", err)
				}
			}
			if len != tt.len {
				t.Fatalf("expected length %q got %q", tt.len, len)
			}
		})
	}
}

func Test_Store_Insert(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  string
		store *Store
		tn    string
		mlist []Model
		mods  []Model
	}{
		{
			name: "insert one item",
			store: &Store{
				data: data{
					"orders": {
						{"desktop": "dell"},
						{"mobile": "apple"},
					},
				},
			},
			tn:    "orders",
			mlist: []Model{{"smartwatch": "xioami"}},
			mods: []Model{
				{"desktop": "dell"},
				{"mobile": "apple"},
				{"smartwatch": "xioami"},
			}},
		{
			name: "insert two item",
			store: &Store{
				data: data{
					"orders": {
						{"desktop": "dell"},
						{"mobile": "apple"},
					},
				},
			},
			tn:    "orders",
			mlist: []Model{{"smartwatch": "xioami"}, {"gps": "garmin"}},
			mods: []Model{
				{"desktop": "dell"},
				{"mobile": "apple"},
				{"smartwatch": "xioami"},
				{"gps": "garmin"},
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tt.store.Insert(tt.tn, tt.mlist...)
			db := tt.store.db()
			mods := db[tt.tn]
			assertModels(t, mods, tt.mods)

		})
	}
}

func Test_Store_Select(t *testing.T) {
	t.Parallel()
	table := []struct {
		name string
		str  *Store
		tn   string
		cls  Clauses
		mods Models
		exp  error
	}{
		{
			name: "store with table and model",
			str: &Store{
				data: data{
					"orders": {
						{"desktop": "dell"},
						{"mobile": "apple"},
						{"smartwatch": "xioami"},
						{"gps": "garmin"},
					},
				},
			},
			tn:   "orders",
			cls:  Clauses{"gps": "garmin"},
			mods: []Model{{"gps": "garmin"}},
			exp:  nil},
		{
			name: "store with table but no model",
			str: &Store{
				data: data{
					"orders": {},
				},
			},
			tn:   "orders",
			cls:  Clauses{"gps": "garmin"},
			mods: nil,
			exp: &errNoRows{
				clauses: Clauses{"gps": "garmin"},
				table:   "orders",
			},
		},
		{
			name: "query with no clause",
			str: &Store{
				data: data{
					"orders": {
						{"smartwatch": "xioami"},
					},
				},
			},
			tn:   "orders",
			cls:  Clauses{},
			mods: []Model{{"smartwatch": "xioami"}},
			exp: &errNoRows{
				clauses: Clauses{},
				table:   "orders",
			}},
		{
			name: "query with no table name",
			str: &Store{
				data: data{
					"orders": {
						{"smartwatch": "xioami"},
					},
				},
			},
			tn:   "",
			cls:  Clauses{},
			mods: nil,
			exp: ErrTableNotFound{
				table: "",
			}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			mods, err := tt.str.Select(tt.tn, tt.cls)
			if err != nil {
				if err.Error() != tt.exp.Error() {
					t.Fatalf("expected %q got %q", tt.exp.Error(), err.Error())
				}
			}

			assertModels(t, mods, tt.mods)

		})
	}
}

func assertModels(t testing.TB, exp Models, act Models) {
	t.Helper()

	if len(exp) != len(act) {
		t.Fatalf("expected %d models, got %d", len(exp), len(act))
	}

	exps := fmt.Sprintf("%#v", exp)
	acts := fmt.Sprintf("%#v", act)

	if acts != exps {
		t.Fatalf("expected %s, got %s", exps, acts)
	}
}
