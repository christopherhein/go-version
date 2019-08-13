package version

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestToJSON(t *testing.T) {
	v := New("dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019")
	json := v.ToJSON()

	expected := "{\"Version\":\"dev\",\"Commit\":\"fee8dd1f7c24da23508e26347694be0acce5631b\",\"Date\":\"Fri Jun 21 10:50:15 PDT 2019\"}\n"
	if json != expected {
		t.Errorf("Expected json %s to equal %s", json, expected)
	}
}

func TestToShortened(t *testing.T) {
	v := New("dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019")
	short := v.ToShortened()

	expected := `Version: dev
Commit: fee8dd1f7c24da23508e26347694be0acce5631b
Date: Fri Jun 21 10:50:15 PDT 2019
`

	if short != expected {
		t.Errorf("Expected shortened %s to equal %s", short, expected)
	}
}

type testcase struct {
	expected  string
	shortened bool
}

func TestFunc(t *testing.T) {
	cases := []testcase{
		testcase{
			expected:  "{\"Version\":\"dev\",\"Commit\":\"fee8dd1f7c24da23508e26347694be0acce5631b\",\"Date\":\"Fri Jun 21 10:50:15 PDT 2019\"}\n",
			shortened: false,
		},
		testcase{
			expected: `Version: dev
Commit: fee8dd1f7c24da23508e26347694be0acce5631b
Date: Fri Jun 21 10:50:15 PDT 2019
`,
			shortened: true,
		},
	}

	for _, c := range cases {
		buffer := &bytes.Buffer{}
		Func(buffer, c.shortened, "dev", "fee8dd1f7c24da23508e26347694be0acce5631b", "Fri Jun 21 10:50:15 PDT 2019")(&cobra.Command{}, []string{})
		resp := buffer.String()
		if resp != c.expected {
			t.Errorf("got %q want %q", resp, c.expected)
		}
	}
}
