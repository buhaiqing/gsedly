package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var search string
	var replace string
	var filename string
	//var backup bool

	//flag.BoolVar(&backup, "b", "false", "Backup original file with .bak extension")
	flag.StringVar(&search, "s", "", "Search regex or string")
	flag.StringVar(&replace, "r", "", "Replace string")
	flag.StringVar(&filename, "f", "", "Filename to replace")
	flag.Parse()

	if filename == "" {
		fmt.Printf("Filename not specified \n")
		flag.Usage()
		os.Exit(1)
	}

	// File must fit in available ram
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%v \n", err)
		flag.Usage()
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")

	for i, l := range lines {
		//
		if strings.Contains(l, search) {
			rline := strings.Replace(l, search, replace, -1)
			lines[i] = rline
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		fmt.Printf("%v \n", err)
		flag.Usage()
		os.Exit(1)
	}

}
