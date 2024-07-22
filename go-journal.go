package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// this is temp while deving
// This should be stored in the program files equivalent
var STOREPATH string = "./dev"

func ReceiveEntry() bool {

	var entry string

	entname := "test-entry"

	// scanln might be better for continuous updates
	fmt.Scan(&entry)

	// add new line to start if file already exists
	entbytes := []byte(entry)

	err := os.WriteFile(fmt.Sprintf("%s/%s.txt", STOREPATH, entname), entbytes, 0644)
	check(err)

	return true
}

func ReadEntry(trange ...string) error {

	if len(trange) == 0 {
		// get all or the default range
	} else if len(trange) == 2 {
		// get between reange
	} else {
		return errors.New("Invalid number of range arguments")
	}

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func flagParser(f []string) {
	// entry point parser returning funcs to be called
	fmt.Println(f)
	ReceiveEntry()

}

func main() {
	// TODO import using os.Getenv() for consistency
	// GOBIN := "/usr/bin"
	// STORE_FILES := "/var/lib/go-journal"
	// fmt.Println(os.Environ())
	fmt.Println(reflect.TypeOf(os.Args))

	flagParser(os.Args[1:])

}
