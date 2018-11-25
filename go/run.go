package main

import "os"
import "fmt"
import "strconv"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 4 {
		fmt.Println("Not Enough args")
		os.Exit(0)
	}
	args := os.Args[2:]
	numberOfArgs := len(args)
	var searchInts []int

	searchInts = make([]int, numberOfArgs)

	var currentInt, findValue, k int = 0, 0, 0
	var stringCurrent, KString, findValueString string

	var found bool

	KString = os.Args[1]
	fmt.Println(KString)
	var err error

	k, err = strconv.Atoi(KString)
	if err != nil {
		os.Exit(1)
	}

	found = false

	for i := 0; i < numberOfArgs; i++ {
		//fmt.Println(i)
		stringCurrent = args[i]

		currentInt, err = strconv.Atoi(stringCurrent)

		if currentInt >= k {
			continue
		}
		findValue = k - currentInt
		searchInts[i] = findValue
		findValueString = strconv.Itoa(findValue)

		if contains(searchInts, currentInt) || containsString(args, findValueString) {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("Found two integers %d and %d ", currentInt, findValue)
		fmt.Printf("add to %s", KString)
	} else {
		fmt.Println("No two integers sum")
	}
}
