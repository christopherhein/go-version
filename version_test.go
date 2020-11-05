package version

import (
	"testing"
)

type testcase struct {
	expected  string
	shortened bool
}

func TestToJSON(t *testing.T) {
	v := New("dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019")
	json := v.ToJSON()

	expected := "{\"Version\":\"dev\",\"Commit\":\"fee8dd1f7c24da23508e26347694be0acce5631b\",\"Date\":\"Fri Jun 21 10:50:15 PDT 2019\"}\n"
	if json != expected {
		t.Errorf("Expected json %s to equal %s", json, expected)
	}
}

func TestToYAML(t *testing.T) {
	v := New("dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019")
	yaml := v.ToYAML()

	expected := `Commit: fee8dd1f7c24da23508e26347694be0acce5631b
Date: Fri Jun 21 10:50:15 PDT 2019
Version: dev
`

	if yaml != expected {
		t.Errorf("Expected yaml %s to equal %s", yaml, expected)
	}
}

func TestToShortened(t *testing.T) {
	v := New("dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019")
	short := v.ToShortened()

	expected := `Version: dev
`

	if short != expected {
		t.Errorf("Expected shortened %s to equal %s", short, expected)
	}
}

func TestFunc(t *testing.T) {
	cases := []testcase{
		{
			expected:  "{\"Version\":\"dev\",\"Commit\":\"fee8dd1f7c24da23508e26347694be0acce5631b\",\"Date\":\"Fri Jun 21 10:50:15 PDT 2019\"}\n",
			shortened: false,
		},
		{
			expected: `Version: dev
`,
			shortened: true,
		},
	}

	for _, c := range cases {
		resp := FuncWithOutput(c.shortened, "dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019", JSON)
		if resp != c.expected {
			t.Errorf("got %q want %q", resp, c.expected)
		}
	}
}
