package main

import "testing"

var findDefinitionsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("barehand-n.f.adj."), [][]byte{[]byte("barehand-n.f.adj.")}},
	{[]byte("baseball (sport)-n.m."), [][]byte{[]byte("baseball (sport)-n.m.")}},
	{[]byte("All-Star Break-n.m."), [][]byte{[]byte("All-Star Break-n.m.")}},
	{[]byte("All-Star-n.m."), [][]byte{[]byte("All-Star-n.m.")}},
	{[]byte("World Series Champions-n.m."), [][]byte{[]byte("World Series Champions-n.m.")}},
	{[]byte("Commissioner's Trophy-n.m."), [][]byte{[]byte("Commissioner's Trophy-n.m.")}},
}

func TestFindDefinitions(t *testing.T) {
	for _, tt := range findDefinitionsTests {
		out, err := findDefinitions(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("findDefinitions(%q) did not find anything!", tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		got := string(out[0])
		expected := string(tt.out[0])
		if got != expected {
			t.Errorf("findDefinitions(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}
