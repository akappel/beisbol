package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type entry struct {
	term         string
	translations []string
}

func main() {
	dat, err := ioutil.ReadFile("dict.txt")
	check(err)
	rootEntries, err := findRootEntries(dat)
	check(err)

	for _, rootEntry := range rootEntries {
		entries, err := findEntries(rootEntry)
		check(err)
		for _, entry := range entries {
			fmt.Printf("%s\n", entry.term)
			for _, trans := range entry.translations {
				fmt.Printf("\t%s\n", trans)
			}
		}
	}
}

func findRootEntries(in []byte) (out []string, err error) {
	r, err := regexp.Compile(".+((\\n .+)*|(\\n\\(.+)*)*[^\\n]")
	if err != nil {
		return
	}

	// Covert entirety to string
	str := string(in)

	out = r.FindAllString(str, -1)
	return
}

func findEntries(in string) (out []entry, err error) {
	if len(in) < 1 {
		err = errors.New("Length of input must be greater than 1")
		return
	}

	in = removeNewlines(in)

	// split on semi-colons
	entryParts := strings.Split(in, ";")

	for _, part := range entryParts {
		// Sometimes there will just be a single space as an entry part; we ignore it
		if len(part) < 1 {
			break
		}

		term, err := findTerm(part)
		check(err)

		if len(term) > 0 {
			orig := term
			// Check if the term is just a different version (masc. -> fem. or vice versa)
			if (strings.LastIndex(term, "-")) == -1 {
				temp := term
				// Get the term from the previous entry
				prev := out[len(out)-1].term
				prev = prev[:strings.LastIndex(prev, "-")+1]
				term = prev + "n." + temp
			}

			// we've found the term in our part, create a new entry struct and add as string
			out = append(out, entry{term, []string{}})

			// Remove those term's bytes
			part = strings.Replace(part, orig, "", -1)
		}

		// trim any excess whitespace
		part = strings.Trim(part, " ")

		// add to last generated translation
		out[len(out)-1].translations = append(out[len(out)-1].translations, part)
	}
	return
}

func removeNewlines(in string) (out string) {
	out = strings.Replace(in, "\n", "", -1)
	return
}

func findTerm(in string) (out string, err error) {
	termRegex, err := regexp.Compile(".+\\.")
	if err != nil {
		return
	}

	out = strings.Trim(termRegex.FindString(in), " ")
	return

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
