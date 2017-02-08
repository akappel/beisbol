package main

import (
	"bytes"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// dat, err := ioutil.ReadFile("dict.txt")
	// check(err)
	// out, err := findDefinitionGroups(dat)
	// check(err)

	// for _, el := range out {
	// 	fmt.Println(strings.Replace(string(el), "\n", "", -1))
	// }
}

func findDefinitionGroups(in []byte) (out [][]byte, err error) {
	r, err := regexp.Compile(".+((\\n .+)*|(\\n\\(.+)*)*[^\\n]")
	if err != nil {
		return
	}

	out = r.FindAll(in, -1)
	return
}

func findDefinitions(in []byte) (out [][]byte, err error) {
	// remove all newlines
	in = bytes.Replace(in, []byte("\n"), []byte(""), -1)
	r, err := regexp.Compile(".+")
	if err != nil {
		return
	}

	out = r.FindAll(in, -1)
	return
}
