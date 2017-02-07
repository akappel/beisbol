package main

import "testing"

var findDefinitionsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("baseball-n.m."), [][]byte{[]byte("baseball-n.m.")}},
	{[]byte("All-Star Break-n.m."), [][]byte{[]byte("All-Star Break-n.m.")}},
	{[]byte("All-Star-n.m."), [][]byte{[]byte("All-Star-n.m.")}},
	{[]byte("World Series Champions-n.m."), [][]byte{[]byte("World Series Champions-n.m.")}},
}

func TestFindDefinitions(t *testing.T) {
	for _, tt := range findDefinitionsTests {
		out, err := findDefinitions(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		got := string(out[0])
		expected := string(tt.out[0])

		if got != expected {
			t.Errorf("findDefinitions(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}
