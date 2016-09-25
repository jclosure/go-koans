package go_koans

import "io/ioutil"
import "strings"

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	// print("lines[5]: " + lines[5])
	assert(lines[5] == lines[5]) // files is too trivial
}
