package demo

import (
	"reflect"
	"testing"
)

func Test_db(t *testing.T) {
	t.Parallel()

	s := &Store{}

	table := []struct {
		name  string
		store *Store
		exp   data
	}{
		{name: "empty store", store: s, exp: s.data},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.store.db()

			if !reflect.DeepEqual(s.data, data) {
				t.Fatalf("%s expected %q got %q", tt.name, tt.exp, data)
			}
		})
	}
}

func Test_All(t *testing.T) {
	t.Parallel()
	tn1 := "users"
	tn2 := "orders"

	s1 := &Store{}
	db1 := s1.db()
	err1 := ErrTableNotFound{
		table: "users",
	}

	s2 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}
	db2 := s2.db()

	table := []struct {
		name  string
		store *Store
		tn    string
		mods  Models
		exp   error
	}{
		{name: "empty store", store: s1, tn: tn1, mods: db1[tn1], exp: err1},
		{name: "store with orders", store: s2, tn: tn2, mods: db2[tn2], exp: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			mods, err := tt.store.All(tt.tn)

			if err != nil {
				if ok := IsErrTableNotFound(err); !ok {
					t.Fatalf("%s expected error type %s got %T", tt.name, "ErrTableNotFound", err)
				}

				if !reflect.DeepEqual(err.Error(), tt.exp.Error()) {
					t.Fatalf("%s expected %q got %q", tt.name, tt.exp.Error(), err.Error())
				}
			}

			if !reflect.DeepEqual(tt.mods, mods) {
				t.Fatalf("%s expected %q got %q", tt.name, tt.mods, mods)
			}
		})
	}
}

func Test_Len(t *testing.T) {
	t.Parallel()
	tn1 := "users"
	tn2 := "orders"

	s1 := &Store{}

	s2 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}

	err1 := ErrTableNotFound{
		table: "users",
	}

	table := []struct {
		name  string
		store *Store
		tn    string
		len   int
		exp   error
	}{
		{name: "empty store", store: s1, tn: tn1, len: 0, exp: err1},
		{name: "store with orders", store: s2, tn: tn2, len: 2, exp: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			len, err := tt.store.Len(tt.tn)

			// need to think if this block should be inside here
			if err != nil {
				if ok := IsErrTableNotFound(err); !ok {
					t.Fatalf("%s expected error type %s got %T", tt.name, "ErrTableNotFound", err)
				}

				if !reflect.DeepEqual(err.Error(), tt.exp.Error()) {
					t.Fatalf("%s expected %q got %q", tt.name, tt.exp.Error(), err.Error())
				}
			}
			if len != tt.len {
				t.Fatalf("%s expected length %q got %q", tt.name, tt.len, len)
			}
		})
	}
}

func Test_Insert(t *testing.T) {

	tn := "orders"

	s1 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}

	m1 := Model{"id": 3, "item": "mobile"}

	s2 := &Store{
		data: data{
			"orders": {{"id": 1, "item": "desktop"}, {"id": 2, "item": "laptop"}},
		},
	}
	m2 := Model{"id": 3, "item": "mobile"}
	m3 := Model{"id": 4, "item": "smartwatch"}

	addOne := []Model{m1}
	addTwo := []Model{m2, m3}

	expMods1 := []Model{
		{"id": 1, "item": "desktop"},
		{"id": 2, "item": "laptop"},
		{"id": 3, "item": "mobile"}}

	expMods2 := []Model{
		{"id": 1, "item": "desktop"},
		{"id": 2, "item": "laptop"},
		{"id": 3, "item": "mobile"},
		{"id": 4, "item": "smartwatch"}}

	table := []struct {
		name  string
		store *Store
		tn    string
		mods  []Model
		exp   []Model
	}{
		{name: "insert one item", store: s1, tn: tn, mods: addOne, exp: expMods1},
		{name: "insert two item", store: s2, tn: tn, mods: addTwo, exp: expMods2},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			tt.store.Insert(tt.tn, tt.mods...)
			db := tt.store.db()
			mods := db[tt.tn]
			if reflect.DeepEqual(tt.exp, mods) {
				t.Fatalf("%s expected models %v got %v", tt.name, tt.exp, mods)
			}

		})
	}
}

func Test_Select(t *testing.T) {
	tn := "orders"

	// store with table and model
	s1 := &Store{
		data: data{
			"orders": {
				{"desktop": "dell"},
				{"mobile": "apple"},
				{"smartwatch": "xioami"},
				{"gps": "garmin"},
			},
		},
	}
	c1 := Clauses{"gps": "garmin"}
	len1 := len(c1)
	m1 := Model{"gps": "garmin"}

	//store with table but no model
	s2 := &Store{
		data: data{
			"orders": {},
		},
	}
	c2 := Clauses{"gps": "garmin"}
	len2 := len(c2)
	m2 := Model{}
	err2 := &errNoRows{
		clauses: c2,
		table:   tn,
	}

	//query with no clause
	s3 := &Store{
		data: data{
			"orders": {
				{"smartwatch": "xioami"},
			},
		},
	}
	c3 := Clauses{}
	len3 := len(c3)
	m3 := Model{"smartwatch": "xioami"}
	err3 := &errNoRows{
		clauses: c3,
		table:   tn,
	}

	//query with no table name
	s4 := &Store{
		data: data{
			"orders": {
				{"smartwatch": "xioami"},
			},
		},
	}
	c4 := Clauses{}
	len4 := len(c3)
	m4 := Model{"smartwatch": "xioami"}
	err4 := ErrTableNotFound{
		table: "",
	}
	//table test
	table := []struct {
		name string
		str  *Store
		tn   string
		cls  Clauses
		mds  Model
		len  int
		exp  error
	}{
		{name: "store with table and model", str: s1, tn: tn, cls: c1, mds: m1, len: len1, exp: nil},
		{name: "store with table but no model", str: s2, tn: tn, cls: c2, mds: m2, len: len2, exp: err2},
		{name: "query with no clause", str: s3, tn: tn, cls: c3, mds: m3, len: len3, exp: err3},
		{name: "query with no table name", str: s4, tn: "", cls: c4, mds: m4, len: len4, exp: err4},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			mods, err := tt.str.Select(tt.tn, tt.cls)
			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.exp.Error()) {
					t.Fatalf("%s expected %q got %q", tt.name, tt.exp.Error(), err.Error())
				}

				for _, mod := range mods {
					if !reflect.DeepEqual(mod, tt.mds) {
						t.Fatalf("%s expected %q got %q", tt.name, tt.mds, mod)
					}
				}
			}

			if len(tt.cls) != tt.len {
				t.Fatalf("length of clause %d need at least 1", len(tt.cls))
			}

			if len(mods) == 0 {
				if !reflect.DeepEqual(err.Error(), tt.exp.Error()) {
					t.Fatal("no models returned")
				}
			}
		})
	}
}
