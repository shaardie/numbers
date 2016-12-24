package numbers

import "testing"

func TestIntegers(t *testing.T) {
	for _, element := range []string{
		"0",
		"123131",
		"123123123123123123123123123123123123123123123123123123123123",
		"-13031",
		"-1000"} {
		if in, err := Integers.Contains(element); err != nil {
			t.Fatal(err)
		} else if !in {
			t.Errorf("%v not in set")
		}
	}

	if in, err := Integers.Contains(42); err != nil {
		t.Fatal(err)
	} else if in {
		t.Errorf("Type Integer in set")
	}

	if in, err := Integers.Contains("-0"); err != nil {
		t.Fatal(err)
	} else if !in {
		t.Errorf("'-0' in set")
	}

	if in, err := Integers.Contains("00"); err != nil {
		t.Fatal(err)
	} else if !in {
		t.Errorf("'00' in set")
	}
}
