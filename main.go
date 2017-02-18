package main

import (
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
	splits := strings.Split(in, ";")

	for _, split := range splits {
		trans := split

		if term, _ := findTerm(split); term != "" {
			// we've found the term in our split, create a new translation struct and add as string
			out = append(out, entry{string(term), []string{}})

			// Remove those term's bytes
			trans = strings.Replace(trans, term, "", -1)
		}

		// add to last generated translation
		out[len(out)-1].translations = append(out[len(out)-1].translations, strings.Trim(trans, " "))
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

	out = termRegex.FindString(in)
	return

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
