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

	termRegex, err := regexp.Compile(".+\\-")
	translationRegex, err := regexp.Compile(".+\\.")
	if err != nil {
		return
	}

	// remove all newlines
	in = bytes.Replace(in, []byte("\n"), []byte(""), -1)

	// split on semi-colons
	splits := bytes.Split(in, []byte(";"))

	for _, split := range splits {
		if t := termRegex.Find(split); t != nil {
			// we've found the term in our split, create a new dict struct
			term := string(t)
			term = strings.TrimRight(term, "-")
			out = append(out, translation{string(t), []string{}})
		}

		if d := translationRegex.Find(split); d != nil {
			// translation found, add it
			translation := string(d)
			fmt.Println(translation)
			// add to last generated translation
			out[len(out)-1].translations = append(out[len(out)-1].translations, translation)
		}
	}

	// if just a single translation was found, generate the struct and return
	// if len(splits) == 1 {
	// 	t := termRegex.Find(splits[0])
	// 	trans := definitionRegex.Find(splits[0])
	//
	// 	t = strings.TrimRight(t, "-")
	// 	out = append(out, definition{term: t, translations: []string{}})
	// 	fmt.Println(out[0].term)
	// }

	// for _, split := range splits {
	// 	fmt.Println(string(split))
	// }

	return
	// r, err := regexp.Compile(".+")
	// if err != nil {
	// 	return
	// }
	//
	// out = r.FindAll(in, -1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
