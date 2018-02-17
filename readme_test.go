package mustcheck

import "testing"

func TestCheck(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Check doesn't working somehow.")
		}
	}()
	var err error
	Check(err)
}

func TestMust(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Must doesn't working somehow.")
		}
	}()

	fn := func(val int) (int, error) {
		return val, nil
	}

	if Must(fn(2)).(int) != 2 {
		t.Error("Must doesn't pass through same value")
	}
}
