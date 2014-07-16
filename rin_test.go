package ringo

import "testing"

func TestInsertAndDump(t *testing.T) {
	r := New(10)
	for i := 0; i < 15; i++ {
		p := r.Insert(i)

		t.Logf("Dump: %+v Out: %+v", r.Dump(), p)

		if i < 10 {
			if p != nil {
				t.Errorf("Invalid returned value: %+v", p)
			}
		} else {
			if p.(int) != i-10 {
				t.Errorf("Invalid returned value: %+v", p)
			}
		}
	}

	d := r.Dump()
	t.Logf("Dump: %+v", d)
	for j, v := range d {
		if v == nil {
			t.Errorf("Invalid returned value: %+v", v)
		}

		vi, ok := v.(int)
		if !ok {
			t.Errorf("Invalid returned value: %+v", v)
		}

		if vi != j+5 {
			t.Errorf("Invalid returned value: %+v", vi)
		}
	}
}

func TestShiftLeft(t *testing.T) {
	r := New(5)

	for i := 0; i < 5; i++ {
		r.Insert(i)
	}

	d1 := r.Dump()
	v1 := d1[0].(int)
	if v1 != 0 {
		t.Errorf("Invalid returned value: %+v", v1)
	}

	r.ShiftLeft()

	d2 := r.Dump()
	v2 := d2[0].(int)
	if v2 != 1 {
		t.Errorf("Invalid returned value: %+v", v2)
	}

	r.ShiftLeft()
	r.ShiftLeft()
	r.ShiftLeft()

	d3 := r.Dump()
	v3 := d3[0].(int)
	if v3 != 4 {
		t.Errorf("Invalid returned value: %+v", v3)
	}

	r.ShiftLeft()
	r.ShiftLeft()
	r.ShiftLeft()

	d4 := r.Dump()
	v4 := d4[0].(int)
	if v4 != 2 {
		t.Errorf("Invalid returned value: %+v", v4)
	}
}
