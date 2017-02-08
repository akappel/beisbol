package main

import "testing"

var findDefinitionGroupsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("ba-n.f.adj. m li"), [][]byte{[]byte("ba-n.f.adj. m li")}},
	{[]byte("b b-n.f. b; b;\n mó\nb-n.m."), [][]byte{[]byte("b b-n.f. b; b;\n mó")}},
	{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)"), [][]byte{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)")}},
}

func TestFindParentDefinitions(t *testing.T) {
	for _, tt := range findDefinitionGroupsTests {
		out, err := findDefinitionGroups(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("findDefinitionGroups(%q) did not find anything!", tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		got := string(out[0])
		expected := string(tt.out[0])
		if got != expected {
			t.Errorf("findDefinitionGroups(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}

var findDefinitionsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("ba-n.f.adj. m li"), [][]byte{[]byte("ba-n.f.adj. m li")}},
	{[]byte("ba-n.f.adj. m li; f"), [][]byte{[]byte("ba-n.f.adj. m li; f")}},
	{[]byte("ba-n.f.adj. m li; f; f b-n.m. am d s"), [][]byte{[]byte("ba-n.f.adj. m li; f"), []byte("f b-n.m am d s")}},
}

func TestFindChildDefinitions(t *testing.T) {
	for _, tt := range findDefinitionsTests {
		out, err := findDefinitions(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		if len(out) < 1 {
			t.Errorf("findDefinitions(%q) did not find anything!", tt.in)
			return
		}

		for _, ttchild := range tt.out {
			expected := string(ttchild)
			found := false

			for _, child := range out {
				got := string(child)
				t.Logf("got: %q", got)

				if got == expected {
					found = true
				}
			}

			if !found {
				t.Errorf("findDefinitions(%q) => %q not found in output set!", string(tt.in), expected)
			}
		}
	}
}
