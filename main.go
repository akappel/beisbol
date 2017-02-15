package main

import (
	"bytes"
	"errors"
	"regexp"
	"strings"
)

type entry struct {
	term         string
	translations []string
}

func main() {
	// dat, err := ioutil.ReadFile("dict.txt")
	// check(err)
	// groups, err := findTranslationGroups(dat)
	// check(err)
	//
	// for _, group := range groups {
	// 	translations, err := findTranslations(group)
	// 	check(err)
	// 	for _, t := range translations {
	// 		fmt.Printf("%s\n", t.term)
	// 		for _, trans := range t.translations {
	// 			fmt.Printf("\t%s\n", trans)
	// 		}
	// 	}
	// }
}

func findRootEntries(in []byte) (out [][]byte, err error) {
	r, err := regexp.Compile(".+((\\n .+)*|(\\n\\(.+)*)*[^\\n]")
	if err != nil {
		return
	}

	out = r.FindAll(in, -1)
	return
}

func findEntries(in []byte) (out []entry, err error) {
	if len(in) < 1 {
		err = errors.New("Length of input must be greater than 1")
		return
	}

	in = removeNewlines(in)

	// split on semi-colons
	splits := bytes.Split(in, []byte(";"))

	for _, split := range splits {
		trans := split

		if term, _ := findTerm(split); term != nil {
			// we've found the term in our split, create a new translation struct and add as string
			out = append(out, entry{string(term), []string{}})

			// Remove those term's bytes
			trans = bytes.Replace(trans, term, []byte(""), -1)
		}

		// add to last generated translation
		out[len(out)-1].translations = append(out[len(out)-1].translations, strings.Trim(string(trans), " "))
	}
	return
}

func removeNewlines(in []byte) (out []byte) {
	out = bytes.Replace(in, []byte("\n"), []byte(""), -1)
	return
}

func findTerm(in []byte) (out []byte, err error) {
	termRegex, err := regexp.Compile(".+\\.")
	if err != nil {
		return
	}

	out = termRegex.Find(in)
	return

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
