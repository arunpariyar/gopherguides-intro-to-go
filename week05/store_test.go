package demo

import (
	"fmt"
	"sort"
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

func Test_store_All(t *testing.T) {
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
					t.Fatalf("%s expected error type %s got %T", tt.name, "ErrTableNotFound", err)
				}
			}

			assertSameModels(t, mods, tt.mods)
		})
	}
}

func Test_store_Len(t *testing.T) {
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
					t.Fatalf("%s expected error type %s got %T", tt.name, "ErrTableNotFound", err)
				}
			}
			if len != tt.len {
				t.Fatalf("%s expected length %q got %q", tt.name, tt.len, len)
			}
		})
	}
}

func Test_store_Insert(t *testing.T) {
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
			assertSameModels(t, mods, tt.mods)

		})
	}
}

func Test_store_Select(t *testing.T) {
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
			mods: []Model{},
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
			mods: []Model{{"smartwatch": "xioami"}},
			exp: ErrTableNotFound{
				table: "",
			}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			mods, err := tt.str.Select(tt.tn, tt.cls)
			if err != nil {
				if err.Error() != tt.exp.Error() {
					t.Fatalf("%s expected %q got %q", tt.name, tt.exp.Error(), err.Error())
				}
			}

			assertSameModels(t, mods, tt.mods)

		})
	}
}

func assertSameModels(t testing.TB, act Models, exp Models) {
	t.Helper()

	k1 := []string{}
	v1 := []string{}

	k2 := []string{}
	v2 := []string{}

	for _, m := range act {
		for k, v := range m {
			k1 = append(k1, k)
			vc1 := fmt.Sprintf("%s", v)
			v1 = append(v1, vc1)
		}
	}

	for _, m := range exp {
		for k, v := range m {
			k2 = append(k2, k)
			vc2 := fmt.Sprintf("%s", v)
			v2 = append(v2, vc2)
		}
	}

	sort.Strings(k1)
	sort.Strings(k2)
	sort.Strings(v1)
	sort.Strings(v2)

	for i := 0; i < len(k1); i++ {
		if k1[i] != k2[i] || v1[i] != v2[i] {
			t.Fatal("models don't match")
		}
	}
}
