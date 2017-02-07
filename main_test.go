package main

import "testing"

var findRootDefinitionsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("barehand-n.f.adj. mano limpia"), [][]byte{[]byte("barehand-n.f.adj. mano limpia")}},
	{[]byte("batted ball-n.f. bola bateada; m. batazo;\n machucón\nbatter-n.m."), [][]byte{[]byte("batted ball-n.f. bola bateada; m. batazo;\n machucón")}},
	// {[]byte("All-Star Break-n.m."), [][]byte{[]byte("All-Star Break-n.m.")}},
	// {[]byte("All-Star-n.m."), [][]byte{[]byte("All-Star-n.m.")}},
	// {[]byte("World Series Champions-n.m."), [][]byte{[]byte("World Series Champions-n.m.")}},
	// {[]byte("Commissioner's Trophy-n.m."), [][]byte{[]byte("Commissioner's Trophy-n.m.")}},
}

func TestFindRootDefinitions(t *testing.T) {
	for _, tt := range findRootDefinitionsTests {
		out, err := findRootDefinitions(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("findRootDefinitions(%q) did not find anything!", tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		got := string(out[0])
		expected := string(tt.out[0])
		if got != expected {
			t.Errorf("findRootDefinitions(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}
