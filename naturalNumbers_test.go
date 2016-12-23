package numbers

import "testing"

func TestNaturalNumbers(t *testing.T) {
	for _, element := range []string{"0", "123131", "123123123123123123123123123123123123123123123123123123123123"} {
		if in, err := NatualNumbers.Contains(element); err != nil {
			t.Fatal(err)
		} else if !in {
			t.Errorf("%v not in set")
		}
	}

	if in, err := NatualNumbers.Contains(42); err != nil {
		t.Fatal(err)
	} else if in {
		t.Errorf("Integer in set")
	}

	if in, err := NatualNumbers.Contains("0"); err != nil {
		t.Fatal(err)
	} else if !in {
		t.Errorf("'0' not in set")
	}

	if in, err := NatualNumbers.Contains("00"); err != nil {
		t.Fatal(err)
	} else if !in {
		t.Errorf("'00' in set")
	}
}
