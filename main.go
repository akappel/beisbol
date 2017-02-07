package main

import "regexp"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// dat, err := ioutil.ReadFile("dict.txt")
	// check(err)

}

func findRootDefinitions(in []byte) (out [][]byte, err error) {
	// r, err := regexp.Compile("([a-zA-Z\\(\\)\\s']+-*)+-([a-zA-Z].)+")
	r, err := regexp.Compile(".+((\\n .+)*|(\\n\\(.+)*)*[^\\n]")
	if err != nil {
		return
	}

	out = r.FindAll(in, -1)
	return
}
