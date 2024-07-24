package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
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

type MenuLocation struct {
	InitialStep int
}

func InitMenu() MenuLocation {
	menu := MenuLocation{}
	menu.InitialStep = 0
	return menu
}

type SelectionItem struct {
	Path     string
	Name     string
	Date     string
	OnSelect func()
}

func NewSelectionItem(filename string, name string, date string, callback func()) SelectionItem {

	sitem := SelectionItem{}
	sitem.Path = fmt.Sprintf("%s/%s.txt", STOREPATH, filename)
	sitem.Name = name
	sitem.Date = date
	sitem.OnSelect = callback

	return sitem
}

func RenderHeader() {
	fmt.Print(HEADERSTRING)
}

func ReceiveEntry() bool {
	entname := "test-entry"
	input := bufio.NewReader(os.Stdin)
	entbytes := []byte("==================\n")

	// works for some linux distros only
	out, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	check(err)

	in, err := os.OpenFile("/dev/tty", os.O_RDONLY, 0)
	check(err)

	// figure out better ways to render this
	// animations and shit perhaps
	for {

		fmt.Println(reflect.TypeOf(in))

		fmt.Print(">>> ")

		line, err := input.ReadString('\n')
		check(err)

		// testing tty
		inbuff := []byte{}
		_, err = in.Read(inbuff)
		check(err)

		fmt.Println(inbuff)

		outbuff := []byte{}
		_, err = out.Read(outbuff)
		check(err)

		fmt.Println(outbuff)

		linebite := []byte(line)
		entbytes = append(entbytes, linebite...)

		if len(linebite) == 1 && linebite[0] == 10 {
			// EOF
			break
		}
	}

	fmt.Println(entbytes)
	err = os.WriteFile(fmt.Sprintf("%s/%s.txt", STOREPATH, entname), entbytes, 0644)
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
