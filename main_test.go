package main

import (
	"reflect"
	"testing"
)

var findTranslationGroupsTests = []struct {
	in  []byte
	out [][]byte
}{
	{[]byte("ba-n.f.adj. m li"), [][]byte{[]byte("ba-n.f.adj. m li")}},
	{[]byte("b b-n.f. b; b;\n mó\nb-n.m."), [][]byte{[]byte("b b-n.f. b; b;\n mó")}},
	{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)"), [][]byte{[]byte("b (o)-n.f. p; b; p\n(p) m. e (s); m. c\n d c (h)")}},
}

func TestFindTranslationGroups(t *testing.T) {
	for _, tt := range findTranslationGroupsTests {
		out, err := findTranslationGroups(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		// if the regex just completely failed to find anything, fail out
		if len(out) < 1 {
			t.Errorf("findTranslationGroups(%q) did not find anything!", tt.in)
			return
		}

		// else, check that what it did find matches what we expect
		got := string(out[0])
		expected := string(tt.out[0])
		if got != expected {
			t.Errorf("findTranslationGroups(%q) => %q, want %q", string(tt.in), got, expected)
		}

	}
}

var findTranslationsTests = []struct {
	in  []byte
	out []translation
}{
	{[]byte("b-a-n.f.adj. m li"), []translation{translation{"b-a-n.f.adj.", []string{"m li"}}}},
	{[]byte("ba-n.f.adj. m li; f"), []translation{translation{"ba-n.f.adj.", []string{"m li", "f"}}}},
	// {[]byte("ba-n.f.adj. m li; f; f b-n.m. am d s"), [][]byte{[]byte("ba-n.f.adj. m li; f"), []byte("f b-n.m am d s")}},
}

func TestFindTranslations(t *testing.T) {
	for _, tt := range findTranslationsTests {
		translations, err := findTranslations(tt.in)
		if err != nil {
			t.Fatal(err)
		}

		if len(translations) < 1 {
			t.Errorf("findTranslations(%q) did not find anything!", tt.in)
			return
		}

		for _, expectedTranslation := range tt.out {
			found := false

			for _, actualTranslation := range translations {
				if actualTranslation.term == expectedTranslation.term && reflect.DeepEqual(actualTranslation.translations, expectedTranslation.translations) {
					found = true
				}
			}

			if !found {
				t.Errorf("findTranslations(%q) => want %q", string(tt.in), tt.out)
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
