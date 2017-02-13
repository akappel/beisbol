package main

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type translation struct {
	term         string
	translations []string
}

func main() {
	// dat, err := ioutil.ReadFile("dict.txt")
	// check(err)
	// out, err := findTranslationGroups(dat)
	// check(err)

	// for _, el := range out {
	// 	fmt.Println(strings.Replace(string(el), "\n", "", -1))
	// }
}

func findTranslationGroups(in []byte) (out [][]byte, err error) {
	r, err := regexp.Compile(".+((\\n .+)*|(\\n\\(.+)*)*[^\\n]")
	if err != nil {
		return
	}

	out = r.FindAll(in, -1)
	return
}

func findTranslations(in []byte) (out []translation, err error) {
	if len(in) < 1 {
		err = errors.New("Length of input must be greater than 1")
		return
	}

	termRegex, err := regexp.Compile(".+\\.")
	if err != nil {
		return
	}

	// remove all newlines
	in = bytes.Replace(in, []byte("\n"), []byte(""), -1)

	// split on semi-colons
	splits := bytes.Split(in, []byte(";"))

	for _, split := range splits {
		trans := split

		if term := termRegex.Find(split); term != nil {
			// we've found the term in our split, create a new translation struct and add as string
			out = append(out, translation{string(term), []string{}})

			// Remove those term's bytes
			trans = bytes.Replace(trans, term, []byte(""), -1)
		}

		fmt.Println(string(trans))
		// add to last generated translation
		out[len(out)-1].translations = append(out[len(out)-1].translations, strings.Trim(string(trans), " "))

		fmt.Println(out[len(out)-1])
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
