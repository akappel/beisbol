package main

import (
	"reflect"
	"testing"
)

var findRootEntriesTests = []struct {
	in  []byte
	out []string
}{
	{[]byte("ba-n.f.adj. m li"), []string{"ba-n.f.adj. m li"}},
	{[]byte("b b-n.f. b; b;\n mó\nb-n.m."), []string{"b b-n.f. b; b;\n mó"}},
	{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)"), []string{"b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)"}},
}

var findEntriesTests = []struct {
	in  string
	out []entry
}{
	{"b-a-n.f.adj. m li", []entry{entry{"b-a-n.f.adj.", []string{"m li"}}}},
	{"ba-n.f.adj. m li; f", []entry{entry{"ba-n.f.adj.", []string{"m li", "f"}}}},
	// {[]byte("ba-n.f.adj. m li; f; f b-n.m. am d s"), [][]byte{[]byte("ba-n.f.adj. m li; f"), []byte("f b-n.m am d s")}},
}

var removeNewlinesTests = []struct {
	in  string
	out string
}{
	{"Hello\nWorld", "HelloWorld"},
	{"Hello\n World", "Hello World"},
	{"Hello\n\tWorld", "Hello	World"},
}

var findTermTests = []struct {
	in  string
	out string
}{
	{"hello-n.f. this is a translation", "hello-n.f."},
}

func TestFindRootEntries(t *testing.T) {
	name := "findRootEntries"

	for _, tt := range findRootEntriesTests {
		out, err := findRootEntries(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("%s(%q) did not find anything!", name, tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		if out[0] != tt.out[0] {
			t.Errorf("%s(%q) => %q, want %q", name, tt.in, out[0], tt.out[0])
		}

	}
}

func TestFindEntries(t *testing.T) {
	name := "findEntries"

	for _, tt := range findEntriesTests {
		entries, err := findEntries(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		if len(entries) < 1 {
			t.Errorf("%s(%q) did not find anything!", name, tt.in)
			return
		}

		for _, expectedEntry := range tt.out {
			found := false

			for _, actualEntry := range entries {
				if actualEntry.term == expectedEntry.term && reflect.DeepEqual(actualEntry.translations, expectedEntry.translations) {
					found = true
				}
			}

			if !found {
				t.Errorf("%s(%q) => want %q", name, string(tt.in), tt.out)
			}
		}
	}
}

func TestRemoveNewlines(t *testing.T) {
	name := "removeNewlines"

	for _, tt := range removeNewlinesTests {
		rn := removeNewlines(tt.in)
		if rn != tt.out {
			t.Errorf("%s(%q) => %q, want %q", name, tt.in, rn, tt.out)
		}
	}
}

func TestFindTerm(t *testing.T) {
	name := "findTerm"

	for _, tt := range findTermTests {
		term, _ := findTerm(tt.in)

		if t == nil {
			t.Errorf("%s(%q) did not find anything!", name, tt.in)
			return
		}

		if term != tt.out {
			t.Errorf("%s(%q) => %q, want %q", name, tt.in, term, tt.out)
		}
	}
}
