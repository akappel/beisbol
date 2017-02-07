package main

import "testing"

var findParentDefinitionsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("ba-n.f.adj. m li"), [][]byte{[]byte("ba-n.f.adj. m li")}},
	{[]byte("b b-n.f. b; b;\n mó\nb-n.m."), [][]byte{[]byte("b b-n.f. b; b;\n mó")}},
	{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)"), [][]byte{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)")}},
}

func TestFindParentDefinitions(t *testing.T) {
	for _, tt := range findParentDefinitionsTests {
		out, err := findParentDefinitions(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("findParentDefinitions(%q) did not find anything!", tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		got := string(out[0])
		expected := string(tt.out[0])
		if got != expected {
			t.Errorf("findParentDefinitions(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}

func TestFindChildDefinitions(t *testing.T) {
	t.Skip("Implement me")
}
