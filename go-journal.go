package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// this is temp while deving
// This should be stored in the program files equivalent
var STOREPATH string = "./dev"

var HEADERSTRING string = fmt.Sprintf(`
=====================================================
Welcome to your Journal. Today's date is <Date here>.
Lay that shit out i guess:
=====================================================
`)

func RenderHeader() {

}

func ReceiveEntry() bool {

	entname := "test-entry"

	// scanln might be better for continuous updates
	input := bufio.NewReader(os.Stdin)

	entbytes := []byte("==================\n")
	// figure out better ways to render this
	// animations and shit perhaps
	fmt.Print(HEADERSTRING)
	for {
		fmt.Print(">>> ")
		line, err := input.ReadString('\n')
		check(err)

		linebite := []byte(line)
		entbytes = append(entbytes, linebite...)
		if len(linebite) == 1 && linebite[0] == 10 {
			// EOF
			break
		}
	}

	fmt.Println(entbytes)
	// add new line to start if file already exists

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

func FlagParser(f []string) {
	RenderHeader()
	ReceiveEntry()

}

func main() {
	// TODO import using os.Getenv() for consistency
	// GOBIN := "/usr/bin"
	// STORE_FILES := "/var/lib/go-journal"
	// fmt.Println(os.Environ())

	// fmt.Println(reflect.TypeOf(os.Args))

	FlagParser(os.Args[1:])

}
