package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("dict.txt")
	check(err)

	r, err := regexp.Compile("[a-zA-Z -']+-")
	check(err)

	// Replace with one space to make up for removal via the regex
	// out := r.ReplaceAll(dat, []byte(" "))
	// ioutil.WriteFile("removedspaces.txt", out, 0644)

	fmt.Println(r.FindAllString(string(dat), -1))

	// Create new regex for pulling each term and its definition(s)
	// r, err = regexp.Compile("expr")

}

func findDefinitions(in []byte) (out [][]byte, err error) {
	r, err := regexp.Compile("([a-zA-Z\\s]+-*)+-([a-zA-Z].)+")
	if err != nil {
		return
	}

	out = r.FindAll(in, -1)
	return
}
