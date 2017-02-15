package main

import (
	"reflect"
	"testing"
)

var findRootEntriesTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("ba-n.f.adj. m li"), [][]byte{[]byte("ba-n.f.adj. m li")}},
	{[]byte("b b-n.f. b; b;\n mó\nb-n.m."), [][]byte{[]byte("b b-n.f. b; b;\n mó")}},
	{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)"), [][]byte{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)")}},
}

func TestFindRootEntries(t *testing.T) {
	for _, tt := range findRootEntriesTests {
		out, err := findRootEntries(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("findRootEntries(%q) did not find anything!", tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		got := string(out[0])
		expected := string(tt.out[0])
		if got != expected {
			t.Errorf("findRootEntries(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}

var findEntriesTests = []struct {
	in  []byte
	out []entry
}{
	{[]byte("b-a-n.f.adj. m li"), []entry{entry{"b-a-n.f.adj.", []string{"m li"}}}},
	{[]byte("ba-n.f.adj. m li; f"), []entry{entry{"ba-n.f.adj.", []string{"m li", "f"}}}},
	// {[]byte("ba-n.f.adj. m li; f; f b-n.m. am d s"), [][]byte{[]byte("ba-n.f.adj. m li; f"), []byte("f b-n.m am d s")}},
}

func TestFindTranslations(t *testing.T) {
	for _, tt := range findEntriesTests {
		entries, err := findEntries(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		if len(entries) < 1 {
			t.Errorf("findEntries(%q) did not find anything!", tt.in)
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
				t.Errorf("findEntries(%q) => want %q", string(tt.in), tt.out)
			}
		}
	}
}

var removeNewlinesTests = []struct {
	in  []byte
	out []byte
}{
	{[]byte("Hello\nWorld"), []byte("HelloWorld")},
	{[]byte("Hello\n World"), []byte("Hello World")},
	{[]byte("Hello\n\tWorld"), []byte("Hello	World")},
}

func TestRemoveNewlines(t *testing.T) {
	for _, tt := range removeNewlinesTests {
		rn := removeNewlines(tt.in)
		if string(rn) != string(tt.out) {
			t.Errorf("removeNewlines(%q) => %q, want %q", tt.in, rn, tt.out)
		}
	}
}

var findTermTests = []struct {
	in  []byte
	out []byte
}{
	{[]byte("hello-n.f. this is a translation"), []byte("hello-n.f.")},
}

func TestFindTerm(t *testing.T) {
	for _, tt := range findTermTests {
		term, _ := findTerm(tt.in)

		if t == nil {
			t.Errorf("findTerm(%q) did not find anything!", tt.in)
			return
		}

		if string(term) != string(tt.out) {
			t.Errorf("removeNewlines(%q) => %q, want %q", tt.in, term, tt.out)
		}
	}
}
